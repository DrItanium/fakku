package fakku

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestContentGetSimple_1(t *testing.T) {
	// Test the manga listed in the API docs
	compareName := "Right now, while cleaning the pool"
	compareUrl := "http://www.fakku.net/manga/right-now-while-cleaning-the-pool"
	compareTag := "Vanilla"
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
	if c.Name != compareName {
		t.Errorf("Expected: %s, Got: %s", compareName, c.Name)
	}
	if c.Url != compareUrl {
		t.Errorf("Expected: %s, Got: %s", compareUrl, c.Url)
	}
	foundTag := false
	for i := 0; i < len(c.Tags); i++ {
		if c.Tags[i].Attribute == compareTag {
			foundTag = true
			break
		}
	}
	if !foundTag {
		t.Errorf("Didn't find tag %s but should have!", compareTag)
	}
}

func TestContentGetSimple_2(t *testing.T) {
	// Test the manga listed in the API docs
	compareName := "Right now, while cleaning the pool"
	compareUrl := "http://www.fakku.net/manga/right-now-while-cleaning-the-pool"
	compareTag := "Vanilla"
	c, err := GetContentInformation("manga", "right-now-while-cleaning-the-pool")
	if err != nil {
		t.Error(err)
	}
	if c.Name != compareName {
		t.Errorf("Expected: %s, Got: %s", compareName, c.Name)
	}
	if c.Url != compareUrl {
		t.Errorf("Expected: %s, Got: %s", compareUrl, c.Url)
	}
	foundTag := false
	for i := 0; i < len(c.Tags); i++ {
		if c.Tags[i].Attribute == compareTag {
			foundTag = true
			break
		}
	}
	if !foundTag {
		t.Errorf("Didn't find tag %s but should have!", compareTag)
	}
}

func TestContentGetComments_1(t *testing.T) {
	var c Comments
	resp, err := http.Get("https://api.fakku.net/manga/right-now-while-cleaning-the-pool/comments")
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
}

func TestContentGetComments_2(t *testing.T) {
	_, err := GetContentComments("manga", "right-now-while-cleaning-the-pool")
	if err != nil {
		t.Error(err)
	}
}
