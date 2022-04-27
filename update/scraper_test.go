package update

import (
	"fmt"
	"testing"

	"github.com/jaroddev/katana/fake"
	"github.com/stretchr/testify/assert"
)

func TestGetUrl(t *testing.T) {
	scraper := New()

	table := []struct {
		name  string
		value int
		wants string
	}{
		{
			name:  "pageNumber lower than 0",
			value: -2,
			wants: "https://mangakatana.com/latest/page/1",
		},
		{
			name:  "pageNumber equals 0",
			value: 0,
			wants: "https://mangakatana.com/latest/page/1",
		},
		{
			name:  "pageNumber equals 1",
			value: 1,
			wants: "https://mangakatana.com/latest/page/1",
		},
		{
			name:  "pageNumber equals 3",
			value: 3,
			wants: "https://mangakatana.com/latest/page/3",
		},
	}

	for _, row := range table {

		t.Run(row.name, func(t *testing.T) {
			assert.Equal(t, scraper.Url(row.value), row.wants)
		})

	}

}

func TestScrapManga(t *testing.T) {
	// create a listener with the desired port.
	ts := fake.Server()

	ts.Start()
	defer ts.Close()

	url := fmt.Sprintf("%s%s", ts.URL, fake.Endpoint)

	scraper := New()
	scraper.Collector.Visit(url)

	// 	printMangaList(*scraper.Mangas)

	assert.NotNil(t, scraper.Mangas)
	assert.NotEmpty(t, scraper.Mangas)
	assert.Len(t, *scraper.Mangas, 20)
}

func printMangaList(mangas []Update) {
	fmt.Println("[")
	for _, manga := range mangas {
		fmt.Println("\t{")
		fmt.Println("\t\t", "\"", "Url", "\":", "\"", manga.Url, "\",")
		fmt.Println("\t\t", "\"", "Title", "\":", "\"", manga.Title, "\"")
		fmt.Println("\t},")
	}
	fmt.Println("]")
}
