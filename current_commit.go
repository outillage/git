package git

import (
	"github.com/go-git/go-git/v5/plumbing/object"
)

// CurrentCommit returns the commit that HEAD is at
func (g *Git) CurrentCommit() (*object.Commit, error) {
	head, err := g.repo.Head()

	if err != nil {
		return nil, err
	}

	currentCommitHash := head.Hash()

	g.DebugLogger.Printf("current commitHash: %v \n", currentCommitHash)

	commitObject, err := g.repo.CommitObject(currentCommitHash)

	if err != nil {
		return nil, err
	}

	return commitObject, nil
}
