package fakku

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var ApiHeader = "https://api.fakku.net/"

type ApiFunction interface {
	ConstructApiFunction() string
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
	err = json.Unmarshal(body, &c)
	if err != nil {
		return err
	}
	return nil
}
