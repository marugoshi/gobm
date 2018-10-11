package Router

import (
	"net/http"
)

var Routes = []Route{
	Route{Pattern: "^/bookmarks$", Method: http.MethodGet, Handler: func(e *Exchange) {
		e.Text(200, "hello")
	}},
}
