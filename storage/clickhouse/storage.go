package clickhouse

import (
	"github.com/jmoiron/sqlx"
	"ucode_go_query_service/storage/repo"
)

type StorageCHI interface {
	Table() repo.TableCHI
	Object() repo.ObjectCHI
	Field() repo.FieldCHI
	Relation() repo.RelationCHI
}

type storageCH struct {
	tableRepo    repo.TableCHI
	objectRepo   repo.ObjectCHI
	fieldRepo    repo.FieldCHI
	relationRepo repo.RelationCHI
}

func NewStorageCH(db *sqlx.DB) StorageCHI {
	return &storageCH{
		objectRepo:   NewObjectCHRepo(db),
		tableRepo:    NewTableCHRepo(db),
		fieldRepo:    NewFieldCHRepo(db),
		relationRepo: NewRelationCHRepo(db),
	}
}

func (s *storageCH) Table() repo.TableCHI {
	return s.tableRepo
}

func (s *storageCH) Object() repo.ObjectCHI {
	return s.objectRepo
}

func (s *storageCH) Field() repo.FieldCHI {
	return s.fieldRepo
}

func (s *storageCH) Relation() repo.RelationCHI {
	return s.relationRepo
}
