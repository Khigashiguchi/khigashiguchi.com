package usecase_test

//
//import (
//	"database/sql"
//	"testing"
//
//	"github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"
//	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/repository"
//	"github.com/Khigashiguchi/khigashiguchi.com/api/usecase"
//)
//
//func TestPostEntriesUseCase_Run(t *testing.T) {
//	tests := []struct {
//		name     string
//		input    entity.Entry
//		expected error
//	}{
//		{
//			name: "save_entry",
//			input: entity.Entry{
//				Title: "test title",
//				URL:   "http://localhost/entries",
//			},
//			expected: nil,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			u := usecase.ExportPostEntriesUseCase{
//				DB:        &mockBeginner{},
//				EntryRepo: &mockEntryRepo{},
//			}
//			res := u.Run(tt.input)
//			if tt.expected != res {
//				t.Errorf("expected error: %#v, given error: %#v", tt.expected, res)
//			}
//		})
//	}
//}
//
//type mockBeginner struct{}
//
//func (*mockBeginner) Begin() (repository.Tx, error) {
//	return &mockTx{}, nil
//}
//
//type mockTx struct{}
//
//func (*mockTx) Rollback() error {
//	panic("implement me")
//}
//
//func (*mockTx) Commit() error {
//	return nil
//}
//
//func (*mockTx) Exec(string, ...interface{}) (sql.Result, error) {
//	panic("implement me")
//}
//
//func (*mockTx) Query(string, ...interface{}) (*sql.Rows, error) {
//	panic("implement me")
//}
//
//func (*mockTx) QueryRow(string, ...interface{}) *sql.Row {
//	panic("implement me")
//}
//
//func (*mockTx) Prepare(string) (*sql.Stmt, error) {
//	panic("implement me")
//}
//
//type mockEntryRepo struct {
//}
//
//func (*mockEntryRepo) GetAll(db repository.DB) ([]entity.Entry, error) {
//	panic("implement me")
//}
//
//func (*mockEntryRepo) Save(db repository.DB, entry entity.Entry) error {
//	return nil
//}
