package manga

type Manga struct {
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Status      string    `json:"status"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Chapters    []Chapter `json:"chapters"`
}

type Chapter struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
