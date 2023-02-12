package katana

type updatesPtr struct {
	Updates []Update
}

func NewPtr() updatesPtr {
	updates := make([]Update, 0)

	ptr := updatesPtr{
		Updates: updates,
	}

	return ptr
}

type Update struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type Manga struct {
	Chapters    []Chapter `json:"chapters"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Status      string    `json:"status"`
	Tags        []string  `json:"tags"`
	Title       string    `json:"title"`
}

func NewManga() *Manga {
	manga := new(Manga)
	manga.Chapters = make([]Chapter, 0)
	manga.Tags = make([]string, 0)

	return manga
}

type Chapter struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
