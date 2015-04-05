package fakku

import (
	"testing"
)

func Test_NewestManga_1(t *testing.T) {
	// currently this is broken
	_, err := NewestManga()
	if err != nil {
		t.Logf("NOTE: that this test is currently broken since GetCategoryIndex returns nothing!")
		t.Error(err)
	}
}
func Test_NewestDoujinshi_1(t *testing.T) {
	// currently this is broken
	_, err := NewestDoujinshi()
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
	_, err := ContentSearch("pool")
	if err != nil {
		t.Error(err)
	}
}
