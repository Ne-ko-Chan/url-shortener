package main

import (
	"link-shortener/internal/api"
	"log"
	"net/http"
)

func main() {
	// Init database, shortener service and qr code generator service
	h, err := api.NewHandler(nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	log.Print("Server started")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
