package fakku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var ApiHeader = "https://api.fakku.net/"

type ApiFunction interface {
	ConstructApiFunction() string
}
type ErrorStatus struct {
	ErrorCode    int
	ErrorMessage string `json:"error"`
	KnownError   bool
}

func (e ErrorStatus) Error() string {
	return fmt.Sprintf("Error %d: %s", e.ErrorCode, e.ErrorMessage)
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
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(body, &c)
		if err != nil {
			return err
		}
		return nil
	case 404, 451, 503:
		// right now just harvest the error code
		var ec ErrorStatus
		err = json.Unmarshal(body, &ec)
		ec.KnownError = true
		ec.ErrorCode = resp.StatusCode
		if err != nil {
			return err
		} else {
			return &ec
		}
	default:
		return &ErrorStatus{ErrorCode: resp.StatusCode, ErrorMessage: resp.Status, KnownError: false}
	}
}

func PaginateString(s string, page uint) string {
	return fmt.Sprintf("%s/page/%d", s, page)
}
