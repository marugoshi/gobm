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
	pathPrefix string
}

func NewBookmarkService() BookmarkService {
	current, _ := os.Getwd()
	return &bookmarkService{current + "/presentation/view/bookmark"}
}

func(b *bookmarkService) Bookmarks(ctx context.Context, http httputils.Http) error {
	data := struct{
		Key string
	}{
		Key: "hoge",
	}
	return http.Html(200, "index", b.pathPrefix + "/index.html", data)
}

func(b *bookmarkService) Bookmark(ctx context.Context, http httputils.Http) error {
	return http.RawText(200, http.Params[0])
}