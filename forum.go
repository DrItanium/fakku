package fakku

import (
	"fmt"
)

type ForumCategoriesApiFunction struct{}

func (c ForumCategoriesApiFunction) ConstructApiFunction() string {
	return fmt.Sprintf("%s/forums", ApiHeader)
}

func GetForumCategories() (*ForumCategories, error) {
	var c ForumCategories
	url := ForumCategoriesApiFunction{}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type ForumCategories struct {
	Categories []*ForumCategory `json:"categories"`
}
type ForumCategory struct {
	Title  string   `json:"category_title"`
	Order  string   `json:"category_order"`
	Forums []*Forum `json:"forums"`
}

type Forum struct {
	Name        string `json:"forum_name"`
	Description string `json:"forum_description"`
	Url         string `json:"forum_url"`
	Posts       uint   `json:"forum_posts"`
	Topics      uint   `json:"forum_topics"`
	Silent      uint   `json:"forum_silent"`
	RecentTopic *Topic `json:"forum_recent_topic"`
	// There are more....
}

type Topic struct {
	Title       string `json:"topic_title"`
	Url         string `json:"topic_url"`
	Time        uint   `json:"topic_time"`
	FirstPostId uint   `json:"topic_first_post_id"`
	LastPostId  uint   `json:"topic_last_post_id"`
	FrontPage   uint   `json:"front_page"`
	Status      uint   `json:"topic_status"`
	Vote        uint   `json:"topic_vote"`
	Type        uint   `json:"topic_type"`
	Poster      string `json:"topic_poster"`
	PosterUrl   string `json:"topic_poster_url"`
}

type ForumTopicsApiFunction struct {
	Forum string
	Page  uint
}

func (c ForumTopicsApiFunction) ConstructApiFunction() string {
	base := fmt.Sprintf("%s/forums/%s", ApiHeader, c.Forum)
	return PaginateString(base, c.Page)
}
func GetForumTopics(forum string) (*ForumTopics, error) {
	return GetForumTopicsPage(forum, 0)
}
func GetForumTopicsPage(forum string, page uint) (*ForumTopics, error) {
	var c ForumTopics
	url := ForumTopicsApiFunction{Forum: forum, Page: page}
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
	Page  uint
}

func (c ForumPostsApiFunction) ConstructApiFunction() string {
	base := fmt.Sprintf("%s/forums/%s/%s", ApiHeader, c.Forum, c.Topic)
	return PaginateString(base, c.Page)
}
func GetForumPosts(forum, topic string) (*ForumPosts, error) {
	return GetForumPostsPage(forum, topic, 0)
}
func GetForumPostsPage(forum, topic string, page uint) (*ForumPosts, error) {
	var c ForumPosts
	url := ForumPostsApiFunction{Forum: forum, Topic: topic, Page: page}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}
