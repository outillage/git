package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
)

func TestGetTags(t *testing.T) {
	repo, err := git.PlainOpen("../testdata/git_tags")
	testGit := &Git{repo: repo, Debug: false}

	assert.NoError(t, err)

	tags, err := testGit.getTags()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(tags))

	assert.Equal(t, "refs/tags/v0.0.2", tags[0].Name)
	assert.Equal(t, "refs/tags/v0.0.1", tags[1].Name)

}

func TestAnnotatedGetTags(t *testing.T) {
	repo, err := git.PlainOpen("../testdata/annotated_git_tags_mix")
	testGit := &Git{repo: repo, Debug: false}

	assert.NoError(t, err)

	tags, err := testGit.getTags()

	assert.NoError(t, err)
	assert.Equal(t, 3, len(tags))

	assert.Equal(t, "v0.0.3", tags[0].Name)
	assert.Equal(t, "refs/tags/v0.0.2", tags[1].Name)

}
