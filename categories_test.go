package fakku

import (
	"testing"
)

func TestGetCategoryIndex_1(t *testing.T) {
	// currently this is broken
	_, err := GetCategoryIndex(CategoryManga)
	if err != nil {
		t.Logf("NOTE: that this test is currently broken since GetCategoryIndex returns nothing!")
		t.Error(err)
	}
}

func TestGetTags_1(t *testing.T) {
	tags, err := Tags()
	foundTag := false
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(tags); i++ {
		tag := tags[i]
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
