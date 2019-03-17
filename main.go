package main

import (
	"github.com/marugoshi/gobm/presentation/httputils"
	. "github.com/marugoshi/gobm/presentation/router"
	. "github.com/marugoshi/gobm/registry"
	"github.com/marugoshi/gobm/shared/app_log"
	"net/http"
)

func main() {
	registry, err := NewRegistry()
	if err != nil {
		app_log.Fatalf("panic: %#v", err)
		panic("could not start.")
	}
	defer registry.DB.Close()
	router := NewRouter(registry, httputils.ContentTypeTextHtml)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		app_log.Fatalf("panic: %#v", err)
		panic("could not start.")
	}
}
