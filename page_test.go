package fakku

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	ContentName     = "right-now-while-cleaning-the-pool"
	ContentCategory = CategoryDoujinshi
)

func loadPages() (PageList, error) {
	content, err := ReadOnline(ContentCategory, ContentName)
	if err != nil {
		return nil, err
	} else {
		return content.Pages, nil
	}
}
func TestPage_ImageBytes_1(t *testing.T) {
	pages, err := loadPages()
	if err != nil {
		t.Fatal(err)
	}
	if len(pages) < 1 {
		t.Fatal("No pages found!")
	}
	_, err2 := pages[0].ImageBytes()
	if err2 != nil {
		t.Error(err2)
	}
}

func TestPage_SaveImage_1(t *testing.T) {
	pages, err := loadPages()
	if err != nil {
		t.Fatal(err)
	}
	if len(pages) < 1 {
		t.Fatal("No pages found!")
	}
	file, err2 := ioutil.TempFile("", "image")
	if err2 != nil {
		t.Fatal(err2)
	}
	filename := file.Name()
	if err3 := pages[0].SaveImage(filename, 0644); err3 != nil {
		t.Fatal(err2)
	}
	file.Close()
	os.Remove(filename)
}

func TestPage_GetImage_1(t *testing.T) {
	pages, err := loadPages()
	if err != nil {
		t.Fatal(err)
	}
	if len(pages) < 1 {
		t.Fatal("No pages found!")
	}
	img, err2 := pages[0].Image()
	if err2 != nil {
		t.Fatal(err2)
	}
	t.Logf("Image Bounds: %s", img.Bounds().String())
}

func TestPage_ThumbBytes_1(t *testing.T) {
	pages, err := loadPages()
	if err != nil {
		t.Fatal(err)
	}
	if len(pages) < 1 {
		t.Fatal("No pages found!")
	}
	_, err2 := pages[0].ThumbBytes()
	if err2 != nil {
		t.Error(err2)
	}
}
func TestPage_SaveThumb_1(t *testing.T) {
	pages, err := loadPages()
	if err != nil {
		t.Fatal(err)
	}
	if len(pages) < 1 {
		t.Fatal("No pages found!")
	}
	file, err2 := ioutil.TempFile("", "thumbnail")
	if err2 != nil {
		t.Fatal(err2)
	}
	filename := file.Name()
	if err3 := pages[0].SaveThumb(filename, 0644); err3 != nil {
		t.Fatal(err2)
	}
	file.Close()
	os.Remove(filename)
}

func TestPage_GetThumbnail_1(t *testing.T) {
	pages, err := loadPages()
	if err != nil {
		t.Fatal(err)
	}
	if len(pages) < 1 {
		t.Fatal("No pages found!")
	}
	img, err2 := pages[0].Thumbnail()
	if err2 != nil {
		t.Fatal(err2)
	}
	t.Logf("Thumbnail Bounds: %s", img.Bounds().String())

}
