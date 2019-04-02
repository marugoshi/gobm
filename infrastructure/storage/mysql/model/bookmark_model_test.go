package model_test

import (
	"context"
	"github.com/marugoshi/gobm/domain/entity"
	"github.com/marugoshi/gobm/infrastructure/storage/mysql"
	"github.com/marugoshi/gobm/infrastructure/storage/mysql/model"
	. "github.com/marugoshi/gobm/shared/app_testutils"
	"testing"
)

func TestBookmarkModel_All(t *testing.T) {
	tests := map[string]struct {
		queries []string
		tables  []string
		page    int
		perPage int
		results []*entity.Bookmark
	}{
		"no bookmark": {
			queries: []string{},
			tables:  []string{},
			page:    1,
			perPage: 1,
			results: []*entity.Bookmark{},
		},
		"bookmarks": {
			queries: []string{
				`INSERT INTO bookmarks (directory_id, url, title, created_at, updated_at) VALUES
				(1, "https://a.makototokuyama.com", "test1", '2018-01-01 00:00:00', '2018-01-01 00:00:00')`,
				`INSERT INTO bookmarks (directory_id, url, title, created_at, updated_at) VALUES
				(1, "https://b.makototokuyama.com", "test2", '2018-01-01 00:00:00', '2018-01-01 00:00:00')`,
				`INSERT INTO bookmarks (directory_id, url, title, created_at, updated_at) VALUES
				(1, "https://c.makototokuyama.com", "test3", '2018-01-01 00:00:00', '2018-01-01 00:00:00')`,
				`INSERT INTO bookmarks (directory_id, url, title, created_at, updated_at) VALUES
				(1, "https://d.makototokuyama.com", "test4", '2018-01-01 00:00:00', '2018-01-01 00:00:00')`,
			},
			tables: []string{
				"bookmarks",
			},
			page:    2,
			perPage: 2,
			results: []*entity.Bookmark{
				&entity.Bookmark{
					Id:          3,
					DirectoryId: mysql.NullInt64{Int64: 1, Valid: true},
					Url:         "https://c.makototokuyama.com",
					Title:       "test3",
					CreatedAt:   mysql.NewNullDateTime(2018, 01, 01, 00, 00, 00),
					UpdatedAt:   mysql.NewNullDateTime(2018, 01, 01, 00, 00, 00),
				},
				&entity.Bookmark{
					Id:          4,
					DirectoryId: mysql.NullInt64{Int64: 1, Valid: true},
					Url:         "https://d.makototokuyama.com",
					Title:       "test4",
					CreatedAt:   mysql.NewNullDateTime(2018, 01, 01, 00, 00, 00),
					UpdatedAt:   mysql.NewNullDateTime(2018, 01, 01, 00, 00, 00),
				},
			},
		},
	}

	for name, tt := range tests {
		tm := NewTestMySql(t)
		tm.Fixtures(tt.queries)
		m := model.NewBookmarkModel(tm.GetInstance())

		bookmarks, err := m.All(context.TODO(), tt.page, tt.perPage)
		if err != nil {
			t.Errorf("%s: got error: %v", name, err)
		}

		if len(bookmarks) != len(tt.results) {
			t.Errorf("%s: got %d, want %d", name, len(bookmarks), len(tt.results))
		}

		for i, bookmark := range bookmarks {
			if bookmark.Id != tt.results[i].Id {
				t.Errorf("%s: got %d, want %d", name, bookmark.Id, tt.results[i].Id)
			}

			if bookmark.DirectoryId.Valid != tt.results[i].DirectoryId.Valid {
				t.Errorf("%s: got %t, want %t",
					name,
					bookmark.DirectoryId.Valid,
					tt.results[i].DirectoryId.Valid)
			}

			if bookmark.DirectoryId.Int64 != tt.results[i].DirectoryId.Int64 {
				t.Errorf("%s: got %d, want %d",
					name,
					bookmark.DirectoryId.Int64,
					tt.results[i].DirectoryId.Int64)
			}

			if bookmark.Url != tt.results[i].Url {
				t.Errorf("%s: got %s, want %s", name, bookmark.Url, tt.results[i].Url)
			}

			if bookmark.Title != tt.results[i].Title {
				t.Errorf("%s: got %s, want %s", name, bookmark.Title, tt.results[i].Title)
			}

			if bookmark.CreatedAt.Valid != tt.results[i].CreatedAt.Valid {
				t.Errorf("%s: got %t, want %t",
					name,
					bookmark.CreatedAt.Valid,
					tt.results[i].CreatedAt.Valid)
			}

			if bookmark.CreatedAt.Time.String() != tt.results[i].CreatedAt.Time.String() {
				t.Errorf("%s: got %s, want %s",
					name,
					bookmark.CreatedAt.Time.String(),
					tt.results[i].CreatedAt.Time.String())
			}
		}

		tm.Truncates(tt.tables)
	}
}

