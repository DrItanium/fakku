// Categories search functions
package fakku

import (
	"fmt"
)

type ContentSearchResults struct {
	Content ContentList `json:"content"`
	Total   uint        `json:"total"`
	Pages   uint        `json:"pages"`
}

type contentSearchApiFunction struct {
	Terms string
	supportsPagination
}

func (c contentSearchApiFunction) Construct() string {
	base := fmt.Sprintf("%s/search/%s", apiHeader, c.Terms)
	return paginateString(base, c.Page)
}

func ContentSearchPage(terms string, page uint) (*ContentSearchResults, error) {
	var c ContentSearchResults
	url := contentSearchApiFunction{
		Terms:              terms,
		supportsPagination: supportsPagination{Page: page},
	}
	if err := apiCall(url, c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func ContentSearch(terms string) (*ContentSearchResults, error) {
	return ContentSearchPage(terms, 0)
}
