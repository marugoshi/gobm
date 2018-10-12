package router

import (
	"net/http"
)

var Routes = []Route{
	Route{Pattern: `^/bookmarks$`, Method: http.MethodGet, HandleFunc: func(e *Exchange) {
		e.Text(200, "hello")
	}},
}
