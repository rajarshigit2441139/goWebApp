package main

import (
	"log"
	"net/http"

	"github.com/rajarshigit2441139/goWebApp/router"
)

func main() {
	r := router.SetupRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	port := ":8089" // Change the port to 8089
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
