package manga

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jaroddev/katana/fake"
	"github.com/stretchr/testify/assert"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	ts = fake.Server()

	ts.Start()
	defer ts.Close()
	os.Exit(m.Run())
}

func TestScrapInfo(t *testing.T) {
	url := fmt.Sprintf("%s%s", ts.URL, fake.MangaEndpoint)

	scraper := GetScraper()

	scraper.Collector.Visit(url)

	manga := scraper.Manga

	assert.NotEmpty(t, manga.Title)
	assert.NotEmpty(t, manga.Status)
}

func TestScrapDescription(t *testing.T) {
	url := fmt.Sprintf("%s%s", ts.URL, fake.MangaEndpoint)

	scraper := GetScraper()

	scraper.Collector.Visit(url)

	manga := scraper.Manga

	assert.NotEmpty(t, manga.Description)
}

func TestScrapCover(t *testing.T) {
	url := fmt.Sprintf("%s%s", ts.URL, fake.MangaEndpoint)

	scraper := GetScraper()

	scraper.Collector.Visit(url)

	manga := scraper.Manga
	assert.NotEmpty(t, manga.Image)
}

func TestScrapTags(t *testing.T) {
	url := fmt.Sprintf("%s%s", ts.URL, fake.MangaEndpoint)

	scraper := GetScraper()

	scraper.Collector.Visit(url)

	manga := scraper.Manga

	assert.NotNil(t, manga.Tags)
	assert.NotEmpty(t, manga.Tags)
	assert.Len(t, manga.Tags, 7)

	assert.Equal(t, manga.Tags[0], "Action")
	assert.Equal(t, manga.Tags[len(manga.Tags)-1], "Gender Bender")
}

func TestScrapChapter(t *testing.T) {
	url := fmt.Sprintf("%s%s", ts.URL, fake.MangaEndpoint)

	scraper := GetScraper()
	scraper.Collector.Visit(url)

	manga := scraper.Manga

	assert.NotNil(t, manga.Chapters)
	assert.NotEmpty(t, manga.Chapters)
	assert.Len(t, manga.Chapters, 137)
	assert.Equal(t, manga.Chapters[0].Name, "Chapter 136")

}
