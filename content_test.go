package fakku

import (
	"testing"
)

const (
	// Test the manga listed in the API docs
	ContentTestingCategory           = "manga"
	ContentTestingName               = "right-now-while-cleaning-the-pool"
	ContentTestingUrl                = "http://www.fakku.net/manga/right-now-while-cleaning-the-pool"
	ContentTestingTag                = "Vanilla"
	ContentTestingDisplayName        = "Right now, while cleaning the pool"
	ContentTestingDisplayNameRelated = "Before the Pool Opens"
	ContentTestingPoster             = "Jacob"
	ContentTestingRelatedTotal       = 11236
	ContentTestingRelatedPages       = 1124
)

func TestContentGetSimple_1(t *testing.T) {
	c, err := GetContentInformation(ContentTestingCategory, ContentTestingName)
	if err != nil {
		t.Error(err)
	}
	if c.Name != ContentTestingDisplayName {
		t.Errorf("Expected: %s, Got: %s", ContentTestingDisplayName, c.Name)
	}
	if c.Url != ContentTestingUrl {
		t.Errorf("Expected: %s, Got: %s", ContentTestingUrl, c.Url)
	}
	foundTag := false
	for i := 0; i < len(c.Tags); i++ {
		if c.Tags[i].Attribute == ContentTestingTag {
			foundTag = true
			break
		}
	}
	if !foundTag {
		t.Errorf("Didn't find tag %s but should have!", ContentTestingTag)
	}
}

func TestContentGetComments_1(t *testing.T) {
	/* tests to make sure that the functions work */
	_, err := GetContentComments(ContentTestingCategory, ContentTestingName)
	if err != nil {
		t.Error(err)
	}
}

func TestContentGetReadOnline_1(t *testing.T) {
	onlineContent, err := GetContentReadOnline(ContentTestingCategory, ContentTestingName)
	if err != nil {
		t.Error(err)
	}

	if len(onlineContent.Pages) != 25 {
		t.Errorf("Expected 25 pages, got %d pages", len(onlineContent.Pages))
	}
}

func TestGetContentDownloads_1(t *testing.T) {
	downloads, err := GetContentDownloads(ContentTestingCategory, ContentTestingName)
	if err != nil {
		t.Error(err)
	}
	if downloads.Total != 1 {
		t.Errorf("Expected 1 page, got %d pages", downloads.Total)
	} else {
		if len(downloads.Downloads) != 1 {
			t.Errorf("Expected 1 download, got %d downloads", len(downloads.Downloads))
		} else {
			if downloads.Downloads[0].Poster != ContentTestingPoster {
				t.Errorf("Expected download poster: %s, got %s", ContentTestingPoster, downloads.Downloads[0].Poster)
			}
		}
	}
}

func TestGetContentRelated_1(t *testing.T) {
	related, err := GetRelatedContentAll(ContentTestingCategory, ContentTestingName)
	if err != nil {
		t.Error(err)
	}
	if related.Related[0].Name != ContentTestingDisplayNameRelated {
		t.Errorf("First entry in related content was %s not %s as expected!", related.Related[0].Name, ContentTestingDisplayNameRelated)
	}
	// Since this can change, we need to do at least checks
	if related.Total < ContentTestingRelatedTotal {
		t.Errorf("Was unable to find an adequte enough of related content! Expected at least: %d, Got: %d", ContentTestingRelatedTotal, related.Total)
	}

	if related.Pages < ContentTestingRelatedPages {
		t.Errorf("Was unable to find an adequte enough of related content pages! Expected at least: %d, Got: %d", ContentTestingRelatedPages, related.Total)
	}
}

// Test to see if the DMCA takedowns are correctly caught
func TestDMCATakedown_1(t *testing.T) {
	_, err := GetContentReadOnline(CategoryManga, "renai-sample-ch01-english")
	if err == nil {
		t.Error("DMCA takedown notice not found!")
	} else {
		t.Log(err)
	}
}
