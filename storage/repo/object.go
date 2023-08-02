package repo

import (
	"ucode_go_query_service/models"
)

// table clickhouse interface

type ObjectCHI interface {
	Create(query string, args []interface{}) error
	Delete(query string) error
	Update(query string, args []interface{}) error
	GetQuery(input models.CommonInput) ([]map[string]interface{}, error)
}
