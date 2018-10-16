package handler

import (
	"github.com/marugoshi/gobm/presentation/httputils"
)

type BookmarkHandler interface {
	Bookmarks(params httputils.Http) error
}

type bookmarkHandler struct {

}

func NewBookmarkHandler() BookmarkHandler {
	return &bookmarkHandler{}
}

func (b *bookmarkHandler) Bookmarks(http httputils.Http) error {
	http.Text(200, "hoge")
	return nil
}