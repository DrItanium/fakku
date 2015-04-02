package fakku

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	ErrorContentDoesntExist = "Content doesn't exist"
	ErrorUnknownJsonData    = "Got unknown json data back from content request. API Change?"
	ErrorUnknownJsonLayout  = "Got an unknown layout back from content request. API Change?"
)

type Attribute struct {
	Attribute     string `json:"attribute"`
	AttributeLink string `json:"attribute_link"`
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
	Tags        []*Attribute `json:"content_tags"`
	Translators []*Attribute `json:"content_translators"`
	Series      []*Attribute `json:"content_series"`
	Artists     []*Attribute `json:"content_artists"`
	Images      struct {
		Cover  string
		Sample string
	}
}

func newContentFromPopulation(v map[string]interface{}) *Content {
	var c Content
	c.populateContent(v)
	return &c
}

func (c *Content) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)
	switch f.(type) {
	case map[string]interface{}:
		m := f.(map[string]interface{})
		contents := m["content"]
		switch contents.(type) {
		case map[string]interface{}:
			v := contents.(map[string]interface{})
			c.populateContent(v)
			return nil
		case []interface{}:
			q := contents.([]interface{})
			if len(q) == 0 {
				// doesn't exist
				return fmt.Errorf(ErrorContentDoesntExist)
			} else {
				return fmt.Errorf(ErrorUnknownJsonData)
			}
		default:
			return fmt.Errorf(ErrorUnknownJsonLayout)
		}
	case []interface{}:
		q := f.([]interface{})
		if len(q) == 0 {
			// doesn't exist
			return fmt.Errorf(ErrorContentDoesntExist)
		} else {
			return fmt.Errorf(ErrorUnknownJsonData)
		}
	default:
		return fmt.Errorf(ErrorUnknownJsonLayout)
	}
}

func (c *Content) populateContent(v map[string]interface{}) {
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
}

func constructAttributeFields(c map[string]interface{}, field string) []*Attribute {
	try, ok := c[field]
	if !ok {
		return nil
	}
	tmp := try.([]interface{})
	size := len(tmp)
	attrs := make([]*Attribute, size)
	for i := 0; i < size; i++ {
		attrs[i] = NewAttribute(tmp[i].(map[string]interface{}))
	}
	return attrs
}

type ContentApiFunction struct {
	Category string
	Name     string
}

func (a ContentApiFunction) ConstructApiFunction() string {
	return fmt.Sprintf("%s/%s/%s", ApiHeader, a.Category, a.Name)
}

type ContentCommentApiFunction struct {
	ContentApiFunction
	TopComments bool
	SupportsPagination
}

func (a ContentCommentApiFunction) ConstructApiFunction() string {
	base := fmt.Sprintf("%s/comments", a.ContentApiFunction.ConstructApiFunction())
	if a.TopComments {
		return fmt.Sprintf("%s/top", base)
	} else {
		return PaginateString(base, a.Page)
	}
}

func GetContentInformation(category, name string) (*Content, error) {
	var c Content
	url := ContentApiFunction{Category: category, Name: name}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}
func getContentCommentsGeneric(url ApiFunction) (*Comments, error) {
	var c Comments
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}
func GetContentComments(category, name string) (*Comments, error) {
	url := ContentCommentApiFunction{
		ContentApiFunction: ContentApiFunction{
			Category: category,
			Name:     name,
		},
	}
	return getContentCommentsGeneric(url)
}

func GetContentCommentsPage(category, name string, page uint) (*Comments, error) {
	url := ContentCommentApiFunction{
		ContentApiFunction: ContentApiFunction{
			Category: category,
			Name:     name,
		},
		SupportsPagination: SupportsPagination{Page: page},
	}
	return getContentCommentsGeneric(url)
}

