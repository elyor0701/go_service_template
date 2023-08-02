package models

type Field struct {
	Default  string `json:"default"`
	Type     string `json:"type"`
	Slug     string `json:"slug"`
	Required bool   `json:"required"`
}

type Table struct {
	Slug   string  `json:"slug"`
	Fields []Field `json:"fields"`
}

type TableOutboxEvent struct {
	OutboxEvent
	Payload Table `json:"payload"`
}

type UpdateTable struct {
	OlderSlug string  `json:"older_slug"`
	Slug      string  `json:"slug"`
	Fields    []Field `json:"fields"`
}

type UpdateTableOutboxEvent struct {
	OutboxEvent
	Payload UpdateTable `json:"payload"`
}

type UpdateField struct {
	TableSlug      string `json:"table_slug"`
	OlderFieldName string `json:"older_field_name"`
	Field          Field  `json:"field"`
}

type FieldOutboxEvent struct {
	OutboxEvent
	Payload UpdateField `json:"payload"`
}
