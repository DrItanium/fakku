package fakku

import (
	"testing"
)

func TestForumCategoriesApiFunction_1(t *testing.T) {
	output, err := GetForumCategories()
	if err != nil {
		t.Error(err)
	}
	result := output.Categories[0]

	if result.Title != "FAKKU" {
		t.Error("Didn't get FAKKU forum!")
	}
	if result.Forums[0].Name != "Front Page News" {
		t.Error("Front Page News wasn't found!")
	}
	t.Logf("Forum: %s", result.Forums[0].Name)
	rt := result.Forums[0].RecentTopic
	t.Logf("Recent topic title: %s", rt.Title)
}

func TestForumCategoriesApiFunction_2(t *testing.T) {
	output, err := GetForumTopics("random")
	if err != nil {
		t.Error(err)
	}

	if output.Forum.Name != "Random" {
		t.Error("Didn't get the Random forum!")
	}
	if output.Page != 1 {
		t.Error("Didn't get page 11")
	}
}
