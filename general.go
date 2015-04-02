// General Information
package fakku

import (
	"net/url"
)

type generalInformationApiFunction struct{}

func (a generalInformationApiFunction) Construct() string {
	return ApiHeader
}

type GeneralInformation struct {
	Title         string `json:"title"`
	RawUrl        string `json:"url"`
	Documentation string `json:"documentation"`
}

func (this *GeneralInformation) Url() (*url.URL, error) {
	return url.Parse(this.RawUrl)
}

func GetGeneralInformation() (*GeneralInformation, error) {
	var c GeneralInformation
	url := generalInformationApiFunction{}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

// checks to see if we can connect to fakku
func Online() bool {
	_, err := GetGeneralInformation()
	return err != nil
}
