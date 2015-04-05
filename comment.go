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
	Id           int64   `json:"comment_id"`
	AttachedId   int64   `json:"comment_attached_id"`
	Poster       string  `json:"comment_poster"`
	RawPosterUrl string  `json:"comment_poster_url"`
	Reputation   float64 `json:"comment_reputation"`
	Text         string  `json:"comment_text"`
	RawDate      int64   `json:"comment_date"`
}

func (this *Comment) PosterUrl() (*url.URL, error) {
	return url.Parse(this.RawPosterUrl)
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
