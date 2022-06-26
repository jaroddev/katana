package fake

import (
	"net"
	"net/http"
	"net/http/httptest"
)

const Test_DOMAIN = "127.0.0.1:8080"

const (
	HTMLUpdatePath  = "./katana.html"
	HTMLMangaPath   = "./katanac.html"
	HTMLChapterPath = "./katanai.html"
)

const (
	UpdateEndpoint  = "/"
	MangaEndpoint   = "/manga"
	ChapterEndpoint = "/chapter"
)

func Server() *httptest.Server {
	l, _ := net.Listen("tcp", Test_DOMAIN)

	ts := httptest.NewUnstartedServer(handler())
	ts.Listener.Close()
	ts.Listener = l

	return ts
}

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(UpdateEndpoint, updateMockedHandler)
	mux.HandleFunc(MangaEndpoint, mangaMockedHandler)
	mux.HandleFunc(ChapterEndpoint, chapterMockedHandler)
	return mux
}

func updateMockedHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, HTMLUpdatePath)
}

func mangaMockedHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, HTMLMangaPath)
}

func chapterMockedHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, HTMLChapterPath)
}
