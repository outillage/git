package git

import (
	"github.com/go-git/go-git/v5"
)

// Git is the struct used to house all methods in use in Commitsar.
type Git struct {
	repo *git.Repository
}

// OpenGit loads Repo on path and returns a new Git struct to work with.
func OpenGit(path string) (*Git, error) {
	repo, repoErr := git.PlainOpen(path)

	if repoErr != nil {
		return nil, repoErr
	}

	return &Git{repo: repo}, nil
}
