package service

import (
	"context"
	"github.com/marugoshi/gobm/presentation/httputils"
)

type BookmarkService interface {
	Bookmarks(ctx context.Context, http httputils.Http)
	ShowBookmark(ctx context.Context, http httputils.Http)
}

type bookmarkService struct {
}

func NewBookmarkService() BookmarkService {
	return &bookmarkService{}
}

func(b *bookmarkService) Bookmarks(ctx context.Context, http httputils.Http) {

}

func(b *bookmarkService) ShowBookmark(ctx context.Context, http httputils.Http) {

}