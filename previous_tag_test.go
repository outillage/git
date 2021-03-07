package git

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
)

func TestPreviousTag(t *testing.T) {
	repo, _ := git.PlainOpen("./testdata/git_tags")
	testGit := &Git{repo: repo}

	head, err := repo.Head()

	assert.NoError(t, err)

	tag, err := testGit.PreviousTag(head.Hash())

	assert.NoError(t, err)

	commit, err := repo.CommitObject(tag.Hash)
	assert.NoError(t, err)
	assert.Equal(t, "chore: first commit on master\n", commit.Message)

}
