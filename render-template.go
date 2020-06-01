// This script generates README.md from README.md.template and the data
// in data/.  It fetches statistics about GitHub repositories from GitHub.
// For the script to work the environment variable GITHUB_TOKEN must contain
// a valid personal access token (see https://github.com/settings/tokens).
//
// Copyright (c) D. Bohdan 2017, 2018, 2019, 2020.
// License: MIT.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"text/template"
	"time"

	"github.com/magiconair/properties"
	"github.com/rhinoman/go-commonmark"
)

type entry struct {
	Name    string `properties:"name"`
	URL     string `properties:"url"`
	DB      string `properties:"db"`
	API     string `properties:"api"`
	Lang    string `properties:"lang"`
	License string `properties:"license"`
	Stats   string `properties:"stats,default="`
	Notes   string `properties:"notes,default="`
}

type table [][]string

type repo struct {
	Name  string
	Owner string
}

type projStats struct {
	Name             string
	DefaultBranchRef struct {
		Target struct {
			AuthoredDate time.Time
			History      struct {
				TotalCount int
			}
		}
	}
	Stargazers struct {
		TotalCount int
	}
}

type gitHubResponse struct {
	Data    map[string]projStats
	Message string
	Errors  []map[string]interface{}
}

// Load entry data from properties files that match the glob pattern.
func loadEntries(glob string) (entries []entry, err error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}

	entries = make([]entry, len(matches))
	for i, match := range matches {
		p, err := properties.LoadFile(match, properties.UTF8)
		if err != nil {
			return nil, err
		}

		err = p.Decode(&entries[i])
		if err != nil {
			return nil, err
		}
	}

	return entries, nil
}

// Convert each entry into a table row.
func entriesToTable(entries []entry) (tbl table) {
	tbl = make(table, len(entries)+1)
	tbl[0] = make([]string, 7)
	tbl[0][0] = "Project name/link"
	tbl[0][1] = "Database(s) supported"
	tbl[0][2] = "API type"
	tbl[0][3] = "Implementation language"
	tbl[0][4] = "License"
	tbl[0][5] = "GitHub stats"
	tbl[0][6] = "Notes"

	for i, ent := range entries {
		j := i + 1
		tbl[j] = make([]string, 7)

		tbl[j][0] = fmt.Sprintf("[%s](%s)",
			html.EscapeString(ent.Name),
			html.EscapeString(ent.URL))
		tbl[j][1] = html.EscapeString(ent.DB)
		tbl[j][2] = html.EscapeString(ent.API)
		tbl[j][3] = html.EscapeString(ent.Lang)
		tbl[j][4] = html.EscapeString(ent.License)
		// Deliberately allow HTML in Stats and Notes.
		tbl[j][5] = ent.Stats
		tbl[j][6] = ent.Notes
	}

	return tbl
}

// <p>foo</p> -> foo
func stripP(html string) (res string) {
	p := regexp.MustCompile("^<p>(.*?)</p>\n?$")
	found := p.FindStringSubmatch(html)
	if len(found) == 2 {
		return found[1]
	}
	return html
}

// Convert a table with each cell containing Markdown to an HTML <table>
// representation.
func (tbl table) Format() (res string) {
	buffer := bytes.NewBufferString("<table>\n")

	buffer.WriteString("  <tr>\n")

	for _, h := range tbl[0] {
		buffer.WriteString(fmt.Sprintf("    <th>%s</th>\n", h))
	}

	buffer.WriteString("  </tr>\n")

	for _, row := range tbl[1:] {
		buffer.WriteString("  <tr>\n")
		for _, col := range row {
			colHTML := stripP(commonmark.Md2Html(col, 0))
			buffer.WriteString(
				fmt.Sprintf("    <td>%s</td>\n", colHTML))
		}
		buffer.WriteString("  </tr>\n")
	}

	buffer.WriteString("</table>\n")

	return buffer.String()
}

