package cli

import (
	"encoding/json"
	"net/http"
)

const (
	GithubAPIHost    = "https://api.github.com/repos"
	GithubRepoHost   = "https://github.com"
	TemplateRepoPath = "/cameronbrill/go-project-template"
	BranchesEndpoint = "/branches"
)

type RepoRes struct{}

type BranchRes struct {
	Name   string `json:"name"`
	Commit struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	Protected bool `json:"protected"`
}

type ApiResponse interface {
	RepoRes | []BranchRes
}

func fetchJSON[T ApiResponse](url string, hc http.Client, data *T) error {
	resp, err := hc.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}
