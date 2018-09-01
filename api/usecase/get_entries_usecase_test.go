package usecase_test

import (
	"errors"
	"testing"

	"github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"
	"github.com/Khigashiguchi/khigashiguchi.com/api/usecase"
	"github.com/google/go-cmp/cmp"
)

func TestGetEntriesUseCase_Run(t *testing.T) {
	tests := []struct {
		name           string
		mGetAll        func() ([]entity.Entry, error)
		expectedEntity []entity.Entry
		expectedErr    string
	}{
		{
			name: "get_no_entity",
			mGetAll: func() ([]entity.Entry, error) {
				return []entity.Entry{}, nil
			},
			expectedEntity: []entity.Entry{},
			expectedErr:    "",
		},
		{
			name: "get_1_entity",
			mGetAll: func() ([]entity.Entry, error) {
				return []entity.Entry{
					{
						Title: "test",
						URL:   "http://example.com",
					},
				}, nil
			},
			expectedEntity: []entity.Entry{
				{
					Title: "test",
					URL:   "http://example.com",
				},
			},
			expectedErr: "",
		},
		{
			name: "get_error",
			mGetAll: func() ([]entity.Entry, error) {
				return nil, errors.New("test error")
			},
			expectedEntity: nil,
			expectedErr:    "test error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mockEntriesRepo{mGetAll: tt.mGetAll}

			u := usecase.ExportGetEntriesUseCase{EntryRepo: &m}
			entries, err := u.Run()

			if err == nil {
				err = errors.New("")
			}
			if tt.expectedErr != err.Error() {
				t.Errorf("expected error: %#v\n, given %#v\n", tt.expectedErr, err.Error())
			}
			if diff := cmp.Diff(tt.expectedEntity, entries); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

type mockEntriesRepo struct {
	mGetAll func() ([]entity.Entry, error)
}

func (m *mockEntriesRepo) GetAll() ([]entity.Entry, error) {
	return m.mGetAll()
}
