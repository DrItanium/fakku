// download related operations
package fakku

import (
	"fmt"
	"net/url"
	"time"
)

type contentDownloadsApiFunction struct {
	contentApiFunction
}

func (a contentDownloadsApiFunction) Construct() string {
	return fmt.Sprintf("%s/download", a.contentApiFunction.Construct())
}

type DownloadList []Download

func (this DownloadList) HasDownloads() bool {
	return len(this) > 0
}

type downloadContent struct {
	Downloads DownloadList `json:"downloads"`
	Total     uint         `json:"total"`
}

func ContentDownloads(category, name string) (DownloadList, error) {
	var c downloadContent
	url := contentDownloadsApiFunction{
		contentApiFunction: contentApiFunction{
			Category: category,
			Name:     name,
		},
	}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return c.Downloads, nil
	}
}

func (this *Content) Downloads() (DownloadList, error) {
	return ContentDownloads(this.Category, this.rawName)
}

type Download struct {
	Type          string  `json:"download_type"`
	RawUrl        string  `json:"download_url"`
	Info          string  `json:"download_info"`
	DownloadCount float64 `json:"download_count"`
	RawTime       float64 `json:"download_time"`
	Poster        string  `json:"download_poster"`
	RawPosterUrl  string  `json:"download_poster_url"`
}

func (this *Download) Url() (*url.URL, error) {
	return url.Parse(this.RawUrl)
}
func (this *Download) PosterUrl() (*url.URL, error) {
	return url.Parse(this.RawPosterUrl)
}
func (this *Download) Time() time.Time {
	return time.Unix(int64(this.RawTime), 0)
}
