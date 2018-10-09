package Router

var Routes = []Route{
	Route{Pattern: "^/bookmarks$", Method: GET, Handler: func(e *Exchange){
		e.Text(200, "hello")
	}},
}