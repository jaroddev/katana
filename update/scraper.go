package update

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/jaroddev/katana/fake"
)

const (
	BASE_URL    = "https://mangakatana.com/latest/page/"
	BASE_DOMAIN = "mangakatana.com"
)

type Scrapper struct {
	Collector colly.Collector
	Mangas    *[]Update
}

func New() (scrapper Scrapper) {
	collector := colly.NewCollector(
		colly.AllowedDomains(
			BASE_DOMAIN,
			fake.Test_DOMAIN,
		),
		colly.AllowURLRevisit(),
	)
	mangas := make([]Update, 0)

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting update page: ", request.URL.String())
	})

	collector.OnHTML("#book_list .item", func(element *colly.HTMLElement) {
		manga := Update{
			Title: element.ChildText(".title a"),
			Url:   element.ChildAttr(".title a", "href"),
		}

		mangas = append(mangas, manga)
	})

	scrapper.Collector = *collector
	scrapper.Mangas = &mangas

	return
}

func (scrapper *Scrapper) Url(pageNumber int) string {
	if pageNumber < 1 {
		pageNumber = 1
	}

	url := BASE_URL + strconv.Itoa(pageNumber)

	return url
}
