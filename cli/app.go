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
		return fmt.Errorf("Failed to get project name %s: %w", a.name, err)
	}
	ref := plumbing.NewBranchReferenceName(a.template)
	_, err = git.PlainClone(a.directory, false, &git.CloneOptions{
		URL:           GithubRepoHost + TemplateRepoPath,
		ReferenceName: ref,
		SingleBranch:  true,
	})
	if err != nil {
		return fmt.Errorf("Failed to clone template: %w\nref: %s", err, ref)
	}
	err = os.RemoveAll(fmt.Sprintf("./%s/.git/refs/remotes", a.directory))
	if err != nil {
		return fmt.Errorf("Failed to remove .git: %w", err)
	}
	return nil
}

func getProjectName(s string) (string, error) {
	if f, err := os.Stat(fmt.Sprintf("./%s", s)); f != nil {
		return getProjectName(s + "-1")
	} else if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return "", err
	}
	return s, nil
}
