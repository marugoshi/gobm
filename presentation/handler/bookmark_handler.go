package handler

import (
	"github.com/marugoshi/gobm/presentation/httputils"
)

type BookmarkHandler interface {
	Bookmarks(params httputils.Api) error
}

type bookmarkHandler struct {

}

func NewBookmarkHandler() BookmarkHandler {
	return &bookmarkHandler{}
}

func (b *bookmarkHandler) Bookmarks(params httputils.Api) error {
	params.Text(200, "hoge")
	return nil
}