package models

import "go.mongodb.org/mongo-driver/bson"

type LogSchema struct {
	Id            string `json:"id" bson:"id"`
	ProjectId     string `json:"project_id" bson:"project_id"`
	EnvironmentId string `json:"environment_id" bson:"environment_id"`
	QueryId       string `json:"query_id" bson:"query_id"`
	UserId        string `json:"user_id" bson:"user_id"`
	Request       bson.M `json:"request" bson:"request"`
	Response      string `json:"response" bson:"response"`
	CreatedAt     string `json:"created_at" bson:"created_at"`
}
