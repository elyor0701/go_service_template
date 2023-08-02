package repo

// relation clickhouse interface

type RelationCHI interface {
	Create(query string) error
	Delete(query string) error
}
