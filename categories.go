// Categories functions
package fakku

import (
	"fmt"
)

const (
	CategoryManga     = "manga"
	CategoryDoujinshi = "doujinshi"
	CategoryVideos    = "videos" // Is this a legal category?
)

func LegalCategory(category string) bool {
	return category == CategoryManga || category == CategoryDoujinshi
}

type categoryIndexApiFunction struct {
	Category string
	Name     string
}

func (c categoryIndexApiFunction) Construct() string {
	return fmt.Sprintf("%s/%s/%s", apiHeader, c.Category, c.Name)
}

func CategoryIndex(category, name string, c interface{}) error {
	url := categoryIndexApiFunction{
		Category: category,
		Name:     name,
	}
	if err := apiCall(url, &c); err != nil {
		return err
	} else {
		return nil
	}

}

func MangaCategoryIndex(name string, c interface{}) error {
	return CategoryIndex(CategoryManga, name, c)
}

func DoujinshiCategoryIndex(name string, c interface{}) error {
	return CategoryIndex(CategoryDoujinshi, name, c)
}

type newestCategory struct {
	Collection ContentList `json:"newest"`
}
type englishCategory struct {
	Collection ContentList `json:"english"`
}
type favoritesCategory struct {
	Collection ContentList `json:"favorites"`
}
type popularCategory struct {
	Collection ContentList `json:"popular"`
}
type controversialCategory struct {
	Collection ContentList `json:"controversial"`
}

func ControversialManga() (ContentList, error) {
	var q controversialCategory
	if err := MangaCategoryIndex("controversial", &q); err != nil {
		return nil, err
	} else {
		return q.Collection, nil
	}
}

func ControversialDoujinshi() (ContentList, error) {
	var q controversialCategory
	if err := DoujinshiCategoryIndex("controversial", &q); err != nil {
		return nil, err
	} else {
		return q.Collection, nil
	}
}

func PopularManga() (ContentList, error) {
	var q popularCategory
	if err := MangaCategoryIndex("popular", &q); err != nil {
		return nil, err
	} else {
		return q.Collection, nil
	}
}

func PopularDoujinshi() (ContentList, error) {
	var q popularCategory
	if err := DoujinshiCategoryIndex("popular", &q); err != nil {
		return nil, err
	} else {
		return q.Collection, nil
	}
}

func FavoritesManga() (ContentList, error) {
	var q favoritesCategory
	if err := MangaCategoryIndex("favorites", &q); err != nil {
		return nil, err
	} else {
		return q.Collection, nil
	}
}

func FavoritesDoujinshi() (ContentList, error) {
	var q favoritesCategory
	if err := DoujinshiCategoryIndex("favorites", &q); err != nil {
		return nil, err
	} else {
		return q.Collection, nil
	}
}

func EnglishManga() (ContentList, error) {
	var q englishCategory
	if err := MangaCategoryIndex("english", &q); err != nil {
		return nil, err
	} else {
		return q.Collection, nil
	}
}

func EnglishDoujinshi() (ContentList, error) {
	var q englishCategory
	if err := DoujinshiCategoryIndex("english", &q); err != nil {
		return nil, err
	} else {
		return q.Collection, nil
	}
}

func NewestManga() (ContentList, error) {
	var q newestCategory
	if err := MangaCategoryIndex("newest", &q); err != nil {
		return nil, err
	} else {
		return q.Collection, nil
	}
}

func NewestDoujinshi() (ContentList, error) {
	var q newestCategory
	if err := DoujinshiCategoryIndex("newest", &q); err != nil {
		return nil, err
	} else {
		return q.Collection, nil
	}
}
