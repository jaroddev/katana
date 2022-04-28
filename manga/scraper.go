package manga

import (
	"strings"

	"github.com/gocolly/colly"
)

type Scraper struct {
	Collector *colly.Collector
	Manga     *Manga
}

func GetScraper() (scraper Scraper) {
	collector := customCollector()

	manga := new(Manga)
	manga.Chapters = make([]Chapter, 0)
	manga.Tags = make([]string, 0)

	collector.OnHTML(".info", scrapInfo(manga))

	collector.OnHTML(".cover img", scrapCover(manga))

	collector.OnHTML(".info .genres a", scrapTags(manga))

	collector.OnHTML(".summary", scrapDescription(manga))

	collector.OnHTML(".chapters tr", scrapChapter(manga))

	scraper.Collector = collector
	scraper.Manga = manga

	return
}

func scrapCover(manga *Manga) func(element *colly.HTMLElement) {
	return func(element *colly.HTMLElement) {
		manga.Image = element.Attr("src")
	}
}

func scrapDescription(manga *Manga) func(element *colly.HTMLElement) {
	return func(element *colly.HTMLElement) {
		manga.Description = strings.TrimSpace(element.ChildText("p"))
	}
}

func scrapInfo(manga *Manga) func(element *colly.HTMLElement) {
	return func(element *colly.HTMLElement) {
		manga.Status = strings.TrimSpace(element.ChildText(".status"))
		manga.Title = strings.TrimSpace(element.ChildText(".heading"))
	}
}

func scrapChapter(manga *Manga) func(element *colly.HTMLElement) {

	chapters := manga.Chapters

	return func(element *colly.HTMLElement) {

		chapter := Chapter{
			Name: strings.TrimSpace(element.ChildText("a")),
			Url:  element.ChildAttr("a", "href"),
		}

		chapters = append(chapters, chapter)
		manga.Chapters = chapters

	}

}

func scrapTags(manga *Manga) func(element *colly.HTMLElement) {
	tags := manga.Tags

	return func(element *colly.HTMLElement) {
		tag := strings.TrimSpace(element.Text)

		tags = append(tags, tag)
		manga.Tags = tags
	}

}
