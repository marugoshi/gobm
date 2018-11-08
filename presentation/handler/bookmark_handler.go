package handler

import (
	"context"
	"github.com/marugoshi/gobm/domain/service"
	"github.com/marugoshi/gobm/presentation/httputils"
	"os"
	"strconv"
)

type BookmarkHandler interface {
	BookmarkIndex(ctx context.Context, api httputils.Api) error
	BookmarkShow(ctx context.Context, api httputils.Api) error
}

type bookmarkHandler struct {
	service.BookmarkService
	partialDir  string
	templateDir string
}

func NewBookmarkHandler(s service.BookmarkService) BookmarkHandler {
	current, _ := os.Getwd()
	partialDir := current + "/presentation/view/partial/"
	templateDir := current + "/presentation/view/bookmark/"
	return &bookmarkHandler{s, partialDir, templateDir}
}

func (b *bookmarkHandler) BookmarkIndex(ctx context.Context, api httputils.Api) error {
	data, err := b.BookmarkService.Bookmarks(ctx)
	if err != nil {
		return err
	}
	return api.Html(200, data, b.templates("index.html")...)
}

func (b *bookmarkHandler) BookmarkShow(ctx context.Context, api httputils.Api) error {
	id, _ := strconv.Atoi(api.Params[0])
	data, err := b.BookmarkService.Bookmark(ctx, id)
	if err != nil {
		return err
	}
	return api.Html(200, data, b.templateDir+"show.html")
}

func (b *bookmarkHandler) templates(main string) []string {
	results := []string{}
	results = append(results, b.partialDir+"header.html")
	results = append(results, b.partialDir+"footer.html")
	results = append(results, b.templateDir+main)
	return results
}