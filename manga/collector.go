package manga

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/jaroddev/katana/fake"
)

const (
	BASE_DOMAIN = "mangakatana.com"
)

func customCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains(
			BASE_DOMAIN,
			fake.Test_DOMAIN,
		),
		colly.AllowURLRevisit(),
	)

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting manga page: ", request.URL.String())
	})

	return collector
}
