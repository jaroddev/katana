package katana

import (
	"os"
	"path/filepath"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

func TestScrapManga(t *testing.T) {
	path := filepath.Join("html", "katanac.html")
	view, _ := os.ReadFile(path)

	doc, err := fromByte(view)
	manga := getManga(doc)

	assert.NoError(t, err)

	assert.NotEmpty(t, manga.Title)
	assert.NotEmpty(t, manga.Status)
	assert.NotEmpty(t, manga.Description)
	assert.NotEmpty(t, manga.Image)

	assert.NotNil(t, manga.Tags)
	assert.NotEmpty(t, manga.Tags)
	assert.Len(t, manga.Tags, 7)

	assert.Equal(t, manga.Tags[0], "Action")
	assert.Equal(t, manga.Tags[len(manga.Tags)-1], "Gender Bender")

	assert.NotNil(t, manga.Chapters)
	assert.NotEmpty(t, manga.Chapters)
	assert.Len(t, manga.Chapters, 137)
	assert.Equal(t, manga.Chapters[0].Name, "Chapter 136")
}

func TestScrapUpdate(t *testing.T) {
	path := filepath.Join("html", "katana.html")
	view, _ := os.ReadFile(path)

	doc, err := fromByte(view)
	ptr := getUpdates(doc)

	assert.NoError(t, err)

	assert.NotNil(t, ptr.Updates)
	assert.NotEmpty(t, ptr.Updates)
	assert.Len(t, ptr.Updates, 20)
}