func GetContentTopComments(category, name string) (*Comments, error) {
	url := ContentCommentApiFunction{
		ContentApiFunction: ContentApiFunction{
			Category: category,
			Name:     name,
		},
		TopComments: true,
	}
	return getContentCommentsGeneric(url)
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

func NewAttribute(c map[string]interface{}) *Attribute {
	return &Attribute{
		Attribute:     c["attribute"].(string),
		AttributeLink: c["attribute_link"].(string),
	}
}

func (a *Attribute) String() string {
	return a.Attribute
}

type Comments struct {
	Comments   []*Comment `json:"comments"`
	PageNumber float64    `json:"page"`
	Total      float64    `json:"total"`
	Pages      float64    `json:"pages"`
}

type ReadOnlineContent struct {
	Content *Content `json:"content"`
	Pages   []*Page  `json:"pages"`
}

func (r *ReadOnlineContent) UnmarshalJSON(b []byte) error {
	var f interface{}
	if err := json.Unmarshal(b, &r.Content); err != nil {
		return err
	}
	json.Unmarshal(b, &f)
	// need to check and make sure that the content exists
	switch f.(type) {
	case map[string]interface{}:
		m := f.(map[string]interface{})
		pages := m["pages"]
		v := pages.(map[string]interface{})
		r.Pages = make([]*Page, len(v))
		for i := 0; i < len(v); i++ {
			ind := strconv.Itoa(i + 1)
			r.Pages[i] = NewPage(ind, v[ind].(map[string]interface{}))
		}
		return nil
	case []interface{}:
		q := f.([]interface{})
		if len(q) == 0 {
			// doesn't exist
			return fmt.Errorf("Content doesn't exist")
		} else {
			return fmt.Errorf("Got unknown json data back from content request. API Change?")
		}
	default:
		return fmt.Errorf("Got an unknown layout back from content request. API Change?")
	}
}

type Page struct {
	Id    string
	Thumb string
	Image string
}

func NewPage(id string, c map[string]interface{}) *Page {
	return &Page{
		Id:    id,
		Thumb: c["thumb"].(string),
		Image: c["image"].(string),
	}
}

type ContentReadOnlineApiFunction struct {
	ContentApiFunction
}

func (a ContentReadOnlineApiFunction) ConstructApiFunction() string {
	return fmt.Sprintf("%s/read", a.ContentApiFunction.ConstructApiFunction())
}

func GetContentReadOnline(category, name string) (*ReadOnlineContent, error) {
	var c ReadOnlineContent
	url := ContentReadOnlineApiFunction{
		ContentApiFunction: ContentApiFunction{
			Category: category,
			Name:     name,
		},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetContentDownloads(category, name string) (*DownloadContent, error) {
	var c DownloadContent
	url := ContentDownloadsApiFunction{
		ContentApiFunction: ContentApiFunction{
			Category: category,
			Name:     name,
		},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type ContentDownloadsApiFunction struct {
	ContentApiFunction
}

func (a ContentDownloadsApiFunction) ConstructApiFunction() string {
	return fmt.Sprintf("%s/download", a.ContentApiFunction.ConstructApiFunction())
}

type DownloadContent struct {
	Downloads []*Download `json:"downloads"`
	Total     uint        `json:"total"`
}

func (this *DownloadContent) HasDownloads() bool {
	return this.Total > 0
}

type Download struct {
	Type          string  `json:"download_type"`
	Url           string  `json:"download_url"`
	Info          string  `json:"download_info"`
	DownloadCount float64 `json:"download_count"`
	Time          float64 `json:"download_time"`
	Poster        string  `json:"download_poster"`
	PosterUrl     string  `json:"download_poster_url"`
}

type ContentRelatedApiFunction struct {
	ContentApiFunction
	SupportsPagination
}

func (a ContentRelatedApiFunction) ConstructApiFunction() string {
	base := fmt.Sprintf("%s/related", a.ContentApiFunction.ConstructApiFunction())
	return PaginateString(base, a.Page)
}

type RelatedContent struct {
	Related []*Content `json:"related"`
	Total   uint       `json:"total"`
	Pages   uint       `json:"pages"`
}

func GetRelatedContentAll(category, name string) (*RelatedContent, error) {
	return GetRelatedContent(category, name, 0)
}

func GetRelatedContent(category, name string, page uint) (*RelatedContent, error) {
	var c RelatedContent
	url := ContentRelatedApiFunction{
		ContentApiFunction: ContentApiFunction{
			Category: category,
			Name:     name,
		},
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func (c *RelatedContent) UnmarshalJSON(b []byte) error {
	// slightly different
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	related := m["related"]
	v := related.([]interface{})
	c.Related = make([]*Content, len(v))
	for i := 0; i < len(v); i++ {
		c.Related[i] = newContentFromPopulation(v[i].(map[string]interface{}))
	}
	c.Total = uint(m["total"].(float64))
	c.Pages = uint(m["pages"].(float64))
	return nil
}
