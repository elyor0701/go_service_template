package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"ucode_go_query_service/storage/repo"
)

type StorageMI interface {
	Commit() repo.CommitRepoI
	Folder() repo.FolderRepoI
	Query() repo.QueryRepoI
	Log() repo.LogRepoI
}

type storageM struct {
	commitRepo repo.CommitRepoI
	folderRepo repo.FolderRepoI
	queryRepo  repo.QueryRepoI
	logRepo    repo.LogRepoI
}

func NewStorageM(db *mongo.Database) StorageMI {
	return &storageM{
		commitRepo: NewCommitRepo(db),
		folderRepo: NewFolderRepo(db),
		queryRepo:  NewQueryRepo(db),
		logRepo:    NewLogRepo(db),
	}
}

func (s *storageM) Commit() repo.CommitRepoI {
	return s.commitRepo
}

func (s *storageM) Folder() repo.FolderRepoI {
	return s.folderRepo
}

func (s *storageM) Query() repo.QueryRepoI {
	return s.queryRepo
}

func (s *storageM) Log() repo.LogRepoI {
	return s.logRepo
}
