package api

import "net/http"

type UrlShortenerService interface {
	Shorten(url string) (string, error)
}

type QrGeneratorService interface {
	Generate(url string) ([]byte, error)
}

type DatabaseService interface {
	Store(url string) (key string, err error)
	Load(key string) (url string, err error)
}

type Handler struct {
	db          DatabaseService
	shortener   UrlShortenerService
	qtGenerator QrGeneratorService
}

func NewHandler(
	db DatabaseService,
	shortener UrlShortenerService,
	qtGenerator QrGeneratorService,
) (*Handler, error) {
	return &Handler{
		db:          db,
		shortener:   shortener,
		qtGenerator: qtGenerator,
	}, nil
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dummy root handler:)\nHello!"))
	})
}
