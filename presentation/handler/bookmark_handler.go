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
	BookmarkNew(ctx context.Context, api httputils.Api) error
	BookmarkCreate(ctx context.Context, api httputils.Api) error
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

func (b *bookmarkHandler) BookmarkNew(ctx context.Context, api httputils.Api) error {
	return api.Html(200, nil, b.templates("new.html")...)
}

func (b *bookmarkHandler) BookmarkCreate(ctx context.Context, api httputils.Api) error {
	title := api.Request.FormValue("title")
	url := api.Request.FormValue("url")
	memo := api.Request.FormValue("memo")
	params := &data.Bookmark{0, url, title, memo}
	_, err := b.BookmarkService.Create(ctx, params)
	if err != nil {
		return err
	}
	// TODO: redirect
	return b.BookmarkIndex(ctx, api)
}

func (b *bookmarkHandler) BookmarkEdit(ctx context.Context, api httputils.Api) error {
	id, _ := strconv.ParseInt(api.Params[0], 10, 64)
	a, err := b.BookmarkService.Bookmark(ctx, id)
	if err != nil {
		return err
	}
	return api.Html(200, a, b.templates("edit.html")...)
}

func (b *bookmarkHandler) BookmarkUpdate(ctx context.Context, api httputils.Api) error {
	id, _ := strconv.ParseInt(api.Params[0], 10, 64)
	title := api.Request.FormValue("title")
	url := api.Request.FormValue("url")
	memo := api.Request.FormValue("memo")
	params := &data.Bookmark{id, url, title, memo}
	a, err := b.BookmarkService.Update(ctx, params)
	if err != nil {
		return err
	}
	// TODO: redirect
	return api.Html(200, a, b.templates("edit.html")...)
}

func (b *bookmarkHandler) templates(main string) []string {
	results := []string{}
	results = append(results, b.templateDir+main)
	results = append(results, b.partialDir+"header.html")
	results = append(results, b.partialDir+"footer.html")
	return results
}
