package git

import (
	"log"
	"testing"
	"time"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/stretchr/testify/assert"
)

func createBranch(repo *git.Repository) {
	headRef, _ := repo.Head()

	ref := plumbing.NewHashReference("refs/heads/my-branch", headRef.Hash())

	storerErr := repo.Storer.SetReference(ref)

	if storerErr != nil {
		log.Println(storerErr)
	}

	worktree, _ := repo.Worktree()

	checkoutErr := worktree.Checkout(&git.CheckoutOptions{
		Branch: ref.Name(),
	})

	if checkoutErr != nil {
		log.Println(checkoutErr)
	}

}

func setupRepo() *git.Repository {
	repo, _ := git.Init(memory.NewStorage(), memfs.New())

	return repo
}

func createCommit(repo *git.Repository, message string) *object.Commit {
	w, _ := repo.Worktree()

	commit, _ := w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	obj, _ := repo.CommitObject(commit)

	return obj
}

func createTestHistory(repo *git.Repository) {
	createCommit(repo, "test commit on master")
	createBranch(repo)
	createCommit(repo, "commit on new branch")
	createCommit(repo, "second commit on new branch\n\n Long message")
	createCommit(repo, "third commit on new branch")
}

func TestCommitsOnBranch(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	head, _ := repo.Head()

	testGit := &Git{repo: repo}

	commits, err := testGit.CommitsOnBranch(head.Hash())

	assert.Equal(t, 4, len(commits))

	commit, commitErr := repo.CommitObject(commits[0])

	assert.NoError(t, commitErr)
	assert.Equal(t, "third commit on new branch", commit.Message)
	assert.Equal(t, err, nil)

	lastCommit, _ := repo.CommitObject(commits[3])

	assert.Equal(t, "test commit on master", lastCommit.Message)
}
