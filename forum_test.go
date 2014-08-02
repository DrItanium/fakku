package fakku

import (
	"testing"
)

func TestForumCategoriesApiFunction_1(t *testing.T) {
	_, err := GetForumCategories()
	if err != nil {
		t.Error(err)
	}
}
