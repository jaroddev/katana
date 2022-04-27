package fake

import (
	"net"
	"net/http"
	"net/http/httptest"
)

const (
	HTMLFilePath    = "./katana.html"
	HTMLChapterPath = "./katanac.html"
	Endpoint        = "/"
	ChapterEndpoint = "/chapter"

	Test_DOMAIN = "127.0.0.1:8080"
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
	mux.HandleFunc(Endpoint, mangaMockedHandler)
	mux.HandleFunc(ChapterEndpoint, chapterMockedHandler)
	return mux
}

func mangaMockedHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, HTMLFilePath)
}

func chapterMockedHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, HTMLChapterPath)
}
