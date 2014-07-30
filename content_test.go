package fakku

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestContentGetSimple_1(t *testing.T) {
	// Test the manga listed in the API docs
	var c Content
	resp, err := http.Get("https://api.fakku.net/manga/right-now-while-cleaning-the-pool")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Name: %s\n", c.Name)
	t.Logf("Url: %s\n", c.Url)
	t.Logf("Description: %s\n", c.Description)
	t.Log("Tags:")
	for i := 0; i < len(c.Tags); i++ {
		t.Logf("\t- %s", c.Tags[i])
	}
}
