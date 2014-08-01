package fakku

import (
	"testing"
)

func TestGetCategoryIndex_1(t *testing.T) {
	_, err := GetCategoryIndex(CategoryManga)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTags_1(t *testing.T) {
	tags, err := GetTags()
	foundTag := false
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(tags.Tags); i++ {
		tag := tags.Tags[i]
		if tag.Name == ContentTestingTag {
			foundTag = true
		}
	}
	if !foundTag {
		t.Error("Couldn't find vanilla tag...something is very wrong :(")
	}
}

func TestContentSearch_1(t *testing.T) {
	_, err := GetContentSearchResults("pool")
	if err != nil {
		t.Error(err)
	}
}
