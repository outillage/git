package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenGit(t *testing.T) {
	_, err := OpenGit("../", nil)

	// Should not error if this git repository is valid
	assert.NoError(t, err)

	_, unhappyErr := OpenGit(".", nil)

	// Should error opening a folder with missing .git
	assert.Error(t, unhappyErr)

}
