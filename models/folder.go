package models

type FolderSchema struct {
	Id        string `json:"id" bson:"id"`
	Title     string `json:"title" bson:"title"`
	ProjectId string `json:"project_id" bson:"project_id"`
	ParentId  string `json:"parent_id" bson:"parent_id"`
	IsDeleted bool   `json:"is_deleted" bson:"is_deleted"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}
