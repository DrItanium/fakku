package fakku

import "testing"

func TestGetFrontPagePosts_1(t *testing.T) {
	_, err := GetFrontPage()
	if err != nil {
		t.Error(err)
	}
}

func TestGetFrontPagePoll_1(t *testing.T) {
	_, err := GetFrontPagePoll()
	if err != nil {
		t.Error(err)
	}
}

func TestGetFrontPageFeaturedTopics_1(t *testing.T) {
	_, err := GetFrontPageFeaturedTopics()

	if err != nil {
		t.Error(err)
	}
}
