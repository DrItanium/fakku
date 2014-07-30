package fakku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Attribute struct {
	Attribute     string
	AttributeLink string
}
type Content struct {
	Name        string
	Url         string
	Description string
	Language    string
	Category    string
	Date        float64
	FileSize    float64
	Favorites   float64
	Comments    float64
	Pages       float64
	Poster      string
	PosterUrl   string
	Tags        []*Attribute
	Translators []*Attribute
	Series      []*Attribute
	Artists     []*Attribute
	Images      struct {
		Cover  string
		Sample string
	}
}

func (c *Content) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})

	contents := m["content"]
	v := contents.(map[string]interface{})

	c.Name = v["content_name"].(string)
	c.Url = v["content_url"].(string)
	c.Description = v["content_description"].(string)
	c.Language = v["content_language"].(string)
	c.Category = v["content_category"].(string)
	c.Date = v["content_date"].(float64)
	c.FileSize = v["content_filesize"].(float64)
	c.Favorites = v["content_favorites"].(float64)
	c.Comments = v["content_comments"].(float64)
	c.Pages = v["content_pages"].(float64)
	c.Poster = v["content_poster"].(string)
	c.PosterUrl = v["content_poster_url"].(string)
	c.Tags = constructAttributeFields(v, "content_tags")
	c.Translators = constructAttributeFields(v, "content_translators")
	c.Series = constructAttributeFields(v, "content_series")
	c.Artists = constructAttributeFields(v, "content_artists")

	tmp := v["content_images"]
	z := tmp.(map[string]interface{})
	c.Images.Cover = z["cover"].(string)
	c.Images.Sample = z["sample"].(string)

	return nil
}

func PaginateString(s string, page uint) string {
	return fmt.Sprintf("%s/page/%d", s, page)
}

type ApiFunction interface {
	ConstructApiFunction() string
}

type ContentApiFunction struct {
	Category string
	Name     string
}

var ApiHeader = "https://api.fakku.net/"

func (a *ContentApiFunction) ConstructApiFunction() string {
	return fmt.Sprintf("%s/%s/%s", ApiHeader, a.Category, a.Name)
}

type ContentCommentApiFunction struct {
	*ContentApiFunction
	TopComments bool
	Page        uint
}

func (a *ContentCommentApiFunction) ConstructApiFunction() string {
	base := fmt.Sprintf("%s/comments", a.ContentApiFunction.ConstructApiFunction())
	if a.TopComments {
		return fmt.Sprintf("%s/top", base)
	} else {
		if a.Page == 0 {
			return base
		} else {
			return PaginateString(base, a.Page)
		}
	}
}

type ContentDownloadsApiFunction struct {
	*ContentApiFunction
}

func (a *ContentDownloadsApiFunction) ConstructApiFunction() string {
	return fmt.Sprintf("%s/downloads", a.ContentApiFunction.ConstructApiFunction())
}

func ApiCall(url ApiFunction, c interface{}) error {
	resp, err := http.Get(url.ConstructApiFunction())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		return err
	}
	return nil
}

/*
func GetContentInformation(category, name string) (*Content, error) {
	var c Content
	url := AppendApiHeader(fmt.Sprintf("%s/%s", category, name))
	err := ApiCall(url, &c)
	if err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}
func GetContentComments(category, name string) (*Comments, error) {
	var c Comments
	url := AppendApiHeader(fmt.Sprintf("%s/%s/comments", category, name))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
func GetContentCommentsPage(category, name string, page int) (*Comments, error) {
	var c Comments
	url := AppendApiHeader(AppendPagination(fmt.Sprintf("%s/%s/comments", category, name), page))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func GetContentTopComments(category, name string) {

}
*/

func constructAttributeFields(c map[string]interface{}, field string) []*Attribute {
	tmp := c[field].([]interface{})
	size := len(tmp)
	attrs := make([]*Attribute, size)
	for i := 0; i < size; i++ {
		attrs[i] = NewAttribute(tmp[i].(map[string]interface{}))
	}
	return attrs
}

func NewAttribute(c map[string]interface{}) *Attribute {
	return &Attribute{
		Attribute:     c["attribute"].(string),
		AttributeLink: c["attribute_link"].(string),
	}
}

func (a *Attribute) String() string {
	return a.Attribute
}

type Comment struct {
	Id         float64 `json:"comment_id"`
	AttachedId string  `json:"comment_attached_id"`
	Poster     string  `json:"comment_poster"`
	PosterUrl  string  `json:"comment_poster_url"`
	Reputation float64 `json:"comment_reputation"`
	Text       string  `json:"comment_text"`
	Date       float64 `json:"comment_date"`
}
type Comments struct {
	Comments   []*Comment `json:"comments"`
	PageNumber float64    `json:"page"`
	Total      float64    `json:"total"`
	Pages      float64    `json:"pages"`
}
