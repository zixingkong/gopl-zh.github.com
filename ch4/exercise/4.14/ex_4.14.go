package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Issue struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
}

type Issues []Issue

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		user := parts[1]
		repo := parts[2]

		issues, err := fetchIssues(user, repo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.New("issues").Parse(`
			<h1>{{.Repo}} Issues</h1>
			<ul>
			{{range .Issues}}
				<li><a href="https://github.com/{{$.User}}/{{$.Repo}}/issues/{{.Number}}">{{.Title}}</a></li>
			{{end}}
			</ul>
		`))

		tmpl.Execute(w, struct {
			User   string
			Repo   string
			Issues Issues
		}{
			User:   user,
			Repo:   repo,
			Issues: issues,
		})
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func fetchIssues(user, repo string) (Issues, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", user, repo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issues Issues
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}

	return issues, nil
}
