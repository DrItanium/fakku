package fakku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ResponseOk                         = 200
	ResponseNotFound                   = 404
	ResponseUnavailableForLegalReasons = 451 // got DMCA'd son
	ResponseServiceUnavailable         = 503 // Information unavailable
	ApiHeader                          = "https://api.fakku.net/"
)

type ApiFunction interface {
	Construct() string
}

type SupportsPagination struct {
	Page uint
}

type ErrorStatus struct {
	ErrorCode    int
	ErrorMessage string `json:"error"`
	KnownError   bool
}

type UnknownEntry struct {
	Message string
}

func (e UnknownEntry) Error() string {
	return e.Message
}

func (e ErrorStatus) Error() string {
	return fmt.Sprintf("Error %d: %s", e.ErrorCode, e.ErrorMessage)
}

func ApiCall(url ApiFunction, c interface{}) error {
	resp, err := http.Get(url.Construct())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case ResponseOk:
		err = json.Unmarshal(body, &c)
		if err != nil {
			return err
		}
		return nil
	case ResponseNotFound, ResponseUnavailableForLegalReasons, ResponseServiceUnavailable:
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
	// If page is zero then it is meaningless so just return the string
	if page == 0 {
		return s
	} else {
		return fmt.Sprintf("%s/page/%d", s, page)
	}
}
