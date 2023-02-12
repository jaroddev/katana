package katana_test

import (
	"testing"

	"github.com/jaroddev/katana"
	"github.com/stretchr/testify/assert"
)

func TestGetUrl(t *testing.T) {

	table := []struct {
		name  string
		value uint64
		wants string
	}{
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
			assert.Equal(t, katana.Url(row.value), row.wants)
		})

	}

}
