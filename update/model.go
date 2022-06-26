package update

import "strconv"

const (
	URL = "https://mangakatana.com/latest/page"
)

type Update struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

func Url(pageNumber uint64) string {
	if pageNumber == 0 {
		pageNumber = 1
	}

	url := BASE_URL + strconv.FormatUint(pageNumber, 10)

	return url
}
