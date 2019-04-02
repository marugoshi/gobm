package registry

import (
	"database/sql"
	"github.com/marugoshi/gobm/domain/service"
	"github.com/marugoshi/gobm/infrastructure/storage/mysql"
	"github.com/marugoshi/gobm/infrastructure/storage/mysql/model"
	"github.com/marugoshi/gobm/presentation/handler"
	"github.com/pkg/errors"
)

type Registry struct {
	*sql.DB
	handler.BookmarkHandler
}

func NewRegistry() (Registry, error) {
	db, err := mysql.NewInstance()
	if err != nil {
		return Registry{}, errors.Wrap(err, "can not instantiate db.")
	}
	if err := db.Ping(); err != nil {
		return Registry{}, errors.Wrap(err, "can not connect db.")
	}
	bookmarkHandler := handler.NewBookmarkHandler(service.NewBookmarkService(model.NewBookmarkModel(db)))
	return Registry{db, bookmarkHandler}, nil
}
