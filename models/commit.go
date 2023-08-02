package models

type CommitSchema struct {
	Id         string   `json:"id" bson:"id"`
	ProjectId  string   `json:"project_id" bson:"project_id"`
	VersionIds []string `json:"version_ids" bson:"version_ids"`
	CommitType string   `json:"commit_type" bson:"commit_type"`
	AuthorId   string   `json:"author_id" bson:"author_id"`
	Name       string   `json:"name" bson:"name"`
	CreatedAt  string   `json:"created_at" bson:"created_at"`
	UpdatedAt  string   `json:"updated_at" bson:"updated_at"`
}
