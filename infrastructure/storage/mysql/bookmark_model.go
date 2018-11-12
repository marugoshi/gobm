package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/marugoshi/gobm/domain/data"
	"github.com/marugoshi/gobm/domain/model"
)

var (
	id           int64
	directory_id sql.NullInt64
	url          string
	title        string
)

type BookmarkModel struct {
	*sql.DB
}

func NewBookmarkModel(db *sql.DB) model.BookmarkModel {
	return &BookmarkModel{db}
}

func (b *BookmarkModel) All(ctx context.Context, page int, perPage int) (interface{}, error) {
	rows, err := b.DB.Query("SELECT * FROM bookmarks")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	records := make([]*data.Bookmark, 0)
	for rows.Next() {
		if err := rows.Scan(&id, &directory_id, &url, &title); err != nil {
			return nil, err
		}
		records = append(records, &data.Bookmark{id, directory_id, url, title})
	}

	data := struct {
		Records []*data.Bookmark
	}{
		records,
	}

	return data, nil
}

func (b *BookmarkModel) FindById(ctx context.Context, id int64) (interface{}, error) {
	if err := b.DB.QueryRow("SELECT * FROM bookmarks WHERE id = ?", id).Scan(&id, &directory_id, &url, &title); err != nil {
		return nil, err
	}

	record := &data.Bookmark{id, directory_id, url, title}
	return record, nil
}

func (b *BookmarkModel) Create(ctx context.Context, params *data.Bookmark) (interface{}, error) {
	result, err := b.DB.Exec("INSERT INTO bookmarks (url, directory_id, title) VALUES(?, ?, ?)", params.Url, params.DirectoryId, params.Title)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return b.FindById(ctx, id)
}

func (b *BookmarkModel) Update(ctx context.Context, params *data.Bookmark) (interface{}, error) {
	tx, _ := b.DB.Begin()
	_, err := tx.Exec("UPDATE bookmarks SET directory_id = ?, url = ?, title = ? WHERE id = ?", params.DirectoryId, params.Url, params.Title, params.Id)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return b.FindById(ctx, params.Id)
}

func (b *BookmarkModel) Delete(ctx context.Context, id int64) error {
	tx, _ := b.DB.Begin()
	_, err := tx.Exec("DELETE FROM bookmarks WHERE id = ?", id)
	if err != nil {
		return nil
	}
	tx.Commit()
	return nil
}
