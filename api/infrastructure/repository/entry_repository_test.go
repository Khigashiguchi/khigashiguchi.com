package repository_test

import (
	"testing"

	"github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"
	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/repository"
	"github.com/google/go-cmp/cmp"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestEntryStore_GetAll(t *testing.T) {
	tests := []struct {
		name             string
		expectedEntities []entity.Entry
		expectedErr      error
	}{
		{
			name: "get_records",
			expectedEntities: []entity.Entry{
				{
					Title: "ECS(Fargate)で動かすコンテナにSSMからクレデンシャル情報を渡す",
					URL:   "http://khigashigashi.hatenablog.com/entry/2018/08/28/214417",
				},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			defer db.Close()
			columns := []string{"title", "url"}
			rows := sqlmock.NewRows(columns).
				AddRow("ECS(Fargate)で動かすコンテナにSSMからクレデンシャル情報を渡す", "http://khigashigashi.hatenablog.com/entry/2018/08/28/214417")
			expectedQuery := "SELECT title, url FROM entries"
			mock.ExpectQuery(expectedQuery).WillReturnRows(rows)

			s := repository.EntryRepository{}
			entries, err := s.GetAll(db)
			if tt.expectedErr != err {
				t.Errorf("expected err: %#v,\n given: %#v", tt.expectedErr, err)
			}
			if diff := cmp.Diff(tt.expectedEntities, entries); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func TestEntryRepository_Save(t *testing.T) {
	tests := []struct {
		name     string
		input    entity.Entry
		expected error
	}{
		{
			name: "save_1_record",
			input: entity.Entry{
				Title: "test title",
				URL:   "http://example.com",
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			defer db.Close()
			mock.ExpectPrepare("INSERT INTO").
				ExpectExec().
				WithArgs(tt.input.Title, tt.input.URL).
				WillReturnResult(sqlmock.NewResult(1, 1))

			r := repository.EntryRepository{}

			res := r.Save(db, tt.input)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			if tt.expected != res {
				t.Errorf("expected error: %#v, given error: %#v", tt.expected, res)
			}
		})
	}
}
