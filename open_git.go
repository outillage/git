package git

import (
	"io/ioutil"
	"log"

	"github.com/go-git/go-git/v5"
)

// Git is the struct used to house all methods in use in Commitsar.
type Git struct {
	repo *git.Repository
	// DebugLogger flag is passed to make debugging easier during development/problematic deploys
	DebugLogger *log.Logger
}

// OpenGit loads Repo on path and returns a new Git struct to work with.
func OpenGit(path string, debugLogger *log.Logger) (*Git, error) {
	repo, repoErr := git.PlainOpen(path)

	if repoErr != nil {
		return nil, repoErr
	}

	// Failsafe for when debug logger is omitted
	if debugLogger == nil {
		return &Git{repo: repo, DebugLogger: log.New(ioutil.Discard, "", 0)}, nil
	}

	return &Git{repo: repo, DebugLogger: debugLogger}, nil
}
