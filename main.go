package main

import (
	"github.com/marugoshi/gobm/Presentation/Router"
	"log"
	"net/http"
)

func main() {
	router := Router.New("text/html")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Could not start: %s\n", err.Error())
	}
}
