package repo

// table clickhouse interface

type TableCHI interface {
	Create(query string) error
	Delete(query string) error
	Update(query string) error
}

// field clickhouse interface

type FieldCHI interface {
	Create(query string) error
	Update(query string) error
	Delete(query string) error
}
