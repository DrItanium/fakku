package fakku

import (
	"fmt"
	"log"
	"net/url"
	"time"
)

type contentCommentApiFunction struct {
	contentApiFunction
	supportsPagination
	TopComments bool
}

func (a contentCommentApiFunction) Construct() string {
	base := fmt.Sprintf("%s/comments", a.contentApiFunction.Construct())
	log.Printf("url: %s", base)
	if a.TopComments {
		return fmt.Sprintf("%s/top", base)
	} else {
		return paginateString(base, a.Page)
	}
}

func getContentCommentsGeneric(url apiFunction) (*ContentCommentList, error) {
	var c ContentCommentList
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}
func ContentComments(category, name string) (*ContentCommentList, error) {
	url := contentCommentApiFunction{
		contentApiFunction: contentApiFunction{
			Category: category,
			Name:     name,
		},
		supportsPagination: supportsPagination{Page: 0},
		TopComments:        false,
	}
	return getContentCommentsGeneric(url)
}

func ContentCommentsPage(category, name string, page uint) (*ContentCommentList, error) {
	url := contentCommentApiFunction{
		contentApiFunction: contentApiFunction{
			Category: category,
			Name:     name,
		},
		supportsPagination: supportsPagination{Page: page},
		TopComments:        false,
	}
	return getContentCommentsGeneric(url)
}

func ContentTopComments(category, name string) (*ContentCommentList, error) {
	url := contentCommentApiFunction{
		contentApiFunction: contentApiFunction{
			Category: category,
			Name:     name,
		},
		TopComments: true,
	}
	return getContentCommentsGeneric(url)
}

func (this *Content) TopComments() (*ContentCommentList, error) {
	return ContentTopComments(this.Category, this.rawName)
}
func (this *Content) CommentsPage(page uint) (*ContentCommentList, error) {
	return ContentCommentsPage(this.Category, this.rawName, page)
}
func (this *Content) Comments() (*ContentCommentList, error) {
	return ContentComments(this.Category, this.rawName)
}

// Most of the data returned from the fakku rest api are statements made by users in some sort of aggregation
// A statement can be a post, comment, or topic
// An aggregation can be a topic, content, user, forum,
//
// The problem is that there are usually two aggregators (ways) which can be used to get to the same core statement
// The difference is the link that refers to the other aggregate embedded in the structure.
//
// For example, there are two types of comment: user-aggregate comment and content-aggregate comment
// the actual content (minus the other aggregate reference) is the same between the two types.
//
// In real terms this means:
// 		- a comment from a content-aggregate has the poster information in it (user-aggregate)
// 		- a comment from a user-aggregate has the content information in it (content-aggregate)
//
// I view this as each statement having two aggregate references to it.
// - Comment has references to: User and Content
// - Post has references to: User and Topic
// - Topic has reference to: User and Forum
//
// Notice that each statement has a reference to the user which made the statement
// The other aggregate is where the statement is located (Content, Topic, Forum)
//
// This implies that a User is really just an aggregator of different types of statements
// A "content" is just an aggregate of Statements made by users
// A topic is just an aggregate of Statements made by users
// A forum is just an aggregate of Statements made by users
//
// While these statements are different the concept is still the same across all types of aggregates
//
// When dealing with the fact that there are two aggregates for each statement
// we have to remain cognizant of the "perspective" of how we got to that
// statement and where we can go and where we can go from that statement.
//
// All of this is leading me to believe that it isn't easy to merge UserComments and ContentComments
// without losing some sort of perspective of where we came from.
type CommentBody struct {
	Id         int64   `json:"comment_id"`
	Parent     int64   `json:"comment_attached_id"`
	Reputation float64 `json:"comment_reputation"`
	Text       string  `json:"comment_text"`
	RawDate    int64   `json:"comment_date"`
}

func (this *CommentBody) Date() time.Time {
	return time.Unix(this.RawDate, 0)
}

type ContentComment struct {
	CommentBody
	Poster string `json:"comment_poster_name"`
	RawUrl string `json:"comment_poster_url"`
}

func (this *ContentComment) PosterUrl() (*url.URL, error) {
	return url.Parse(this.RawUrl)
}

type UserComment struct {
	CommentBody
	Content string `json:"comment_content_name"`
	RawUrl  string `json:"comment_content_url"`
}

func (this *UserComment) ContentUrl() (*url.URL, error) {
	return url.Parse(this.RawUrl)
}

type ContentCommentList struct {
	Comments   []ContentComment `json:"comments"`
	PageNumber int              `json:"page"`
	Total      int              `json:"total"`
	Pages      int              `json:"pages"`
}

type UserCommentList struct {
	Comments []UserComment `json:"comments"`
	Total    uint          `json:"total"`
	Pages    uint          `json:"pages"`
}

type userCommentsApiFunction struct {
	userApiFunction
	supportsPagination
}

func (c userCommentsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/comments", c.userApiFunction.Construct())
	return paginateString(base, c.Page)
}

func GetUserCommentsPage(user string, page uint) (*UserCommentList, error) {
	var c UserCommentList
	url := userCommentsApiFunction{
		userApiFunction:    userApiFunction{Name: user},
		supportsPagination: supportsPagination{Page: page},
	}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}

}

func GetUserComments(user string) (*UserCommentList, error) {
	return GetUserCommentsPage(user, 0)
}

func (this *UserProfile) Comments() (*UserCommentList, error) {
	return GetUserComments(this.Username)
}
func (this *UserProfile) CommentsPage(page uint) (*UserCommentList, error) {
	return GetUserCommentsPage(this.Username, page)
}
