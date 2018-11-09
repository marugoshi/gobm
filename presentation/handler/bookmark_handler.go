package handler

import (
	"context"
	"github.com/marugoshi/gobm/domain/data"
	"github.com/marugoshi/gobm/domain/service"
	"github.com/marugoshi/gobm/presentation/httputils"
	"os"
	"strconv"
)

type BookmarkHandler interface {
	BookmarkIndex(ctx context.Context, api httputils.Api) error
	BookmarkEdit(ctx context.Context, api httputils.Api) error
	BookmarkUpdate(ctx context.Context, api httputils.Api) error
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
	a, err := b.BookmarkService.Bookmarks(ctx)
	if err != nil {
		return err
	}
	return api.Html(200, a, b.templates("index.html")...)
}

func (b *bookmarkHandler) BookmarkEdit(ctx context.Context, api httputils.Api) error {
	id, _ := strconv.Atoi(api.Params[0])
	a, err := b.BookmarkService.Bookmark(ctx, id)
	if err != nil {
		return err
	}
	return api.Html(200, a, b.templates("edit.html")...)
}

func (b *bookmarkHandler) BookmarkUpdate(ctx context.Context, api httputils.Api) error {
	id, _ := strconv.Atoi(api.Params[0])
	title := api.Request.FormValue("title")
	url := api.Request.FormValue("url")
	memo := api.Request.FormValue("memo")
	params := &data.Bookmark{id, url, title, memo}
	a, err := b.BookmarkService.Update(ctx, params)
	if err != nil {
		return err
	}
	return api.Html(200, a, b.templates("edit.html")...)
}

func (b *bookmarkHandler) templates(main string) []string {
	results := []string{}
	results = append(results, b.templateDir+main)
	results = append(results, b.partialDir+"header.html")
	results = append(results, b.partialDir+"footer.html")
	return results
}
