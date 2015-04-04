// tag related functions
package fakku

import (
	"fmt"
	"net/url"
)

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
