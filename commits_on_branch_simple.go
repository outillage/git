package git

import (
	"github.com/apex/log"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// SimpleCommit is a slimed down commit object of just Hash and Message
type SimpleCommit struct {
	Hash    [20]byte
	Message string
}

// CommitsOnBranchSimple iterates through all references and returns simpleCommits on given branch. \n
// Important to note is that this will provide all commits from anything the branch is connected to.
func (g *Git) CommitsOnBranchSimple(
	branchCommit plumbing.Hash,
) ([]SimpleCommit, error) {
	var branchCommits []SimpleCommit

	branchIter, err := g.repo.Log(&git.LogOptions{
		From: branchCommit,
	})

	if err != nil {
		return nil, err
	}

	branchIterErr := branchIter.ForEach(func(commit *object.Commit) error {
		branchCommits = append(branchCommits, SimpleCommit{
			Hash:    commit.Hash,
			Message: commit.Message,
		})
		return nil
	})

	if branchIterErr != nil {
		log.Debugf("Stopped getting commits on branch: %v", branchIterErr)
	}

	return branchCommits, nil
}
