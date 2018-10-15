package handler

type Handlers interface {
	BookmarkHandler
}

type handlers struct {
	BookmarkHandler
}

func NewHandlers() Handlers {
	return &handlers{ NewBookmarkHandler() }
}