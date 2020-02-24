package git

import (
	"sort"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func (g *Git) getTags() ([]*Tag, error) {
	tagrefs, err := g.repo.Tags()

	if err != nil {
		return nil, err
	}

	defer tagrefs.Close()

	var tags []*Tag

	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		commitDate, err := g.commitDate(t.Hash())

		if err != nil {
			g.DebugLogger.Printf("tag: %v ignored due to missing commit date: %v", t.Name(), err)
			return nil
		}

		tags = append(tags, &Tag{Name: t.Name().String(), Date: commitDate, Hash: t.Hash()})
		return nil
	})

	if err != nil {
		return nil, err
	}

	tagObjects, err := g.repo.TagObjects()

	if err != nil {
		return nil, err
	}

	err = tagObjects.ForEach(func(tag *object.Tag) error {
		tags = append(tags, &Tag{Name: tag.Name, Date: tag.Tagger.When, Hash: tag.Target})

		return nil
	})

	// Tags are alphabetically ordered. We need to sort them by date.
	sortedTags := sortTags(g.repo, tags)

	g.DebugLogger.Println("Sorted tag output: ")
	for _, taginstance := range sortedTags {
		g.DebugLogger.Printf("hash: %v time: %v", taginstance.Hash, taginstance.Date.UTC())
	}

	return sortedTags, nil
}

// sortTags sorts the tags according to when their parent commit happened.
func sortTags(repo *git.Repository, tags []*Tag) []*Tag {
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Date.After(tags[j].Date)
	})

	return tags
}