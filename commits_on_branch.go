package git

import (
	"errors"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// ErrCommonCommitFound is used for identifying when the iterator has reached the common commit
var ErrCommonCommitFound = errors.New("common commit found")

// CommitsOnBranch iterates through all references and returns commit hashes on given branch. \n
// Important to note is that this will provide all commits from anything the branch is connected to.
func (g *Git) CommitsOnBranch(
	branchCommit plumbing.Hash,
) ([]plumbing.Hash, error) {
	var branchCommits []plumbing.Hash

	branchIter, err := g.repo.Log(&git.LogOptions{
		From: branchCommit,
	})

	if err != nil {
		return nil, err
	}

	branchIterErr := branchIter.ForEach(func(commit *object.Commit) error {
		branchCommits = append(branchCommits, commit.Hash)
		return nil
	})

	if branchIterErr != nil {
		g.DebugLogger.Printf("Stopped getting commits on branch: %v", branchIterErr)
	}

	return branchCommits, nil
}
