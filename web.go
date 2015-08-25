package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/carbocation/interpose"
	"github.com/carbocation/interpose/adaptors"
	"github.com/polds/imgbase64"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	// middleware
	middle := interpose.New()
	middle.Use(adaptors.FromNegroni(cors.New(cors.Options{
		// CORS
		AllowedOrigins: []string{"*"},
	})))

	// router
	router := pat.New()
	middle.UseHandler(router)

	router.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")

		if url != "" {
			encoded := imgbase64.FromRemote(url)
			fmt.Fprint(w, encoded)
		} else {
			fmt.Fprint(w, "/?url=<your-image-url> returns the base64 data-URI of that image.")
		}
	}))

	// listen
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Print("listening...")
	http.ListenAndServe(":"+port, middle)
}
