package history

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Tag houses some common info about tags.
type Tag struct {
	Name string
	Hash plumbing.Hash
}
