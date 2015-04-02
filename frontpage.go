package fakku

import (
	"encoding/json"
	"fmt"
)

type FrontPagePostsApiFunction struct {
	SupportsPagination
}

func (c FrontPagePostsApiFunction) Construct() string {
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
			c.Index[i] = newContentFromPopulation(q)
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

type FrontPagePollApiFunction struct{}

func (c FrontPagePollApiFunction) Construct() string {
	return fmt.Sprintf("%s/poll", ApiHeader)
}

func GetFrontPagePoll() (*FrontPagePoll, error) {
	var c FrontPagePoll
	url := FrontPagePollApiFunction{}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type FrontPagePoll struct {
	Question string        `json:"poll_question"`
	Url      string        `json:"poll_url"`
	Options  []*PollOption `json:"poll_options"`
}

type PollOption struct {
	Text  string `json:"option_text"`
	Votes uint   `json:"option_votes"`
}

func (c *FrontPagePoll) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	q := m["poll"].(map[string]interface{})
	c.Question = q["poll_question"].(string)
	c.Url = q["poll_url"].(string)
	tmp := q["poll_options"]
	ops := tmp.([]interface{})
	c.Options = make([]*PollOption, len(ops))
	for i := 0; i < len(ops); i++ {
		lookup := ops[i].(map[string]interface{})
		option := &PollOption{
			Text:  lookup["option_text"].(string),
			Votes: uint(lookup["option_votes"].(float64)),
		}
		c.Options[i] = option
	}
	return nil
}

type FrontPageFeaturedTopicsApiFunction struct{}

func (c FrontPageFeaturedTopicsApiFunction) Construct() string {
	return fmt.Sprintf("%s/featured", ApiHeader)
}

func GetFrontPageFeaturedTopics() (*FrontPageFeaturedTopics, error) {
	var c FrontPageFeaturedTopics
	url := FrontPageFeaturedTopicsApiFunction{}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type FrontPageFeaturedTopics struct {
	Topics []*Topic `json:"topics"`
	Total  uint     `json:"total"`
}
