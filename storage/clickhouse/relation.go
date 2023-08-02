package clickhouse

import (
	"ucode_go_query_service/storage/repo"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type relationRepo struct {
	db *sqlx.DB
}

var _ repo.RelationCHI = &relationRepo{}

func NewRelationCHRepo(db *sqlx.DB) repo.RelationCHI {
	return &relationRepo{
		db: db,
	}
}

func (r *relationRepo) Create(query string) error {
	_, err := r.db.Exec(query)

	if err != nil {
		return errors.Wrap(err, "error while creating relation")
	}

	return nil
}

func (r *relationRepo) Delete(query string) error {
	_, err := r.db.Exec(query)

	if err != nil {
		return errors.Wrap(err, "error while deleting relation")
	}

	return nil
}
