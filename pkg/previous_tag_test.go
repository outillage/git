package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
)

func TestPreviousTag(t *testing.T) {
	repo, _ := git.PlainOpen("../testdata/git_tags")
	testGit := &Git{repo: repo, Debug: true}

	head, err := repo.Head()

	assert.NoError(t, err)

	tag, err := testGit.PreviousTag(head.Hash())

	assert.NoError(t, err)

	commit, err := repo.CommitObject(tag.Hash)
	assert.NoError(t, err)
	assert.Equal(t, "chore: first commit on master\n", commit.Message)

}

func TestPreviousTagDetachedTag(t *testing.T) {
	repo, _ := git.PlainOpen("../testdata/detached_tags")
	testGit := &Git{repo: repo, Debug: true}

	head, err := repo.Head()

	assert.NoError(t, err)

	tag, err := testGit.PreviousTag(head.Hash())

	assert.NoError(t, err)

	commit, err := repo.CommitObject(tag.Hash)
	assert.NoError(t, err)
	assert.Equal(t, "chore: detached commit\n", commit.Message)

}
