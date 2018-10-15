package handler

import (
	"github.com/marugoshi/gobm/presentation/httputils"
)

type BookmarkHandler interface {
	Index(params httputils.Params) error
}

type bookmarkHandler struct {

}

func NewBookmarkHandler() BookmarkHandler {
	return &bookmarkHandler{}
}

func (b *bookmarkHandler) Index(params httputils.Params) error {
	return nil
}