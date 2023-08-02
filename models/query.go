package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type QuerySchema struct {
	Id            string            `json:"id" bson:"id"`
	Title         string            `json:"title" bson:"title"`
	Description   string            `json:"description" bson:"description"`
	Timeout       string            `json:"timeout" bson:"timeout"`
	ProjectId     string            `json:"project_id" bson:"project_id"`
	EnvironmentId string            `json:"environment_id" bson:"environment_id"`
	FolderId      string            `json:"folder_id" bson:"folder_id"`
	QueryType     string            `json:"query_type" bson:"query_type"`
	ResourceId    string            `json:"resource_id" bson:"resource_id"`
	Body          bson.M            `json:"body" bson:"body"`
	Variables     []*VariableSchema `json:"variables" bson:"variables"`
	CommitId      string            `json:"commit_id" bson:"commit_id"`
	VersionId     string            `json:"version_id" bson:"version_id"`
	CreatedAt     string            `json:"created_at" bson:"created_at"`
	UpdatedAt     string            `json:"updated_at" bson:"updated_at"`
	IsDeleted     bool              `json:"is_deleted" bson:"is_deleted"`
}

type VariableSchema struct {
	Id      string `json:"id" bson:"id"`
	Key     string `json:"key" bson:"key"`
	Value   string `json:"value" bson:"value"`
	Type    string `json:"type" bson:"type"`
	Default string `json:"default" bson:"default"`
	QueryId string `json:"query_id" bson:"query_id"`
}

type CommonInput struct {
	Query string                 `json:"query"`
	Data  map[string]interface{} `json:"data"`
}
