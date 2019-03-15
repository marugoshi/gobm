package registry

import (
	"database/sql"
	"github.com/marugoshi/gobm/domain/service"
	"github.com/marugoshi/gobm/infrastructure/storage/mysql"
	"github.com/marugoshi/gobm/presentation/handler"
)

type Registry struct {
	*sql.DB
	handler.BookmarkHandler
}

func NewRegistry() (Registry, error) {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/gobm_d?parseTime=true&loc=Asia%%2FTokyo")
	if err != nil {
		return Registry{}, err
	}
	bookmarkHandler := handler.NewBookmarkHandler(service.NewBookmarkService(mysql.NewBookmarkModel(db)))
	return Registry{db, bookmarkHandler}, nil
}