func TestBookmarkModel_FindById(t *testing.T) {
	tests := map[string]struct {
		queries []string
		tables  []string
		id      int64
		result  *entity.Bookmark
	}{
		"no bookmark": {
			queries: []string{},
			tables:  []string{},
			id:      0,
			result:  nil,
		},
		"bookmark exists": {
			queries: []string{
				`INSERT INTO bookmarks (directory_id, url, title, created_at, updated_at) VALUES
				(1, "https://a.makototokuyama.com", "test1", '2018-01-01 00:00:00', '2018-01-01 00:00:00')`,
			},
			tables: []string{
				"bookmarks",
			},
			id: 1,
			result: &entity.Bookmark{
				Id:          1,
				DirectoryId: mysql.NullInt64{Int64: 1, Valid: true},
				Url:         "https://a.makototokuyama.com",
				Title:       "test1",
				CreatedAt:   mysql.NewNullDateTime(2018, 01, 01, 00, 00, 00),
				UpdatedAt:   mysql.NewNullDateTime(2018, 01, 01, 00, 00, 00),
			},
		},
	}

	for name, tt := range tests {
		tm := NewTestMySql(t)
		tm.Fixtures(tt.queries)
		m := model.NewBookmarkModel(tm.GetInstance())

		bookmark, err := m.FindById(context.TODO(), tt.id)
		if err != nil {
			t.Errorf("%s: got error: %v", name, err)
		}

		if bookmark != nil {
			if bookmark.Id != tt.result.Id {
				t.Errorf("%s: got %d, want %d", name, bookmark.Id, tt.result.Id)
			}

			if bookmark.DirectoryId.Valid != tt.result.DirectoryId.Valid {
				t.Errorf("%s: got %t, want %t",
					name,
					bookmark.DirectoryId.Valid,
					tt.result.DirectoryId.Valid)
			}

			if bookmark.DirectoryId.Int64 != tt.result.DirectoryId.Int64 {
				t.Errorf("%s: got %d, want %d",
					name,
					bookmark.DirectoryId.Int64,
					tt.result.DirectoryId.Int64)
			}

			if bookmark.Url != tt.result.Url {
				t.Errorf("%s: got %s, want %s", name, bookmark.Url, tt.result.Url)
			}

			if bookmark.Title != tt.result.Title {
				t.Errorf("%s: got %s, want %s", name, bookmark.Title, tt.result.Title)
			}

			if bookmark.CreatedAt.Valid != tt.result.CreatedAt.Valid {
				t.Errorf("%s: got %t, want %t",
					name,
					bookmark.CreatedAt.Valid,
					tt.result.CreatedAt.Valid)
			}

			if bookmark.CreatedAt.Time.String() != tt.result.CreatedAt.Time.String() {
				t.Errorf("%s: got %s, want %s",
					name,
					bookmark.CreatedAt.Time.String(),
					tt.result.CreatedAt.Time.String())
			}
		}

		tm.Truncates(tt.tables)
	}
}

