package router

import (
	"github.com/marugoshi/gobm/presentation/httputils"
	"net/http"
)

var RoutesData = []Route{
	Route{Pattern: `^/bookmarks$`, Method: http.MethodGet, Handler: func(params httputils.Params) error {
		return nil
	}},
}
