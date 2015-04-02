// Categories functions
package fakku

import (
	"encoding/json"
	"fmt"
)

const (
	CategoryManga     = "manga"
	CategoryDoujinshi = "doujinshi"
	CategoryVideos    = "videos" // Is this a legal category?
)

type categoryIndexApiFunction struct {
	Category string
}

func (c categoryIndexApiFunction) Construct() string {
	return fmt.Sprintf("%s/%s", ApiHeader, c.Category)
}

func GetCategoryIndex(category string) (*CategoryIndex, error) {
	var c CategoryIndex
	url := categoryIndexApiFunction{Category: category}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}

}

type CategoryIndex struct {
	Latest        ContentList
	Favorites     ContentList
	Popular       ContentList
	Controversial ContentList
}

func (c *CategoryIndex) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	if _, errExists := m["error"]; errExists {
		return fmt.Errorf("CategoryIndex request yielded an error!")
	}
	latest := m["latest"].([]interface{})
	c.Latest = make(ContentList, len(latest))
	for i := 0; i < len(latest); i++ {
		c.Latest[i].populate(latest[i].(map[string]interface{}))
	}
	favorites := m["favorites"].([]interface{})
	c.Favorites = make(ContentList, len(favorites))
	for i := 0; i < len(favorites); i++ {
		c.Favorites[i].populate(favorites[i].(map[string]interface{}))
	}
	popular := m["popular"].([]interface{})
	c.Popular = make([]Content, len(popular))
	for i := 0; i < len(popular); i++ {
		c.Popular[i].populate(popular[i].(map[string]interface{}))
	}
	controversial := m["controversial"].([]interface{})
	c.Controversial = make([]Content, len(controversial))
	for i := 0; i < len(controversial); i++ {
		c.Controversial[i].populate(controversial[i].(map[string]interface{}))
	}
	return nil
}

type tags struct {
	Tags  []Tag `json:"tags"`
	Total uint  `json:"total"`
}

type Tag struct {
	Name        string `json:"tag_name"`
	Url         string `json:"tag_url"`
	ImageSample string `json:"tag_image_sample"`
	Description string `json:"tag_description"`
}
type tagsApiFunction struct{}

func (c tagsApiFunction) Construct() string {
	return fmt.Sprintf("%s/tags", ApiHeader)
}

func Tags() ([]Tag, error) {
	var c tags
	url := tagsApiFunction{}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return c.Tags, nil
	}
}

type ContentSearch struct {
	Content []Content `json:"content"`
	Total   uint      `json:"total"`
	Pages   uint      `json:"pages"`
}

type ContentSearchApiFunction struct {
	Terms string
	SupportsPagination
}

func (c ContentSearchApiFunction) Construct() string {
	base := fmt.Sprintf("%s/search/%s", ApiHeader, c.Terms)
	return PaginateString(base, c.Page)
}

func GetContentSearchResultsPage(terms string, page uint) (*ContentSearch, error) {
	var c ContentSearch
	url := ContentSearchApiFunction{
		Terms:              terms,
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetContentSearchResults(terms string) (*ContentSearch, error) {
	return GetContentSearchResultsPage(terms, 0)
}
