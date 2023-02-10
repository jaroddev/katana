package update

import (
	"github.com/gocolly/colly/v2"
)

const (
	BASE_URL = "https://mangakatana.com/latest/page/"
)

type updatesPtr struct {
	Updates []Update
}

type Scraper struct {
	Collector *colly.Collector
	Ptr       *updatesPtr
}

func GetScraper() (scrapper Scraper) {
	collector := customCollector()
	updates := make([]Update, 0)

	ptr := updatesPtr{
		Updates: updates,
	}

	collector.OnHTML("#book_list .item", scrapUpdate(&ptr))

	scrapper.Collector = collector
	scrapper.Ptr = &ptr

	return
}

func scrapUpdate(updatesPtr *updatesPtr) func(element *colly.HTMLElement) {

	return func(element *colly.HTMLElement) {

		update := Update{
			Title: element.ChildText(".title a"),
			Url:   element.ChildAttr(".title a", "href"),
		}

		updatesPtr.Updates = append(updatesPtr.Updates, update)
	}

}
