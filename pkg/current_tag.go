package history

import "errors"

var (
	// ErrCommitNotOnTag is returned by CurrentTag if the commit is not on a tag
	ErrCommitNotOnTag = errors.New("current commit is not on a tag")
)

// CurrentTag returns a Tag if the current HEAD is on a tag
func (g *Git) CurrentTag() (*Tag, error) {
	head, err := g.repo.Head()

	if err != nil {
		return nil, err
	}

	tags, err := g.getTags()

	for _, tag := range tags {
		if tag.Hash == head.Hash() {
			return &tag, nil
		}
	}

	return nil, ErrCommitNotOnTag
}
