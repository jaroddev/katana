package manga

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jaroddev/katana/fake"
)

const (
	BASE_DOMAIN = "mangakatana.com"
)

type ChaptersScrapper struct {
	Collector colly.Collector
	Chapters  *[]Chapter
}

func NewChaptersScrapper() (scrapper ChaptersScrapper) {
	collector := colly.NewCollector(
		colly.AllowedDomains(
			BASE_DOMAIN,
			fake.Test_DOMAIN,
		),
		colly.AllowURLRevisit(),
	)
	chapters := make([]Chapter, 0)

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting manga page: ", request.URL.String())
	})

	collector.OnHTML(".chapters tr", func(element *colly.HTMLElement) {
		chapter := Chapter{
			Name: strings.TrimSpace(element.ChildText("a")),
			Url:  element.ChildAttr("a", "href"),
		}

		chapters = append(chapters, chapter)
	})

	scrapper.Collector = *collector
	scrapper.Chapters = &chapters

	return
}
