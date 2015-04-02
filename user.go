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
	return fmt.Sprintf("%s/users/%s", ApiHeader, c.Name)
}

func GetUser(name string) (*UserProfile, error) {
	var c user
	url := userApiFunction{Name: name}
	if err := ApiCall(url, &c); err != nil {
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
	Avatar              string `json:"user_avatar"`
	AvatarWidth         uint   `json:"user_avatar_width"`
	AvatarHeight        uint   `json:"user_avatar_height"`
	RawRegistrationDate int64  `json:"user_registration_date"`
	RawLastVisit        int64  `json:"user_last_visit"`
	Subscribed          uint   `json:"user_subscribed"`
	Timezone            int    `json:"user_timezone"`
	PostCount           uint   `json:"user_posts"`
	Topics              uint   `json:"user_topics"`
	CommentCount        uint   `json:"user_comments"`
	Signature           string `json:"user_signature"`
	ForumReputation     int    `json:"user_forum_reputation"`
	CommentReputation   int    `json:"user_comment_reputation"`
	RawGold             uint   `json:"user_gold"`
	RawOnline           uint   `json:"user_online"`
}

func (this *UserProfile) Gold() bool {
	return this.RawGold == 1
}
func (this *UserProfile) Online() bool {
	return this.RawOnline == 1
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

type userFavoritesApiFunction struct {
	userApiFunction
	SupportsPagination
}

func (c userFavoritesApiFunction) Construct() string {
	base := fmt.Sprintf("%s/favorites", c.userApiFunction.Construct())
	return PaginateString(base, c.Page)
}

func GetUserFavoritesPage(user string, page uint) (*UserFavorites, error) {
	var c UserFavorites
	url := userFavoritesApiFunction{
		userApiFunction:    userApiFunction{Name: user},
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
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

type userAchievementsApiFunction struct {
	userApiFunction
}

func (c userAchievementsApiFunction) Construct() string {
	return fmt.Sprintf("%s/achievements", c.userApiFunction.Construct())
}

func GetUserAchievements(user string) (*UserAchievements, error) {
	var c UserAchievements
	url := userAchievementsApiFunction{
		userApiFunction: userApiFunction{Name: user},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type UserAchievements struct {
	Achievements []*UserAchievement `json:"achievements"`
	Total        uint               `json:"total"`
}
type UserAchievement struct {
	Name        string `json:"achievement_name"`
	Description string `json:"achievement_description"`
	Icon        string `json:"achievement_icon"`
	Class       string `json:"achievement_class"`
	RawDate     int64  `json:"achievement_date"`
}

func (this *UserAchievement) Date() time.Time {
	return time.Unix(this.RawDate, 0)
}
func (this *UserAchievement) IconUrl() (*url.URL, error) {
	return url.Parse(this.Icon)
}

type userPostsApiFunction struct {
	userApiFunction
	SupportsPagination
}

func (c userPostsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/posts", c.userApiFunction.Construct())
	return PaginateString(base, c.Page)
}

func GetUserPostsPage(user string, page uint) (*UserPosts, error) {
	var c UserPosts
	url := userPostsApiFunction{
		userApiFunction:    userApiFunction{Name: user},
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
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
	SupportsPagination
}

func (c UserTopicsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/topics", c.userApiFunction.Construct())
	return PaginateString(base, c.Page)
}

func GetUserTopicsPage(user string, page uint) (*UserTopics, error) {
	var c UserTopics
	url := UserTopicsApiFunction{
		userApiFunction:    userApiFunction{Name: user},
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
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

type UserComment struct {
	Id          uint   `json:"comment_id"`
	AttachedId  uint   `json:"comment_attached_id"`
	Reputation  int    `json:"comment_reputation"`
	Text        string `json:"comment_string"`
	RawDate     int64  `json:"comment_date"`
	ContentName string `json:"comment_content_name"`
	ContentUrl  string `json:"comment_content_url"`
}
type UserCommentList []UserComment

func (this *UserComment) Url() (*url.URL, error) {
	return url.Parse(this.ContentUrl)
}
func (this *UserComment) Date() time.Time {
	return time.Unix(this.RawDate, 0)
}

type UserComments struct {
	CommentsList UserCommentList `json:"comments"`
	Total        uint            `json:"total"`
	Pages        uint            `json:"pages"`
}

type userCommentsApiFunction struct {
	userApiFunction
	SupportsPagination
}

func (c userCommentsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/comments", c.userApiFunction.Construct())
	return PaginateString(base, c.Page)
}

func GetUserCommentsPage(user string, page uint) (*UserComments, error) {
	var c UserComments
	url := userCommentsApiFunction{
		userApiFunction:    userApiFunction{Name: user},
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}

}

func GetUserComments(user string) (*UserComments, error) {
	return GetUserCommentsPage(user, 0)
}

func (this *UserProfile) Comments() (*UserComments, error) {
	return GetUserComments(this.Username)
}
func (this *UserProfile) CommentsPage(page uint) (*UserComments, error) {
	return GetUserCommentsPage(this.Username, page)
}
