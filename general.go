// General Information
package fakku

type GeneralInformationApiFunction struct{}

func (a GeneralInformationApiFunction) ConstructApiFunction() string {
	return ApiHeader
}

type GeneralInformation struct {
	Title         string `json:"title"`
	Url           string `json:"url"`
	Documentation string `json:"documentation"`
}

func GetGeneralInformation() (*GeneralInformation, error) {
	var c GeneralInformation
	url := GeneralInformationApiFunction{}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}
