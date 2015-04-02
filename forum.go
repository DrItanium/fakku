package fakku

import (
	"fmt"
	"net/url"
	"time"
)

type forumCategoriesApiFunction struct{}

func (c forumCategoriesApiFunction) Construct() string {
	return fmt.Sprintf("%s/forums", ApiHeader)
}

func GetForumCategories() ([]ForumCategory, error) {
	var c forumCategoriesContainer
	url := forumCategoriesApiFunction{}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return c.Categories, nil
	}
}

type forumCategoriesContainer struct {
	Categories []ForumCategory `json:"categories"`
}
type ForumCategory struct {
	Title  string  `json:"category_title"`
	Order  uint    `json:"category_order"`
	Forums []Forum `json:"forums"`
}

type Forum struct {
	Name        string  `json:"forum_name"`
	Description string  `json:"forum_description"`
	RawUrl      string  `json:"forum_url"`
	Posts       uint    `json:"forum_posts"`
	TopicCount  uint    `json:"forum_topics"`
	Silent      uint    `json:"forum_silent"`
	Total       uint    `json:"total"`
	Page        uint    `json:"page"`
	PageCount   uint    `json:"pages"`
	Topics      []Topic `json:"topics"`
	RecentTopic Topic   `json:"forum_recent_topic"` // this is used in some calls but not others :/
}

func (this *Forum) Url() (*url.URL, error) {
	return url.Parse(this.RawUrl)
}
func (this *Forum) IsSlient() bool {
	return this.Silent == 1
}

type Topic struct {
	Title        string `json:"topic_title"`
	RawUrl       string `json:"topic_url"`
	TopicTime    uint   `json:"topic_time"`
	FirstPostId  uint   `json:"topic_first_post_id"`
	LastPostId   uint   `json:"topic_last_post_id"`
	FrontPage    uint   `json:"front_page"`
	Status       uint   `json:"topic_status"`
	Vote         uint   `json:"topic_vote"`
	Type         uint   `json:"topic_type"`
	Poster       string `json:"topic_poster"`
	RawPosterUrl string `json:"topic_poster_url"`
	RawTime      int64  `json:"topic_url"`
}

func (this *Topic) OnFrontPage() bool {
	return this.FrontPage == 1
}
func (this *Topic) Url() (*url.URL, error) {
	return url.Parse(this.RawUrl)
}

func (this *Topic) PosterUrl() (*url.URL, error) {
	return url.Parse(this.RawPosterUrl)
}

func (this *Topic) Time() time.Time {
	return time.Unix(this.RawTime, 0)
}

func (t *Topic) populateTopic(c map[string]interface{}) error {
	t.Title = c["topic_title"].(string)
	t.RawUrl = c["topic_url"].(string)
	t.RawTime = int64(c["topic_time"].(float64))
	t.FirstPostId = uint(c["topic_first_post_id"].(float64))
	t.LastPostId = uint(c["topic_last_post_id"].(float64))
	if _, ok := c["front_page"]; ok {
		t.FrontPage = uint(c["front_page"].(float64))
	}
	t.Status = uint(c["topic_status"].(float64))
	if _, ok := c["topic_vote"]; ok {
		t.Vote = uint(c["topic_vote"].(float64))
	}
	t.Type = uint(c["topic_type"].(float64))
	t.Poster = c["topic_poster"].(string)
	t.RawPosterUrl = c["topic_poster_url"].(string)
	return nil
}

type ForumTopicsApiFunction struct {
	Forum string
	SupportsPagination
}

func (c ForumTopicsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/forums/%s", ApiHeader, c.Forum)
	return PaginateString(base, c.Page)
}
func GetForumTopics(forum string) (*ForumTopics, error) {
	return GetForumTopicsPage(forum, 0)
}
func GetForumTopicsPage(forum string, page uint) (*ForumTopics, error) {
	var c ForumTopics
	url := ForumTopicsApiFunction{
		Forum:              forum,
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type ForumTopics struct {
	Forum  *Forum   `json:"forum"`
	Topics []*Topic `json:"topics"`
	Total  uint     `json:"total"`
	Page   uint     `json:"page"`
	Pages  uint     `json:"pages"`
}

type ForumPost struct {
	Id         uint         `json:"post_id"`
	Date       uint         `json:"post_date"`
	Poster     string       `json:"post_poster"`
	PosterUrl  string       `json:"post_poster_url"`
	Text       string       `json:"post_text"`
	Image      string       `json:"post_image"`
	Thumb      string       `json:"post_thumb"`
	Reputation int          `json:"post_reputation"`
	User       *UserProfile `json:"post_user"`
}

type ForumPosts struct {
	Topic *Topic       `json:"topic"`
	Forum *Forum       `json:"forum"`
	Posts []*ForumPost `json:"posts"`
	Total uint         `json:"total"`
	Page  uint         `json:"page"`
	Pages uint         `json:"pages"`
}

type ForumPostsApiFunction struct {
	Forum string
	Topic string
	SupportsPagination
}

func (c ForumPostsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/forums/%s/%s", ApiHeader, c.Forum, c.Topic)
	return PaginateString(base, c.Page)
}
func GetForumPosts(forum, topic string) (*ForumPosts, error) {
	return GetForumPostsPage(forum, topic, 0)
}
func GetForumPostsPage(forum, topic string, page uint) (*ForumPosts, error) {
	var c ForumPosts
	url := ForumPostsApiFunction{
		Forum:              forum,
		Topic:              topic,
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}
