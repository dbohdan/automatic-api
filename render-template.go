// A script to generate a README.md from README.md.template and the data
// in data/. It fetches statistics about the repositories from GitHub.
// Copyright (c) dbohdan, 2017.
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

	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

type entry struct {
	Name    string
	URL     string
	DB      string
	API     string
	Lang    string
	License string
	Stats   string
	Notes   string
}
type table [][]string
type repo struct {
	Name  string
	Owner string
}
type projStats struct {
	Name string
	Ref  struct {
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

// Load entry data from the YAML files that match the glob pattern.
func loadEntries(glob string) (entries []entry, err error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	entries = make([]entry, len(matches))
	for i, match := range matches {
		data, err := ioutil.ReadFile(match)
		if err != nil {
			return nil, err
		}
		yaml.Unmarshal(data, &entries[i])
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

// Convert a table to its Markdown representation. Each column in the returned
// Markdown will have a consistent width to be more readable as plain text.
func (tbl table) Format() (res string) {
	buffer := bytes.NewBufferString("")
	tw := tablewriter.NewWriter(buffer)
	tw.SetHeader(tbl[0])
	tw.SetAutoFormatHeaders(false)
	tw.SetAutoWrapText(false)
	tw.SetBorders(tablewriter.Border{
		Left:   true,
		Top:    false,
		Right:  true,
		Bottom: false})
	tw.SetCenterSeparator("|")
	tw.AppendBulk(tbl[1:])
	tw.Render()
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
		  ref(qualifiedName: "master") {
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
		latestCommitDate := v.Ref.Target.AuthoredDate.
			Format("2006-01-02")
		statsHTML[i] = fmt.Sprintf(
			"%d&nbsp;★; %d&nbsp;commits, latest&nbsp;%s",
			v.Stargazers.TotalCount,
			v.Ref.Target.History.TotalCount,
			latestCommitDate)
	}
	return statsHTML, nil
}

func main() {
	tmpl, err := template.ParseFiles("README.md.template")
	if err != nil {
		log.Fatal(err)
	}
	entries, err := loadEntries("data/*.yml")
	if err != nil {
		log.Fatal(err)
	}

	re, err := regexp.Compile(
		"https?://github.com/([a-z0-9-]+)/([a-zA-Z0-9_-]+)/?")
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
