package model

import (
	"context"
	"github.com/marugoshi/gobm/domain/data"
)

type BookmarkModel interface {
	All(ctx context.Context, page int, perPage int) (interface{}, error)
	FindById(ctx context.Context, id int) (interface{}, error)
	Update(ctx context.Context, params *data.Bookmark) (interface{}, error)
}
