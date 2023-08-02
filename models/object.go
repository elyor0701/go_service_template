package models

type Object struct {
	Data       map[string]interface{} `json:"data"`
	FieldTypes map[string]interface{} `json:"field_types"`
	TableSlug  string                 `json:"table_slug"`
}

type ObjectOutboxEvent struct {
	OutboxEvent
	Payload Object `json:"payload"`
}

type DeleteObject struct {
	TableSlug string `json:"table_slug"`
	GUID      string `json:"guid"`
}

type ObjectDeleteEvent struct {
	OutboxEvent
	Payload DeleteObject `json:"payload"`
}
