package router

import (
	"github.com/marugoshi/gobm/presentation/httputils"
	"net/http"
	"regexp"
)

var RoutesData = []Route{
	Route{re(`^/bookmarks$`), http.MethodGet, func(params httputils.Params) error {
		return nil
	}},
}

func re(path string) *regexp.Regexp {
	return regexp.MustCompile(path)
}