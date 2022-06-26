package update

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jaroddev/katana/fake"
	"github.com/stretchr/testify/assert"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	ts = fake.Server()

	ts.Start()
	defer ts.Close()
	os.Exit(m.Run())
}

func TestScrapUpdate(t *testing.T) {
	url := fmt.Sprintf("%s%s", ts.URL, fake.UpdateEndpoint)

	scraper := GetScraper()
	scraper.Collector.Visit(url)

	assert.NotNil(t, scraper.Ptr.Updates)
	assert.NotEmpty(t, scraper.Ptr.Updates)
	assert.Len(t, scraper.Ptr.Updates, 20)
}
