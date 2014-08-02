package fakku

import (
	"encoding/json"
	"fmt"
)

type FrontPagePostsApiFunction struct {
	SupportsPagination
}

func (c FrontPagePostsApiFunction) ConstructApiFunction() string {
	base := fmt.Sprintf("%s/index", ApiHeader)
	return PaginateString(base, c.Page)
}

func GetFrontPagePostsPage(page uint) (*FrontPagePosts, error) {
	var c FrontPagePosts
	url := FrontPagePostsApiFunction{
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetFrontPage() (*FrontPagePosts, error) {
	return GetFrontPagePostsPage(0)
}

type FrontPagePosts struct {
	Index []interface{} `json:"index"`
	Total uint          `json:"total"`
}

func (c *FrontPagePosts) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	c.Total = uint(m["total"].(float64))
	c.Index = make([]interface{}, c.Total)

	contents := m["index"]
	v := contents.([]interface{})
	for i := 0; i < len(c.Index); i++ {
		q := v[i].(map[string]interface{})
		if _, ok := q["content_name"]; ok {
			var j Content
			j.populateContent(q)
			c.Index[i] = j
		} else if _, ok := q["topic_title"]; ok {
			var k Topic
			k.populateTopic(q)
			c.Index[i] = k
		} else {
			return &UnknownEntry{Message: fmt.Sprintf("Couldn't figure out front page entry type %s", q)}
		}
	}
	return nil
}
