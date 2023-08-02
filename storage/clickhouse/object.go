package clickhouse

import (
	"fmt"
	"ucode_go_query_service/models"
	"ucode_go_query_service/storage/repo"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type objectRepo struct {
	db *sqlx.DB
}

var _ repo.ObjectCHI = &objectRepo{}

func NewObjectCHRepo(db *sqlx.DB) repo.ObjectCHI {
	return &objectRepo{
		db: db,
	}
}

func (o *objectRepo) Create(query string, args []interface{}) error {
	tx, err := o.db.Begin()
	if err != nil {
		return errors.Wrap(err, "error beginning transaction")
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	stmt, err := tx.Prepare(query)

	if err != nil {
		return errors.Wrap(err, "error while getting statement")
	}

	defer stmt.Close()

	_, err = stmt.Exec(args...)

	return err
}

func (o *objectRepo) Delete(query string) error {
	_, err := o.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (o *objectRepo) GetQuery(input models.CommonInput) ([]map[string]interface{}, error) {
	stmt, err := o.db.PrepareNamed(input.Query)
	if err != nil {
		return nil, errors.Wrap(err, "error while preparing statement")
	}

	rows, err := stmt.Query(input.Data)
	if err != nil {
		return nil, errors.Wrap(err, "error while getting query rows")
	}
	defer rows.Close()
	columns, err := rows.Columns()

	// for each database row / record, a map with the column names and row values is added to the allMaps slice
	var allMaps []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}
		err = rows.Scan(pointers...)
		resultMap := make(map[string]interface{})
		for i, val := range values {
			fmt.Printf("Adding key=%s val=%v\n", columns[i], val)
			resultMap[columns[i]] = val
		}
		allMaps = append(allMaps, resultMap)
	}

	return allMaps, err
}

func (o *objectRepo) Update(query string, args []interface{}) error {

	stmt, err := o.db.Prepare(query)

	if err != nil {
		return errors.Wrap(err, "error while getting statement")
	}

	defer stmt.Close()

	_, err = stmt.Exec(args...)

	return err
}