func TestBookmarkModel_Create(t *testing.T) {
	tests := map[string]struct {
		tables  []string
		params  *entity.Bookmark
		isError bool
		Id      int64
	}{
		"success": {
			tables: []string{
				"bookmarks",
			},
			params: &entity.Bookmark{
				Url:         "https://a.makototokuyama.com",
				DirectoryId: mysql.NullInt64{Int64: 0, Valid: false},
				Title:       "a",
			},
			isError: false,
			Id:      1,
		},
	}

	for name, tt := range tests {
		tm := NewTestMySql(t)
		m := model.NewBookmarkModel(tm.GetInstance())

		id, err := m.Create(context.TODO(), tt.params)
		if tt.isError && err == nil {
			t.Errorf("%s: got %t, want %t", name, err == nil, tt.isError)
		}

		if !tt.isError {
			bookmark, err := m.FindById(context.TODO(), id)
			if err != nil {
				t.Errorf("%s: got error: %v", name, err)
			}

			if bookmark.Id != tt.Id {
				t.Errorf("%s: got %d, want %d", name, bookmark.Id, tt.Id)
			}

			if bookmark.Url != tt.params.Url {
				t.Errorf("%s: got %s, want %s", name, bookmark.Url, tt.params.Url)
			}

			if bookmark.DirectoryId.Int64 != tt.params.DirectoryId.Int64 {
				t.Errorf("%s: got %d, want %d", name, bookmark.DirectoryId.Int64, tt.params.DirectoryId.Int64)
			}

			if bookmark.Title != tt.params.Title {
				t.Errorf("%s: got %s, want %s", name, bookmark.Title, tt.params.Title)
			}
		}

		tm.Truncates(tt.tables)
	}
}

func TestBookmarkModel_Update(t *testing.T) {
	tests := map[string]struct {
		queries []string
		tables  []string
		params  *entity.Bookmark
		isError bool
		results []*entity.Bookmark
	}{
		"success": {
			queries: []string{
				`INSERT INTO bookmarks (directory_id, url, title, created_at, updated_at) VALUES
				(1, "https://a.makototokuyama.com", "test1", '2018-01-01 00:00:00', '2018-01-01 00:00:00')`,
				`INSERT INTO bookmarks (directory_id, url, title, created_at, updated_at) VALUES
				(1, "https://b.makototokuyama.com", "test2", '2018-02-01 00:00:00', '2018-02-01 00:00:00')`,
			},
			tables: []string{
				"bookmarks",
			},
			params: &entity.Bookmark{
				Id:          1,
				Url:         "https://z.makototokuyama.com",
				DirectoryId: mysql.NullInt64{Int64: 2, Valid: true},
				Title:       "testz",
			},
			isError: false,
			results: []*entity.Bookmark{
				&entity.Bookmark{
					Id:          1,
					DirectoryId: mysql.NullInt64{Int64: 2, Valid: true},
					Url:         "https://z.makototokuyama.com",
					Title:       "testz",
					CreatedAt:   mysql.NewNullDateTime(2018, 01, 01, 00, 00, 00),
				},
				&entity.Bookmark{
					Id:          2,
					DirectoryId: mysql.NullInt64{Int64: 1, Valid: true},
					Url:         "https://b.makototokuyama.com",
					Title:       "test2",
					CreatedAt:   mysql.NewNullDateTime(2018, 02, 01, 00, 00, 00),
				},
			},
		},
	}

	for name, tt := range tests {
		tm := NewTestMySql(t)
		tm.Fixtures(tt.queries)
		m := model.NewBookmarkModel(tm.GetInstance())

		id, err := m.Update(context.TODO(), tt.params)
		if tt.isError && err == nil {
			t.Errorf("%s: got %t, want %t", name, err == nil, tt.isError)
		}

		if !tt.isError {
			if err != nil {
				t.Errorf("%s: got error %v", name, err)
			}

			if id != tt.params.Id {
				t.Errorf("%s: got %d, want %d", name, id, tt.params.Id)
			}

			bookmarks, err := m.All(context.TODO(), 1, 2)
			if err != nil {
				t.Errorf("%s: got error: %v", name, err)
			}

			if len(bookmarks) != len(tt.results) {
				t.Errorf("%s: got %d, want %d", name, len(bookmarks), len(tt.results))
			}

			for i, bookmark := range bookmarks {
				if bookmark.Id != tt.results[i].Id {
					t.Errorf("%s: got %d, want %d", name, bookmark.Id, tt.results[i].Id)
				}

				if bookmark.DirectoryId.Valid != tt.results[i].DirectoryId.Valid {
					t.Errorf("%s: got %t, want %t",
						name,
						bookmark.DirectoryId.Valid,
						tt.results[i].DirectoryId.Valid)
				}

				if bookmark.DirectoryId.Int64 != tt.results[i].DirectoryId.Int64 {
					t.Errorf("%s: got %d, want %d",
						name,
						bookmark.DirectoryId.Int64,
						tt.results[i].DirectoryId.Int64)
				}

				if bookmark.Url != tt.results[i].Url {
					t.Errorf("%s: got %s, want %s", name, bookmark.Url, tt.results[i].Url)
				}

				if bookmark.Title != tt.results[i].Title {
					t.Errorf("%s: got %s, want %s", name, bookmark.Title, tt.results[i].Title)
				}

				if bookmark.CreatedAt.Valid != tt.results[i].CreatedAt.Valid {
					t.Errorf("%s: got %t, want %t",
						name,
						bookmark.CreatedAt.Valid,
						tt.results[i].CreatedAt.Valid)
				}

				if bookmark.CreatedAt.Time.String() != tt.results[i].CreatedAt.Time.String() {
					t.Errorf("%s: got %s, want %s",
						name,
						bookmark.CreatedAt.Time.String(),
						tt.results[i].CreatedAt.Time.String())
				}
			}
		}

		tm.Truncates(tt.tables)
	}
}

