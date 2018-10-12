package Router

import (
	"net/http"
)

var Routes = []Route{
	Route{Pattern: `^/bookmarkschange$`, Method: http.MethodGet, AppHandler: func(e *Exchange) {
		e.Text(200, "hello")
	}},
}
