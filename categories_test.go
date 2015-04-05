package fakku

import (
	"testing"
)

func Test_NewestManga_1(t *testing.T) {
	// currently this is broken
	_, err := NewestManga()
	if err != nil {
		t.Error(err)
	}
}
func Test_NewestDoujinshi_1(t *testing.T) {
	// currently this is broken
	_, err := NewestDoujinshi()
	if err != nil {
		t.Error(err)
	}
}

func Test_EnglishManga_1(t *testing.T) {
	// currently this is broken
	_, err := EnglishManga()
	if err != nil {
		t.Error(err)
	}
}
func Test_EnglishDoujinshi_1(t *testing.T) {
	// currently this is broken
	_, err := EnglishDoujinshi()
	if err != nil {
		t.Error(err)
	}
}
func Test_FavoritesManga_1(t *testing.T) {
	// currently this is broken
	_, err := FavoritesManga()
	if err != nil {
		t.Error(err)
	}
}
func Test_FavoritesDoujinshi_1(t *testing.T) {
	// currently this is broken
	_, err := FavoritesDoujinshi()
	if err != nil {
		t.Error(err)
	}
}

func Test_PopularManga_1(t *testing.T) {
	// currently this is broken
	_, err := PopularManga()
	if err != nil {
		t.Error(err)
	}
}
func Test_PopularDoujinshi_1(t *testing.T) {
	// currently this is broken
	_, err := PopularDoujinshi()
	if err != nil {
		t.Error(err)
	}
}
func Test_ControversialManga_1(t *testing.T) {
	// currently this is broken
	_, err := ControversialManga()
	if err != nil {
		t.Error(err)
	}
}
func Test_ControversialDoujinshi_1(t *testing.T) {
	// currently this is broken
	_, err := ControversialDoujinshi()
	if err != nil {
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