func TestBookmarkModel_Delete(t *testing.T) {
	tests := map[string]struct {
		queries []string
		tables  []string
		isError bool
		id      int64
		results []*entity.Bookmark
	}{
		"success": {
			queries: []string{
				`INSERT INTO bookmarks (directory_id, url, title, created_at, updated_at) VALUES
				(1, "https://a.makototokuyama.com", "test1", '2018-01-01 00:00:00', '2018-01-01 00:00:00')`,
				`INSERT INTO bookmarks (directory_id, url, title, created_at, updated_at) VALUES
				(1, "https://b.makototokuyama.com", "test2", '2018-02-01 00:00:00', '2018-02-01 00:00:00')`,
			},
			tables: []string{
				"bookmarks",
			},
			id: 1,
			results: []*entity.Bookmark{
				&entity.Bookmark{
					Id:          2,
					DirectoryId: mysql.NullInt64{Int64: 1, Valid: true},
					Url:         "https://b.makototokuyama.com",
					Title:       "test2",
					CreatedAt:   mysql.NewNullDateTime(2018, 02, 01, 00, 00, 00),
				},
			},
		},
	}

	for name, tt := range tests {
		tm := NewTestMySql(t)
		tm.Fixtures(tt.queries)
		m := model.NewBookmarkModel(tm.GetInstance())

		err := m.Delete(context.TODO(), tt.id)
		if tt.isError && err == nil {
			t.Errorf("%s: got %t, want %t", name, err == nil, tt.isError)
		}

		if !tt.isError {
			if err != nil {
				t.Errorf("%s: got error %v", name, err)
			}

			bookmarks, err := m.All(context.TODO(), 1, 2)
			if err != nil {
				t.Errorf("%s: got error: %v", name, err)
			}

			if len(bookmarks) != len(tt.results) {
				t.Errorf("%s: got %d, want %d", name, len(bookmarks), len(tt.results))
			}

			for i, bookmark := range bookmarks {
				if bookmark.Id != tt.results[i].Id {
					t.Errorf("%s: got %d, want %d", name, bookmark.Id, tt.results[i].Id)
				}

				if bookmark.DirectoryId.Valid != tt.results[i].DirectoryId.Valid {
					t.Errorf("%s: got %t, want %t",
						name,
						bookmark.DirectoryId.Valid,
						tt.results[i].DirectoryId.Valid)
				}

				if bookmark.DirectoryId.Int64 != tt.results[i].DirectoryId.Int64 {
					t.Errorf("%s: got %d, want %d",
						name,
						bookmark.DirectoryId.Int64,
						tt.results[i].DirectoryId.Int64)
				}

				if bookmark.Url != tt.results[i].Url {
					t.Errorf("%s: got %s, want %s", name, bookmark.Url, tt.results[i].Url)
				}

				if bookmark.Title != tt.results[i].Title {
					t.Errorf("%s: got %s, want %s", name, bookmark.Title, tt.results[i].Title)
				}

				if bookmark.CreatedAt.Valid != tt.results[i].CreatedAt.Valid {
					t.Errorf("%s: got %t, want %t",
						name,
						bookmark.CreatedAt.Valid,
						tt.results[i].CreatedAt.Valid)
				}

				if bookmark.CreatedAt.Time.String() != tt.results[i].CreatedAt.Time.String() {
					t.Errorf("%s: got %s, want %s",
						name,
						bookmark.CreatedAt.Time.String(),
						tt.results[i].CreatedAt.Time.String())
				}
			}
		}


		tm.Truncates(tt.tables)
	}
}