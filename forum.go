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
	// There are more....
}
