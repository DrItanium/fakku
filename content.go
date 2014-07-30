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

func GetContent(category, name string) (*Content, error) {
	var c Content
	url := fmt.Sprintf("https://api.fakku.net/%s/%s", category, name)
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
	Id         string  `json:"comment_id"`
	AttachedId string  `json:"comment_attached_id"`
	Poster     string  `json:"comment_poster"`
	PosterUrl  string  `json:"comment_poster_url"`
	Reputation float64 `json:"comment_reputation"`
	Text       string  `json:"comment_text"`
	Date       float64 `json:"comment_date"`
}
type Comments struct {
	Comments []*Comment `json:"comments"`
	Total    float64    `json:"total"`
	Pages    float64    `json:"pages"`
}
