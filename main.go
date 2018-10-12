package main

import (
	"github.com/marugoshi/gobm/presentation/router"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter(router.ContentTypeTextHtml)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Could not start: %s\n", err.Error())
	}
}
