package chapter

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

type DScraper struct {
	Chapter *Chapter
}

func (scraper *DScraper) GetChapters(url string) {

	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(url, g.Opt.ParseFunc)
		},
		ParseFunc: scrapChapter(scraper.Chapter),
	}).Start()
}

func scrapChapter(chapter *Chapter) func(*geziyor.Geziyor, *client.Response) {

	return func(g *geziyor.Geziyor, r *client.Response) {

		r.HTMLDoc.Find("#imgs .wrap_img").Find("img").Each(
			func(_ int, s *goquery.Selection) {

				image, exist := s.Attr("data-src")

				if exist {
					chapter.Images = append(chapter.Images, image)
				}

			},
		)
	}
}
