package manga

type Manga struct {
	Chapters    []Chapter `json:"chapters"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Status      string    `json:"status"`
	Tags        []string  `json:"tags"`
	Title       string    `json:"title"`
}

type Chapter struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
