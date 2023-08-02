package clickhouse

import (
	"ucode_go_query_service/storage/repo"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type fieldRepo struct {
	db *sqlx.DB
}

var _ repo.FieldCHI = &fieldRepo{}

func NewFieldCHRepo(db *sqlx.DB) repo.FieldCHI {
	return &fieldRepo{
		db: db,
	}
}

func (f *fieldRepo) Create(query string) error {
	_, err := f.db.Exec(query)

	if err != nil {
		return errors.Wrap(err, "error while creating field")
	}

	return nil
}

func (f *fieldRepo) Update(query string) error {
	_, err := f.db.Exec(query)

	if err != nil {
		return errors.Wrap(err, "error while updating field")
	}

	return nil
}

func (f *fieldRepo) Delete(query string) error {
	_, err := f.db.Exec(query)

	if err != nil {
		return errors.Wrap(err, "error while deleting field")
	}

	return nil
}

// just comment
