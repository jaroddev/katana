package update

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jaroddev/katana/fake"
	"github.com/stretchr/testify/assert"
)

func TestGetUrl(t *testing.T) {
	scraper := GetScraper()

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

var ts *httptest.Server

func TestMain(m *testing.M) {
	ts = fake.Server()

	ts.Start()
	defer ts.Close()
	os.Exit(m.Run())
}

func TestScrapUpdate(t *testing.T) {
	url := fmt.Sprintf("%s%s", ts.URL, fake.Endpoint)

	scraper := GetScraper()
	scraper.Collector.Visit(url)

	assert.NotNil(t, scraper.Ptr.Updates)
	assert.NotEmpty(t, scraper.Ptr.Updates)
	assert.Len(t, scraper.Ptr.Updates, 20)
}
