package history

import (
	"time"

	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Tag houses some common info about tags.
type Tag struct {
	Name string
	Hash plumbing.Hash
	Date time.Time
}
