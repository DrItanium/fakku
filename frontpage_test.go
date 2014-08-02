package fakku

import "testing"

func TestGetFrontPagePosts_1(t *testing.T) {
	_, err := GetFrontPage()
	if err != nil {
		t.Error(err)
	}
}
