package main

import (
	"log"
	"net/http"

	"github.com/rajarshigit2441139/goWebApp/router"
)

func main() {
	r := router.SetupRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
