package clickhouse

import (
	"ucode_go_query_service/storage/repo"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type tableRepo struct {
	db *sqlx.DB
}

var _ repo.TableCHI = &tableRepo{}

func NewTableCHRepo(db *sqlx.DB) repo.TableCHI {
	return &tableRepo{
		db: db,
	}
}

func (t *tableRepo) Create(query string) error {
	_, err := t.db.Exec(query)

	if err != nil {
		return errors.Wrap(err, "error while creating table")
	}

	return nil
}

func (t *tableRepo) Delete(query string) error {
	_, err := t.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (t *tableRepo) Update(query string) error {
	_, err := t.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
