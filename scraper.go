package katana

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getManga(doc *goquery.Document) *Manga {
	manga := NewManga()

	scrapHTML(doc, ".info", scrapInfo(manga))
	scrapHTML(doc, ".cover img", scrapCover(manga))
	scrapHTML(doc, ".info .genres a", scrapTags(manga))
	scrapHTML(doc, ".summary", scrapDescription(manga))
	scrapHTML(doc, ".chapters tr", scrapChapter(manga))

	return manga
}

func GetManga(URL string) *Manga {
	doc, err := fromURL(URL)
	if err != nil {
		return nil
	}
	return getManga(doc)
}

func getUpdates(doc *goquery.Document) *updatesPtr {
	ptr := NewPtr()

	scrapHTML(doc, "#book_list .item", scrapUpdate(&ptr))

	return &ptr
}

func GetUpdates(URL string) *updatesPtr {
	doc, err := fromURL(URL)
	if err != nil {
		return nil
	}
	return getUpdates(doc)
}

func scrapCover(manga *Manga) func(element *HTMLElement) {
	return func(element *HTMLElement) {
		manga.Image = element.Attr("src")
	}
}

func scrapDescription(manga *Manga) func(element *HTMLElement) {
	return func(element *HTMLElement) {
		manga.Description = strings.TrimSpace(element.ChildText("p"))
	}
}

func scrapInfo(manga *Manga) func(element *HTMLElement) {
	return func(element *HTMLElement) {
		manga.Status = strings.TrimSpace(element.ChildText(".status"))
		manga.Title = strings.TrimSpace(element.ChildText(".heading"))
	}
}

func scrapChapter(manga *Manga) func(element *HTMLElement) {

	chapters := manga.Chapters

	return func(element *HTMLElement) {

		chapter := Chapter{
			Name: strings.TrimSpace(element.ChildText("a")),
			Url:  element.ChildAttr("a", "href"),
		}

		chapters = append(chapters, chapter)
		manga.Chapters = chapters

	}

}

func scrapTags(manga *Manga) func(element *HTMLElement) {
	tags := manga.Tags

	return func(element *HTMLElement) {

		tag := strings.TrimSpace(element.Text)

		tags = append(tags, tag)
		manga.Tags = tags
	}

}

func scrapUpdate(updatesPtr *updatesPtr) func(element *HTMLElement) {

	return func(element *HTMLElement) {

		update := Update{
			Title: element.ChildText(".title a"),
			Url:   element.ChildAttr(".title a", "href"),
		}

		updatesPtr.Updates = append(updatesPtr.Updates, update)
	}

}

func fromByte(data []byte) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(bytes.NewBuffer(data))
}

func fromURL(URL string) (*goquery.Document, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var bodyReader io.Reader = res.Body

	contentEncoding := strings.ToLower(res.Header.Get("Content-Encoding"))
	if !res.Uncompressed && (strings.Contains(contentEncoding, "gzip") || (contentEncoding == "" && strings.Contains(strings.ToLower(res.Header.Get("Content-Type")), "gzip"))) {
		bodyReader, err = gzip.NewReader(bodyReader)
		if err != nil {
			return nil, err
		}
		defer bodyReader.(*gzip.Reader).Close()
	}

	data, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	return fromByte(data)
}

func scrapHTML(doc *goquery.Document, selector string, f func(*HTMLElement)) {
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			element := NewHTMLElementFromSelectionNode(s, n)
			f(element)
		}
	})
}
