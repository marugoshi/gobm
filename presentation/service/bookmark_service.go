package service

import (
	"context"
	"github.com/marugoshi/gobm/presentation/httputils"
	"os"
)

type BookmarkService interface {
	Bookmarks(ctx context.Context, http httputils.Http) error
	Bookmark(ctx context.Context, http httputils.Http) error
}

type bookmarkService struct {
}

func NewBookmarkService() BookmarkService {
	return &bookmarkService{}
}

func(b *bookmarkService) Bookmarks(ctx context.Context, http httputils.Http) error {
	data := struct{
		Key string
	}{
		Key: "hoge",
	}
	current, _ := os.Getwd()
	return http.Html(200, "index", current + "/presentation/view/bookmark/index.html", data)
}

func(b *bookmarkService) Bookmark(ctx context.Context, http httputils.Http) error {
	return http.RawText(200, http.Params[0])
}