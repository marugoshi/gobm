package main

import (
	"github.com/marugoshi/gobm/presentation/httputils"
	"github.com/marugoshi/gobm/presentation/router"
	"github.com/marugoshi/gobm/registry"
	"log"
	"net/http"
)

func main() {
	registry := registry.NewRegistry()
	defer registry.DB.Close()
	router := router.NewRouter(registry, httputils.ContentTypeTextHtml)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Could not start: %s\n", err.Error())
	}
}
