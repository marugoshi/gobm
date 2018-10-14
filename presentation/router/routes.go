package router

import (
	"github.com/marugoshi/gobm/presentation/httputils"
	"net/http"
)

var Routes = []Route{
	Route{Pattern: `^/bookmarks$`, Method: http.MethodGet, HandleFunc: func(params httputils.HandleFuncParams) error {
		return nil
	}},
}
