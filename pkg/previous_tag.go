package history

import (
	"errors"
	"log"

	"gopkg.in/src-d/go-git.v4/plumbing"
)

var (
	// ErrPrevTagNotAvailable is returned when no previous tag is found.
	ErrPrevTagNotAvailable = errors.New("previous tag is not available")
)

// PreviousTag sorts tags based on when their commit happened and returns the one previous
// to the current.
func (g *Git) PreviousTag(currentHash plumbing.Hash) (*Tag, error) {
	tags, err := g.getTags()

	if err != nil {
		if g.Debug {
			log.Printf("[ERR] getting previous tag failed: %v", err)
		}
		return nil, err
	}

	// If there are fewer than two tags assume that the currentCommit is the newest tag
	if len(tags) < 2 {
		if g.Debug {
			log.Println("[ERR] previous tag not available")
		}
		return nil, ErrPrevTagNotAvailable
	}

	if tags[0].Hash != currentHash {
		if g.Debug {
			log.Println("[ERR] current commit does not have a tag attached, building from this commit")
		}
		return tags[0], nil
	}

	if g.Debug {
		log.Printf("success: previous tag found at %v", tags[1].Hash)
	}

	return tags[1], nil
}
