package router

import (
	"net/http"
)

var Routes = []Route{
	Route{Pattern: `^/bookmarks$`, Method: http.MethodGet, HandleFunc: func(res http.ResponseWriter, req *http.Request, params []string) error {
		return nil
	}},
}
