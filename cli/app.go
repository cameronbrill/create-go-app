package cli

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type app struct {
	hc        http.Client
	name      string
	directory string
	template  string // references a branch name in https://github.com/cameronbrill/go-project-template
}

func (a *app) clone() (err error) {
	a.directory, err = getProjectName(a.name)
	if err != nil {
		fmt.Printf("Failed to get project name: %v\n", err)
		return
	}
	ref := plumbing.NewBranchReferenceName(a.template)
	_, err = git.PlainClone(a.directory, false, &git.CloneOptions{
		URL:           GithubRepoHost + TemplateRepoPath,
		ReferenceName: ref,
		SingleBranch:  true,
	})
	if err != nil {
		fmt.Printf("Failed to clone template: %v\nref: %s\n", err, ref)
		return
	}
	err = os.RemoveAll(fmt.Sprintf("./%s/.git/refs/remotes", a.directory))
	if err != nil {
		fmt.Printf("Failed to remove .git: %v\n", err)
		return
	}
	return nil
}

func getProjectName(s string) (string, error) {
	if _, err := os.Stat(fmt.Sprintf("./%s", s)); errors.Is(err, fs.ErrNotExist) {
		return getProjectName(s + "-1")
	} else if err != nil {
		return "", err
	}
	return s, nil
}