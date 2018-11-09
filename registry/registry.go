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

func NewRegistry() Registry {
	db, _ := sql.Open("mysql", "root@/gobm_d")
	bookmarkHandler := handler.NewBookmarkHandler(service.NewBookmarkService(mysql.NewBookmarkModel(db)))
	return Registry{db, bookmarkHandler}
}
