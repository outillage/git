package git

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Commit find a commit based on commit hash and returns the Commit object
func (g *Git) Commit(hash plumbing.Hash) (*object.Commit, error) {
	commitObject, err := g.repo.CommitObject(hash)

	if err != nil {
		return nil, err
	}

	return commitObject, nil
}
