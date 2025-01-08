// This script generates README.md
// from README.md.tmpl, table.md.tmpl, and the data in data/.
// It fetches statistics about GitHub repositories from GitHub.
// For the script to work, the environment variable GITHUB_TOKEN must contain
// a valid personal access token (see https://github.com/settings/tokens).
//
// Copyright (c) D. Bohdan 2017-2019, 2020, 2023-2025.
// License: MIT.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/yuin/goldmark"
)

type entry struct {
	Name    string `toml:"name"`
	URL     string `toml:"url"`
	DB      string `toml:"db"`
	API     string `toml:"api"`
	Lang    string `toml:"lang"`
	License string `toml:"license"`
	Stats   string `toml:"stats"`
	Notes   string `toml:"notes"`
}

type tableRow struct {
	Cells []string
}

type table struct {
	Rows []tableRow
}

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

const (
	readmeTemplateFile = "README.md.tmpl"
	tableTemplateFile  = "table.md.tmpl"

	repoQuery         = `r%d: repository(owner: "%s", name: "%s") { ...ProjStats }`
	projStatsFragment = `
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
)

const (
	colProjectName = iota
	colDatabase
	colAPIType
	colLanguage
	colLicense
	colStats
	colNotes
	colCount
)

var (
	gitHubURL       = regexp.MustCompile("https?://github.com/([a-zA-Z0-9-]+)/([a-zA-Z0-9_-]+)/?")
	pTag            = regexp.MustCompile(`(?s)^\s*<p>(.*?)</p>\s*$`)
	safeRepoPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
)

// loadEntries loads entry data from TOML files that match the glob pattern.
func loadEntries(glob string) (entries []entry, err error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return nil, fmt.Errorf("loading entries from glob pattern: %w", err)
	}

	entries = make([]entry, len(matches))
	for i, match := range matches {
		buf, err := os.ReadFile(match)
		if err != nil {
			return nil, fmt.Errorf("reading file %s: %w", match, err)
		}

		err = toml.Unmarshal(buf, &entries[i])
		if err != nil {
			return nil, fmt.Errorf("parsing TOML file %s: %w", match, err)
		}
	}

	return entries, nil
}

// entriesToTable converts a slice of entries to a representation of an HTML table.
func entriesToTable(entries []entry) *table {
	rows := make([]tableRow, len(entries))

	for i, ent := range entries {
		cells := make([]string, colCount)

		cells[colProjectName] = fmt.Sprintf("[%s](%s)",
			html.EscapeString(ent.Name),
			html.EscapeString(ent.URL))
		cells[colDatabase] = html.EscapeString(ent.DB)
		cells[colAPIType] = html.EscapeString(ent.API)
		cells[colLanguage] = html.EscapeString(ent.Lang)
		cells[colLicense] = html.EscapeString(ent.License)
		// Deliberately allow HTML in stats and notes.
		cells[colStats] = ent.Stats
		cells[colNotes] = ent.Notes

		rows[i] = tableRow{Cells: cells}
	}

	return &table{Rows: rows}
}

// trimP converts "<p>foo</p>" to "foo".
func trimP(html string) (res string) {
	found := pTag.FindStringSubmatch(html)
	if len(found) == 2 {
		return found[1]
	}

	return html
}

// format converts the table to HTML using tmpl.
func (t *table) format(tmpl template.Template) (string, error) {
	for i := range t.Rows {
		for j, cell := range t.Rows[i].Cells {
			var buf bytes.Buffer

			if err := goldmark.Convert([]byte(cell), &buf); err != nil {
				return "", fmt.Errorf("converting Markdown: %w", err)
			}

			t.Rows[i].Cells[j] = trimP(buf.String())
		}
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, t); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}

	return buf.String(), nil
}

// queryGitHub perform a GraphQL query against the GitHub API v4.
func queryGitHub(token, query string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wrapper := map[string]string{"query": query}
	jsonReq, err := json.Marshal(wrapper)
	if err != nil {
		return nil, fmt.Errorf("marshal query: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://api.github.com/graphql",
		bytes.NewReader(jsonReq),
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Content-Type", "application/graphql")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	return body, nil
}

// buildStatsQuery builds a GraphQL query
// to fetch the GitHub repository statistics for several repositories at once.
func buildStatsQuery(repos []repo) (string, error) {
	var sb strings.Builder
	sb.WriteString("{\n")

	for i, r := range repos {
		if r.Name == "" {
			continue
		}

		if !safeRepoPattern.MatchString(r.Owner) || !safeRepoPattern.MatchString(r.Name) {
			return "", fmt.Errorf("invalid repo name or owner %v", r)
		}

		fmt.Fprintf(&sb, repoQuery, i, r.Owner, r.Name)
		sb.WriteString("\n")
	}

	sb.WriteString("}\n")
	sb.WriteString(projStatsFragment)

	return sb.String(), nil
}

// fetchGitHubStats retrieves the repository statistics
// (the number of stars and commits)
// for the repositories in repos and returns them formatted as HTML.
func fetchGitHubStats(token string, repos []repo) (statsHTML []string, err error) {
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
		return nil, fmt.Errorf("GitHub API error: %s", respData.Message)
	}

	if len(respData.Errors) > 0 {
		return nil, fmt.Errorf("GitHub GraphQL error: %s",
			respData.Errors[0]["message"].(string))
	}

	statsHTML = make([]string, len(repos))
	for k, v := range respData.Data {
		i, err := strconv.Atoi(k[1:])
		if err != nil {
			return nil, err
		}

		latestCommitDate := v.DefaultBranchRef.Target.AuthoredDate.
			Format(time.DateOnly)

		statsHTML[i] = fmt.Sprintf(
			"%d&nbsp;★; %d&nbsp;commits, latest&nbsp;%s",
			v.Stargazers.TotalCount,
			v.DefaultBranchRef.Target.History.TotalCount,
			latestCommitDate)
	}

	return statsHTML, nil
}

func main() {
	readmeTmpl, err := template.ParseFiles(readmeTemplateFile)
	if err != nil {
		log.Fatal(err)
	}

	tableTmpl, err := template.ParseFiles(tableTemplateFile)
	if err != nil {
		log.Fatal(err)
	}

	entries, err := loadEntries("data/*.toml")
	if err != nil {
		log.Fatal(err)
	}

	repos := make([]repo, len(entries))
	for i, ent := range entries {
		found := gitHubURL.FindStringSubmatch(ent.URL)
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
	tableHTML, err := tbl.format(*tableTmpl)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"date":  time.Now().Format(time.DateOnly),
		"table": tableHTML,
	}

	err = readmeTmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
