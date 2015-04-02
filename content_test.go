package fakku

import (
	"testing"
)

const (
	// Test the manga listed in the API docs
	ContentTestingCategory           = "manga"
	ContentTestingName               = "right-now-while-cleaning-the-pool"
	ContentTestingUrl                = "https://www.fakku.net/manga/right-now-while-cleaning-the-pool"
	ContentTestingTag                = "Vanilla"
	ContentTestingDisplayName        = "Right now, while cleaning the pool"
	ContentTestingDisplayNameRelated = "Before the Pool Opens"
	ContentTestingPoster             = "Jacob"
	ContentTestingRelatedTotal       = 11236
	ContentTestingRelatedPages       = 1124
)

func TestContentGetSimple_1(t *testing.T) {
	c, err := GetContent(ContentTestingCategory, ContentTestingName)
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
	_, err := ContentComments(ContentTestingCategory, ContentTestingName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestContentGetReadOnline_1(t *testing.T) {
	onlineContent, err := ReadOnline(ContentTestingCategory, ContentTestingName)
	if err != nil {
		t.Fatal(err)
	}

	if len(onlineContent.Pages) != 25 {
		t.Errorf("Expected 25 pages, got %d pages", len(onlineContent.Pages))
	}
}

func TestGetContentDownloads_1(t *testing.T) {
	_, err := ContentDownloads(ContentTestingCategory, ContentTestingName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetContentRelated_1(t *testing.T) {
	related, err := RelatedContent(ContentTestingCategory, ContentTestingName)
	if err != nil {
		t.Fatal(err)
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
	_, err := ReadOnline(CategoryManga, "nanako-san-english")
	if err == nil {
		t.Fatal("DMCA takedown notice not found!")
	}
	switch err.(type) {
	case *ErrorStatus: // FUUUUU it is a pointer!
		q := err.(*ErrorStatus)
		if q.KnownError {
			t.Log(err)
		} else {
			t.Error(err)
		}
	default:
		t.Error(err)
	}
}

func TestContentDoesntExist_1(t *testing.T) {
	_, err := ReadOnline(CategoryManga, "renai-sample-ch01-english")
	if err == nil {
		// try a second one since this is a little hard to test :/
		t.Fatal("Did not fail as expected!")
	}
	msg := err.Error()
	if msg == ErrorContentDoesntExist {
		t.Log(msg)
	} else {
		t.Error(msg)
	}
}
