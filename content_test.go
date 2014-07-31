package fakku

import (
	"testing"
)

func TestContentGetSimple_1(t *testing.T) {
	// Test the manga listed in the API docs
	compareName := "Right now, while cleaning the pool"
	compareUrl := "http://www.fakku.net/manga/right-now-while-cleaning-the-pool"
	compareTag := "Vanilla"
	c, err := GetContentInformation("manga", "right-now-while-cleaning-the-pool")
	if err != nil {
		t.Error(err)
	}
	if c.Name != compareName {
		t.Errorf("Expected: %s, Got: %s", compareName, c.Name)
	}
	if c.Url != compareUrl {
		t.Errorf("Expected: %s, Got: %s", compareUrl, c.Url)
	}
	foundTag := false
	for i := 0; i < len(c.Tags); i++ {
		if c.Tags[i].Attribute == compareTag {
			foundTag = true
			break
		}
	}
	if !foundTag {
		t.Errorf("Didn't find tag %s but should have!", compareTag)
	}
}

func TestContentGetComments_1(t *testing.T) {
	/* tests to make sure that the functions work */
	_, err := GetContentComments("manga", "right-now-while-cleaning-the-pool")
	if err != nil {
		t.Error(err)
	}
}

func TestContentGetReadOnline_1(t *testing.T) {
	onlineContent, err := GetContentReadOnline("manga", "right-now-while-cleaning-the-pool")
	if err != nil {
		t.Error(err)
	}

	if len(onlineContent.Pages) != 25 {
		t.Errorf("Expected 25 pages, got %d pages", len(onlineContent.Pages))
	}
}

func TestGetContentDownloads_1(t *testing.T) {
	downloadPoster := "Jacob"
	downloads, err := GetContentDownloads("manga", "right-now-while-cleaning-the-pool")
	if err != nil {
		t.Error(err)
	}
	if downloads.Total != 1 {
		t.Errorf("Expected 1 page, got %d pages", downloads.Total)
	} else {
		if len(downloads.Downloads) != 1 {
			t.Errorf("Expected 1 download, got %d downloads", len(downloads.Downloads))
		} else {
			if downloads.Downloads[0].Poster != downloadPoster {
				t.Errorf("Expected download poster: %s, got %s", downloadPoster, downloads.Downloads[0].Poster)
			}
		}
	}
}
