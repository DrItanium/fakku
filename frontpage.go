package fakku

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type frontPagePostsApiFunction struct {
	supportsPagination
}

func (c frontPagePostsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/index", apiHeader)
	return paginateString(base, c.Page)
}

func GetFrontPagePostsPage(page uint) (*FrontPagePosts, error) {
	var c FrontPagePosts
	url := frontPagePostsApiFunction{
		supportsPagination: supportsPagination{Page: page},
	}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetFrontPage() (*FrontPagePosts, error) {
	return GetFrontPagePostsPage(0)
}

type FrontPageList []interface{}
type FrontPagePosts struct {
	Index FrontPageList `json:"index"`
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
			var z Content
			z.populate(q)
			c.Index[i] = &z
		} else if _, ok := q["topic_title"]; ok {
			var k Topic
			err0 := k.populate(q)
			if err0 != nil {
				return err0
			}
			c.Index[i] = k
		} else {
			return &UnknownEntry{Message: fmt.Sprintf("Couldn't figure out front page entry type %s", q)}
		}
	}
	return nil
}

type frontPagePollApiFunction struct{}

func (c frontPagePollApiFunction) Construct() string {
	return fmt.Sprintf("%s/poll", apiHeader)
}

func GetFrontPagePoll() (*FrontPagePoll, error) {
	var c FrontPagePoll
	url := frontPagePollApiFunction{}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type PollOptionList []PollOption
type FrontPagePoll struct {
	Question string         `json:"poll_question"`
	RawUrl   string         `json:"poll_url"`
	Options  PollOptionList `json:"poll_options"`
}

func (this *FrontPagePoll) Url() (*url.URL, error) {
	return url.Parse(this.RawUrl)
}

type PollOption struct {
	Text  string `json:"option_text"`
	Votes uint   `json:"option_votes"`
}

func (this *PollOption) String() string {
	return fmt.Sprintf("%s - %d", this.Text, this.Votes)
}

func (c *FrontPagePoll) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	q := m["poll"].(map[string]interface{})
	c.Question = q["poll_question"].(string)
	c.RawUrl = q["poll_url"].(string)
	tmp := q["poll_options"]
	ops := tmp.([]interface{})
	c.Options = make([]PollOption, len(ops))
	for i := 0; i < len(ops); i++ {
		lookup := ops[i].(map[string]interface{})
		c.Options[i].Text = lookup["option_text"].(string)
		c.Options[i].Votes = uint(lookup["option_votes"].(float64))
	}
	return nil
}

type frontPageFeaturedTopicsApiFunction struct{}

func (c frontPageFeaturedTopicsApiFunction) Construct() string {
	return fmt.Sprintf("%s/featured", apiHeader)
}

func GetFrontPageFeaturedTopics() (*FrontPageFeaturedTopics, error) {
	var c FrontPageFeaturedTopics
	url := frontPageFeaturedTopicsApiFunction{}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type FrontPageFeaturedTopics struct {
	Topics []Topic `json:"topics"`
	Total  uint    `json:"total"`
}
