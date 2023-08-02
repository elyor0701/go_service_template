package models

type OutboxEvent struct {
	Id            string `json:"id"`
	AggregateType string `json:"aggregate_type"`
	AggregateId   string `json:"aggregate_id"`
	Type          string `json:"type"`
	SessionId     string `json:"session_id"`
	ProjectId     string `json:"project_id"`
}

type ErrorModel struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}
