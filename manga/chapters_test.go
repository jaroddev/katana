package manga

import (
	"fmt"
	"testing"

	"github.com/jaroddev/katana/fake"
	"github.com/stretchr/testify/assert"
)

func TestScrapChapter(t *testing.T) {
	// create a listener with the desired port.
	ts := fake.Server()

	ts.Start()
	defer ts.Close()

	url := fmt.Sprintf("%s%s", ts.URL, fake.ChapterEndpoint)

	scraper := NewChaptersScrapper()
	scraper.Collector.Visit(url)

	assert.NotNil(t, scraper.Chapters)
	assert.NotEmpty(t, scraper.Chapters)
	assert.Len(t, *scraper.Chapters, 137)

	fmt.Println(scraper.Chapters)
}
