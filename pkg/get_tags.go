package history

import (
	"log"
	"sort"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func (g *Git) getTags() ([]*Tag, error) {
	tagrefs, err := g.repo.Tags()

	if err != nil {
		return nil, err
	}

	defer tagrefs.Close()

	var tags []*Tag

	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		tags = append(tags, &Tag{Name: t.Name().String(), Hash: t.Hash()})
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Tags are alphabetically ordered. We need to sort them by date.
	sortedTags := sortTags(g.repo, tags)

	if g.Debug {
		log.Println("Sorted tag output: ")
		for _, taginstance := range sortedTags {
			log.Printf("hash: %v name: %v", taginstance.Hash, taginstance.Name)
		}
	}

	return sortedTags, nil
}

// sortTags sorts the tags in reverse.
func sortTags(repo *git.Repository, tags []*Tag) []*Tag {
	sort.Slice(tags, func(i, j int) bool {
		// Reverses the order
		return j < i
	})

	return tags
}
