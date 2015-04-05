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

func getContentCommentsGeneric(url apiFunction) (*Comments, error) {
	var c Comments
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}
func ContentComments(category, name string) (*Comments, error) {
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

func ContentCommentsPage(category, name string, page uint) (*Comments, error) {
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

func ContentTopComments(category, name string) (*Comments, error) {
	url := contentCommentApiFunction{
		contentApiFunction: contentApiFunction{
			Category: category,
			Name:     name,
		},
		TopComments: true,
	}
	return getContentCommentsGeneric(url)
}

func (this *Content) TopComments() (*Comments, error) {
	return ContentTopComments(this.Category, this.rawName)
}
func (this *Content) CommentsPage(page uint) (*Comments, error) {
	return ContentCommentsPage(this.Category, this.rawName, page)
}
func (this *Content) Comments() (*Comments, error) {
	return ContentComments(this.Category, this.rawName)
}

type Comment struct {
	Id         int64   `json:"comment_id"`
	Parent     int64   `json:"comment_attached_id"`
	Reputation float64 `json:"comment_reputation"`
	Text       string  `json:"comment_text"`
	RawDate    int64   `json:"comment_date"`
	Name       string  `json:"comment_poster" json:"comment_content_name"`
	RawUrl     string  `json:"comment_poster_url" json:"comment_content_url"`
}

func (this *Comment) Url() (*url.URL, error) {
	return url.Parse(this.Name)
}
func (this *Comment) Date() time.Time {
	return time.Unix(this.RawDate, 0)
}

/*
func (this *Comment) PosterInformation() (*UserProfile, error) {
	return
}
*/

type CommentList []Comment
type Comments struct {
	Comments   CommentList `json:"comments"`
	PageNumber int         `json:"page"`
	Total      int         `json:"total"`
	Pages      int         `json:"pages"`
}

/*
type UserComment struct {
	Id          uint   `json:"comment_id"`
	AttachedId  uint   `json:"comment_attached_id"`
	Reputation  int    `json:"comment_reputation"`
	Text        string `json:"comment_string"`
	RawDate     int64  `json:"comment_date"`
	ContentName string `json:"comment_content_name"`
	ContentUrl  string `json:"comment_content_url"`
}
*/
//type UserCommentList []UserComment
/*
func (this *UserComment) Url() (*url.URL, error) {
	return url.Parse(this.ContentUrl)
}
func (this *UserComment) Date() time.Time {
	return time.Unix(this.RawDate, 0)
}
*/

type UserComments struct {
	CommentsList CommentList `json:"comments"`
	Total        uint        `json:"total"`
	Pages        uint        `json:"pages"`
}

type userCommentsApiFunction struct {
	userApiFunction
	supportsPagination
}

func (c userCommentsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/comments", c.userApiFunction.Construct())
	return paginateString(base, c.Page)
}

func GetUserCommentsPage(user string, page uint) (*UserComments, error) {
	var c UserComments
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

func GetUserComments(user string) (*UserComments, error) {
	return GetUserCommentsPage(user, 0)
}

func (this *UserProfile) Comments() (*UserComments, error) {
	return GetUserComments(this.Username)
}
func (this *UserProfile) CommentsPage(page uint) (*UserComments, error) {
	return GetUserCommentsPage(this.Username, page)
}
