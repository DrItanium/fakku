package fakku

import (
	"testing"
)

func TestGetCategoryIndex_1(t *testing.T) {
	_, err := GetCategoryIndex(CategoryManga)
	if err != nil {
		t.Error(err)
	}
}
