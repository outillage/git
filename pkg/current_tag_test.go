package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
)

func TestCurrentTagHappy(t *testing.T) {
	repo, err := git.PlainOpen("../testdata/git_tags")
	testGit := &Git{repo: repo, Debug: true}

	assert.NoError(t, err)

	tag, err := testGit.CurrentTag()

	assert.NoError(t, err)
	assert.Equal(t, "refs/tags/v0.0.2", tag.Name)
}

func TestCurrentTagAnnotatedHappy(t *testing.T) {
	repo, err := git.PlainOpen("../testdata/annotated_git_tags_mix")
	testGit := &Git{repo: repo, Debug: true}

	assert.NoError(t, err)

	tag, err := testGit.CurrentTag()

	assert.NoError(t, err)
	assert.Equal(t, "v0.0.3", tag.Name)
}

func TestCurrentTagUnhappy(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	testGit := &Git{repo: repo}

	_, err := testGit.CurrentTag()

	assert.Equal(t, ErrCommitNotOnTag, err)
}
