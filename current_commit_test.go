package git

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentCommit(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	testGit := &Git{repo: repo, DebugLogger: log.New(ioutil.Discard, "", 0)}

	currentCommit, err := testGit.CurrentCommit()

	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch", currentCommit.Message)
}
