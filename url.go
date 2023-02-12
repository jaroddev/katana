package katana

import (
	"strconv"
)

const (
	BASE_DOMAIN = "mangakatana.com"
	latest      = "https://mangakatana.com/latest/page"
)

func Url(pageNumber uint64) string {
	if pageNumber == 0 {
		pageNumber = 1
	}

	url := latest + "/" + strconv.FormatUint(pageNumber, 10)

	return url
}
