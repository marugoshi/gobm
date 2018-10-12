package main

import (
	"github.com/marugoshi/gobm/presentation/httputils"
	"github.com/marugoshi/gobm/presentation/router"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter(httputils.ContentTypeTextHtml)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Could not start: %s\n", err.Error())
	}
}
