package fakku

import (
	"fmt"
	"net/url"
	"time"
)

type userApiFunction struct {
	Name string
}

func (c userApiFunction) Construct() string {
	return fmt.Sprintf("%s/users/%s", apiHeader, c.Name)
}

func GetUser(name string) (*UserProfile, error) {
	var c user
	url := userApiFunction{Name: name}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		//cheat a little bit :D
		return &(c.Profile), nil
	}
}

type user struct {
	Profile UserProfile `json:"user"`
}

// DAMN, I can't just have Go convert these to bools for me.
// I'll need to do the conversion manually
type UserProfile struct {
	Username            string `json:"user_username"`
	RawUrl              string `json:"user_url"`
	Rank                string `json:"user_rank"`
	RawAvatar           string `json:"user_avatar"`
	RawRegistrationDate int64  `json:"user_registration_date"`
	RawLastVisit        int64  `json:"user_last_visit"`
	Subscribed          uint   `json:"user_subscribed"`
	PostCount           uint   `json:"user_posts"`
	Topics              uint   `json:"user_topics"`
	CommentCount        uint   `json:"user_comments"`
	Signature           string `json:"user_signature"`
	ForumReputation     int    `json:"user_forum_reputation"`
	CommentReputation   int    `json:"user_comment_reputation"`
	RawGold             uint   `json:"user_gold"`
}

func (this *UserProfile) Gold() bool {
	return this.RawGold == 1
}
func (this *UserProfile) Url() (*url.URL, error) {
	return url.Parse(this.RawUrl)
}
func (this *UserProfile) RegistrationDate() time.Time {
	return time.Unix(this.RawRegistrationDate, 0)
}

func (this *UserProfile) LastVisit() time.Time {
	return time.Unix(this.RawLastVisit, 0)
}

func (this *UserProfile) Favorites() (*UserFavorites, error) {
	return GetUserFavorites(this.Username)
}
func (this *UserProfile) FavoritesPage(page uint) (*UserFavorites, error) {
	return GetUserFavoritesPage(this.Username, page)
}

func (this *UserProfile) AvatarUrl() (*url.URL, error) {
	return url.Parse(this.RawAvatar)
}

type userFavoritesApiFunction struct {
	userApiFunction
	supportsPagination
}

func (c userFavoritesApiFunction) Construct() string {
	base := fmt.Sprintf("%s/favorites", c.userApiFunction.Construct())
	return paginateString(base, c.Page)
}

func GetUserFavoritesPage(user string, page uint) (*UserFavorites, error) {
	var c UserFavorites
	url := userFavoritesApiFunction{
		userApiFunction:    userApiFunction{Name: user},
		supportsPagination: supportsPagination{Page: page},
	}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetUserFavorites(user string) (*UserFavorites, error) {
	return GetUserFavoritesPage(user, 0)
}

type UserFavorites struct {
	Favorites ContentList `json:"favorites"`
	Total     uint        `json:"total"`
	Pages     uint        `json:"pages"`
}

type userPostsApiFunction struct {
	userApiFunction
	supportsPagination
}

func (c userPostsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/posts", c.userApiFunction.Construct())
	return paginateString(base, c.Page)
}

func GetUserPostsPage(user string, page uint) (*UserPosts, error) {
	var c UserPosts
	url := userPostsApiFunction{
		userApiFunction:    userApiFunction{Name: user},
		supportsPagination: supportsPagination{Page: page},
	}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetUserPosts(user string) (*UserPosts, error) {
	return GetUserPostsPage(user, 0)
}

func (this *UserProfile) Posts() (*UserPosts, error) {
	return GetUserPosts(this.Username)
}
func (this *UserProfile) PostsPage(page uint) (*UserPosts, error) {
	return GetUserPostsPage(this.Username, page)
}

type UserPosts struct {
	Posts UserPostList `json:"posts"`
	Total uint         `json:"total"`
	Pages uint         `json:"pages"`
}

type UserPost struct {
	Id         uint   `json:"post_id"`
	RawDate    int64  `json:"post_date"`
	Text       string `json:"post_text"`
	Reputation int    `json:"post_reputation"`
	TopicTitle string `json:"post_topic_title"`
	TopicUrl   string `json:"post_topic_url"`
}
type UserPostList []UserPost

func (this *UserPost) Url() (*url.URL, error) {
	return url.Parse(this.TopicUrl)
}
func (this *UserPost) Date() time.Time {
	return time.Unix(this.RawDate, 0)
}

type UserTopicsApiFunction struct {
	userApiFunction
	supportsPagination
}

func (c UserTopicsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/topics", c.userApiFunction.Construct())
	return paginateString(base, c.Page)
}

func GetUserTopicsPage(user string, page uint) (*UserTopics, error) {
	var c UserTopics
	url := UserTopicsApiFunction{
		userApiFunction:    userApiFunction{Name: user},
		supportsPagination: supportsPagination{Page: page},
	}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetUserTopics(user string) (*UserTopics, error) {
	return GetUserTopicsPage(user, 0)
}

type UserTopicList []UserTopic
type UserTopics struct {
	Topics UserTopicList `json:"topics"`
	Total  uint          `json:"total"`
	Pages  uint          `json:"pages"`
}

type UserTopic struct {
	Title        string `json:"topic_title"`
	RawUrl       string `json:"topic_url"`
	RawTime      int64  `json:"topic_time"`
	Replies      uint   `json:"topic_replies"`
	Status       uint   `json:"topic_status"`
	Poll         uint   `json:"topic_poll"`
	LastPostId   uint   `json:"topic_last_post_id"`
	PostPreview  string `json:"topic_post_preview"`
	Poster       string `json:"topic_poster"`
	RawPosterUrl string `json:"topic_poster_url"`
}

func (this *UserTopic) Url() (*url.URL, error) {
	return url.Parse(this.RawUrl)
}

func (this *UserTopic) PosterUrl() (*url.URL, error) {
	return url.Parse(this.RawPosterUrl)
}

func (this *UserTopic) Time() time.Time {
	return time.Unix(this.RawTime, 0)
}