// Perform a GraphQL query against the GitHub API v4.
func queryGitHub(token string, query string) (body []byte, err error) {
	url := "https://api.github.com/graphql"

	wrapper := map[string]string{}
	wrapper["query"] = query

	jsonReq, err := json.Marshal(wrapper)
	if err != nil {
		return nil, err
	}

	jsonReqBuf := bytes.NewBuffer([]byte(jsonReq))

	req, err := http.NewRequest("POST", url, jsonReqBuf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
	req.Header.Set("Content-Type", "application/graphql")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Build a GraphQL query to fetch the GitHub repository statistics for several
// repositories at once.
func buildStatsQuery(repos []repo) (query string, err error) {
	repoQuery := `r%d: repository(owner: "%s", name: "%s") { ...ProjStats }`
	projStatsFragment := `
		fragment ProjStats on Repository {
		  name
		  defaultBranchRef {
		    target {
		      ... on Commit {
		        authoredDate
		        history {
		          totalCount
		        }
		      }
		    }
		  }
		  stargazers {
		    totalCount
		  }
		}
	`
	queryBuf := bytes.NewBufferString("{\n")

	safe, err := regexp.Compile("^[a-zA-Z0-9_-]+$")
	if err != nil {
		return "", err
	}

	for i, r := range repos {
		// We do not make a fuss about empty repository entries. We
		// simply skip them here and return an empty string for them at
		// the end.
		if r.Name != "" {
			if !(safe.MatchString(r.Owner) &&
				safe.MatchString(r.Name)) {
				return "", fmt.Errorf("Bad repo data: %v", r)
			}

			t := fmt.Sprintf(repoQuery, i, r.Owner, r.Name)
			queryBuf.WriteString(t)
			queryBuf.WriteString("\n")
		}
	}

	queryBuf.WriteString("}\n")
	queryBuf.WriteString(projStatsFragment)

	return queryBuf.String(), nil
}

// Fetch the repository statistics (the number of stars and commits) for the
// repositories in repos and return them formatted as HTML.
func fetchGitHubStats(token string, repos []repo) (statsHTML []string,
	err error) {
	query, err := buildStatsQuery(repos)
	if err != nil {
		return nil, err
	}

	respBody, err := queryGitHub(token, query)
	if err != nil {
		return nil, err
	}

	var respData gitHubResponse
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		return nil, err
	}

	if respData.Message != "" {
		return nil, fmt.Errorf("GitHub API: %s", respData.Message)
	}

	if len(respData.Errors) > 0 {
		return nil, fmt.Errorf("GitHub API: %s",
			respData.Errors[0]["message"].(string))
	}

	statsHTML = make([]string, len(repos))
	for k, v := range respData.Data {
		i, err := strconv.Atoi(k[1:])
		if err != nil {
			return nil, err
		}

		latestCommitDate := v.DefaultBranchRef.Target.AuthoredDate.
			Format("2006-01-02")

		statsHTML[i] = fmt.Sprintf(
			"%d&nbsp;★; %d&nbsp;commits, latest&nbsp;%s",
			v.Stargazers.TotalCount,
			v.DefaultBranchRef.Target.History.TotalCount,
			latestCommitDate)
	}

	return statsHTML, nil
}

func main() {
	tmpl, err := template.ParseFiles("README.md.template")
	if err != nil {
		log.Fatal(err)
	}

	entries, err := loadEntries("data/*")
	if err != nil {
		log.Fatal(err)
	}

	re, err := regexp.Compile(
		"https?://github.com/([a-zA-Z0-9-]+)/([a-zA-Z0-9_-]+)/?")
	if err != nil {
		log.Fatal(err)
	}

	repos := make([]repo, len(entries))
	for i, ent := range entries {
		found := re.FindStringSubmatch(ent.URL)
		if len(found) == 3 {
			repos[i].Owner = found[1]
			repos[i].Name = found[2]
		} else {
			entries[i].Stats = "n/a"
		}
	}

	statsHTML, err := fetchGitHubStats(os.Getenv("GITHUB_TOKEN"), repos)
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range statsHTML {
		if s != "" {
			entries[i].Stats = s
		}
	}

	tbl := entriesToTable(entries)
	dot := map[string]interface{}{}
	dot["date"] = time.Now().Format("2006-01-02")
	dot["table"] = tbl.Format()

	tmpl.Execute(os.Stdout, dot)
}
