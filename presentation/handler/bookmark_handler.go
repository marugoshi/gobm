package handler

import (
	"github.com/marugoshi/gobm/presentation/httputils"
)

type BookmarkHandler interface {
	Bookmarks(params httputils.Params) error
}

type bookmarkHandler struct {

}

func NewBookmarkHandler() BookmarkHandler {
	return &bookmarkHandler{}
}

func (b *bookmarkHandler) Bookmarks(params httputils.Params) error {
	params.Text(200, "hoge")
	return nil
}