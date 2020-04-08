package git

import (
	"time"

	"github.com/go-git/go-git/v5/plumbing"
)

// Tag houses some common info about tags.
type Tag struct {
	Name string
	Hash plumbing.Hash
	Date time.Time
}
