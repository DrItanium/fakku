package fakku

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (
	ErrorContentDoesntExist = "Content doesn't exist"
	ErrorUnknownJsonData    = "Got unknown json data back from content request. API Change?"
	ErrorUnknownJsonLayout  = "Got an unknown layout back from content request. API Change?"
)

type AttributeList []Attribute
type Attribute struct {
	Attribute     string `json:"attribute"`
	AttributeLink string `json:"attribute_link"`
}

func (this *Attribute) populate(c map[string]interface{}) {
	this.Attribute = c["attribute"].(string)
	this.AttributeLink = c["attribute_link"].(string)
}

func (this *Attribute) String() string {
	return this.Attribute
}

func constructAttributeFields(c map[string]interface{}, field string) AttributeList {
	try, ok := c[field]
	if !ok {
		return nil
	}
	tmp := try.([]interface{})
	attrs := make(AttributeList, len(tmp))
	for i := 0; i < len(tmp); i++ {
		attrs[i].populate(tmp[i].(map[string]interface{}))
	}
	return attrs
}

type Content struct {
	Name         string
	RawUrl       string
	Description  string
	Language     string
	Category     string
	RawDate      int64
	FileSize     int64
	Favorites    int64
	CommentCount int64
	Pages        int64
	Poster       string
	RawPosterUrl string
	Tags         AttributeList `json:"content_tags"`
	Translators  AttributeList `json:"content_translators"`
	Series       AttributeList `json:"content_series"`
	Artists      AttributeList `json:"content_artists"`
	Images       struct {
		Cover  string
		Sample string
	}
	rawName string
}
type ContentList []Content

func (this *Content) Url() (*url.URL, error) {
	return url.Parse(https + this.RawUrl)
}
func (this *Content) PosterUrl() (*url.URL, error) {
	return url.Parse(https + this.RawPosterUrl)
}
func (this *Content) CoverUrl() (*url.URL, error) {
	return url.Parse(https + this.Images.Cover)
}
func (this *Content) SampleUrl() (*url.URL, error) {
	return url.Parse(https + this.Images.Sample)
}
func (this *Content) Date() time.Time {
	return time.Unix(this.RawDate, 0)
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
			c.populate(v)
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

func (c *Content) populate(v map[string]interface{}) {
	c.Name = v["content_name"].(string)
	c.RawUrl = v["content_url"].(string)
	c.Description = v["content_description"].(string)
	c.Language = v["content_language"].(string)
	c.Category = v["content_category"].(string)
	c.RawDate = int64(v["content_date"].(float64))
	c.FileSize = int64(v["content_filesize"].(float64))
	c.Favorites = int64(v["content_favorites"].(float64))
	c.CommentCount = int64(v["content_comments"].(float64))
	c.Pages = int64(v["content_pages"].(float64))
	c.Poster = v["content_poster"].(string)
	c.RawPosterUrl = v["content_poster_url"].(string)
	c.Tags = constructAttributeFields(v, "content_tags")
	c.Translators = constructAttributeFields(v, "content_translators")
	c.Series = constructAttributeFields(v, "content_series")
	c.Artists = constructAttributeFields(v, "content_artists")

	tmp := v["content_images"]
	z := tmp.(map[string]interface{})
	c.Images.Cover = z["cover"].(string)
	c.Images.Sample = z["sample"].(string)
}

func GetContent(category, name string) (*Content, error) {
	var c Content
	url := contentApiFunction{Category: category, Name: name}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		c.rawName = name
		return &c, nil
	}
}

type contentApiFunction struct {
	Category string
	Name     string
}

func (a contentApiFunction) Construct() string {
	return fmt.Sprintf("%s/%s/%s", apiHeader, a.Category, a.Name)
}

type PageList []Page
type ReadOnlineContent struct {
	Content Content  `json:"content"`
	Pages   PageList `json:"pages"`
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
		r.Pages = make([]Page, len(v))
		for i := 0; i < len(v); i++ {
			ind := strconv.Itoa(i + 1)
			r.Pages[i].populate(v[ind].(map[string]interface{}))
		}
		return nil
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

type Page struct {
	Thumb string
	Image string
}

func (this *Page) populate(c map[string]interface{}) {
	this.Thumb = c["thumb"].(string)
	this.Image = c["image"].(string)
}

func (this *Page) ThumbUrl() (*url.URL, error) {
	return url.Parse(this.Thumb)
}
func (this *Page) ImageUrl() (*url.URL, error) {
	return url.Parse(this.Image)
}
func (this *Page) ImageBytes() ([]byte, error) {
	url, err := this.ImageUrl()
	if err != nil {
		return nil, err
	}
	return requestBytes(url)
}
func (this *Page) ThumbBytes() ([]byte, error) {
	url, err := this.ThumbUrl()
	if err != nil {
		return nil, err
	}
	return requestBytes(url)
}

type contentReadOnlineApiFunction struct {
	contentApiFunction
}

func (a contentReadOnlineApiFunction) Construct() string {
	return fmt.Sprintf("%s/read", a.contentApiFunction.Construct())
}

func ReadOnline(category, name string) (*ReadOnlineContent, error) {
	var c ReadOnlineContent
	url := contentReadOnlineApiFunction{
		contentApiFunction: contentApiFunction{
			Category: category,
			Name:     name,
		},
	}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func (this *Content) ReadOnline() (PageList, error) {
	element, err := ReadOnline(this.Category, this.rawName)
	if err != nil {
		return nil, err
	} else {
		return element.Pages, nil
	}
}

type contentRelatedApiFunction struct {
	contentApiFunction
	supportsPagination
}

func (a contentRelatedApiFunction) Construct() string {
	base := fmt.Sprintf("%s/related", a.contentApiFunction.Construct())
	return paginateString(base, a.Page)
}

type RelatedContentList struct {
	Related ContentList `json:"related"`
	Total   uint        `json:"total"`
	Pages   uint        `json:"pages"`
}

func RelatedContent(category, name string) (*RelatedContentList, error) {
	return RelatedContentPage(category, name, 0)
}

func RelatedContentPage(category, name string, page uint) (*RelatedContentList, error) {
	var c RelatedContentList
	url := contentRelatedApiFunction{
		contentApiFunction: contentApiFunction{
			Category: category,
			Name:     name,
		},
		supportsPagination: supportsPagination{Page: page},
	}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}
func (this *Content) RelatedContent() (*RelatedContentList, error) {
	return RelatedContent(this.Category, this.rawName)
}
func (this *Content) RelatedContentPage(page uint) (*RelatedContentList, error) {
	return RelatedContentPage(this.Category, this.rawName, page)
}

func (c *RelatedContentList) UnmarshalJSON(b []byte) error {
	// slightly different
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	related := m["related"]
	v := related.([]interface{})
	c.Related = make(ContentList, len(v))
	for i := 0; i < len(v); i++ {
		c.Related[i].populate(v[i].(map[string]interface{}))
	}
	c.Total = uint(m["total"].(float64))
	c.Pages = uint(m["pages"].(float64))
	return nil
}
