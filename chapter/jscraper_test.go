package chapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrapRealChapterPageWithJS(t *testing.T) {
	url := "https://mangakatana.com/manga/rascal-does-not-dream-of-logical-witch.25909/c10"

	scraper := DScraper{
		Chapter: &Chapter{
			Images: make([]string, 0),
		},
	}
	scraper.GetChapters(url)

	assert.NotNil(t, scraper.Chapter)
	assert.NotEmpty(t, scraper.Chapter)
	assert.Greater(t, len(scraper.Chapter.Images), 0)

	assert.NotEqual(t, scraper.Chapter.Images[0], "about:blank")
	assert.NotEqual(t, scraper.Chapter.Images[len(scraper.Chapter.Images)-1], "about:blank")

	assert.NotEqual(t, scraper.Chapter.Images[0], "#")
	assert.NotEqual(t, scraper.Chapter.Images[len(scraper.Chapter.Images)-1], "#")
}
