// operations related to page interaction
package fakku

import (
	"image"
	"io/ioutil"
	"net/url"
	"os"
)

type PageList []Page
type Page struct {
	rawThumb string
	rawImage string
}

func (this *Page) populate(c map[string]interface{}) {
	this.rawThumb = c["thumb"].(string)
	this.rawImage = c["image"].(string)
}
func (this *Page) ThumbUrl() (*url.URL, error) {
	return url.Parse(this.rawThumb)
}
func (this *Page) ImageUrl() (*url.URL, error) {
	return url.Parse(this.rawImage)
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

func (this *Page) SaveImage(path string, perms os.FileMode) error {
	img, ierr := this.ImageBytes()
	if ierr != nil {
		return ierr
	}
	return ioutil.WriteFile(path, img, perms)
}

func (this *Page) SaveThumb(path string, perms os.FileMode) error {
	img, ierr := this.ThumbBytes()
	if ierr != nil {
		return ierr
	}
	return ioutil.WriteFile(path, img, perms)
}

func (this *Page) Image() (image.Image, error) {
	url, err := this.ImageUrl()
	if err != nil {
		return nil, err
	}
	return requestJpeg(url)
}
func (this *Page) Thumbnail() (image.Image, error) {
	url, err := this.ThumbUrl()
	if err != nil {
		return nil, err
	}
	return requestJpeg(url)
}
