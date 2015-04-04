// Categories functions
package fakku

import (
	"encoding/json"
	"fmt"
	"net/url"
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
	return fmt.Sprintf("%s/%s", apiHeader, c.Category)
}

func GetCategoryIndex(category string) (*CategoryIndex, error) {
	var c CategoryIndex
	url := categoryIndexApiFunction{Category: category}
	if err := apiCall(url, &c); err != nil {
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
	var err error
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	if _, errExists := m["error"]; errExists {
		return fmt.Errorf("CategoryIndex request yielded an error!")
	}
	c.Latest, err = populateCategory("latest", m)
	if err != nil {
		return err
	}
	c.Favorites, err = populateCategory("favorites", m)
	if err != nil {
		return err
	}

	c.Popular, err = populateCategory("popular", m)
	if err != nil {
		return err
	}
	c.Controversial, err = populateCategory("controversial", m)
	if err != nil {
		return err
	}

	return nil
}
func populateCategory(category string, container map[string]interface{}) (ContentList, error) {
	q, result := container[category].([]interface{})
	if !result {
		return nil, fmt.Errorf("Category %s does not exist!", category)
	}
	l := make(ContentList, len(q))
	for i := 0; i < len(q); i++ {
		l[i].populate(q[i].(map[string]interface{}))
	}
	return l, nil

}

type tags struct {
	Tags  []Tag `json:"tags"`
	Total uint  `json:"total"`
}

type Tag struct {
	Name           string `json:"tag_name"`
	RawUrl         string `json:"tag_url"`
	RawImageSample string `json:"tag_image_sample"`
	Description    string `json:"tag_description"`
}

func (this *Tag) Url() (*url.URL, error) {
	return url.Parse(this.RawUrl)
}
func (this *Tag) ImageSampleUrl() (*url.URL, error) {
	return url.Parse(this.RawImageSample)
}

type tagsApiFunction struct{}

func (c tagsApiFunction) Construct() string {
	return fmt.Sprintf("%s/tags", apiHeader)
}

func Tags() ([]Tag, error) {
	var c tags
	url := tagsApiFunction{}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		if c.Total != uint(len(c.Tags)) {
			return nil, fmt.Errorf("Count mismatch! Expected %d tags but got %d instead!", c.Total, len(c.Tags))
		} else {
			return c.Tags, nil
		}
	}
}

type ContentSearchResults struct {
	Content ContentList `json:"content"`
	Total   uint        `json:"total"`
	Pages   uint        `json:"pages"`
}

type contentSearchApiFunction struct {
	Terms string
	supportsPagination
}

func (c contentSearchApiFunction) Construct() string {
	base := fmt.Sprintf("%s/search/%s", apiHeader, c.Terms)
	return paginateString(base, c.Page)
}

func ContentSearchPage(terms string, page uint) (*ContentSearchResults, error) {
	var c ContentSearchResults
	url := contentSearchApiFunction{
		Terms:              terms,
		supportsPagination: supportsPagination{Page: page},
	}
	if err := apiCall(url, c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func ContentSearch(terms string) (*ContentSearchResults, error) {
	return ContentSearchPage(terms, 0)
}
