package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"ucode_go_query_service/genproto/common_messages"
)

//	type Store struct {
//		db         *mongo.Client
//		cfg        config.Config
//		commitRepo repo.CommitRepoI
//		folderRepo repo.FolderRepoI
//		queryRepo  repo.QueryRepoI
//		logRepo    repo.LogRepoI
//		//companyStorageI repo.ApiReferenceRepoI
//		//environmentRepo repo.CommitRepoI
//		//projectRepo     repo.CategoryRepoI
//	}
func NewMongoConn(cred *common_messages.MongodbCredentials) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		cred.Username,
		cred.Password,
		cred.Host,
		cred.Port,
		cred.Database,
	)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, err
}

//
//func NewMongo(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
//	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
//		cfg.MongoUser,
//		cfg.MongoPassword,
//		cfg.MongoHost,
//		cfg.MongoPort,
//		cfg.MongoDB,
//	)
//	clientOptions := options.Client().ApplyURI(uri)
//
//	client, err := mongo.Connect(ctx, clientOptions)
//	if err != nil {
//		return nil, err
//	}
//
//	err = client.Ping(ctx, readpref.Primary())
//	if err != nil {
//		return nil, err
//	}
//
//	return &Store{
//		db:  client,
//		cfg: cfg,
//	}, err
//}
//
//func (s *Store) CloseDB() {
//	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
//	defer cancel()
//	err := s.db.Disconnect(ctx)
//	if err != nil {
//		fmt.Println("ERROR ON DISCONNECTING TO DB")
//		return
//	}
//}
//
//func (s *Store) Commit() repo.CommitRepoI {
//	if s.commitRepo == nil {
//		s.commitRepo = NewCommitRepo(s.db.Database(s.cfg.MongoDB))
//	}
//
//	return s.commitRepo
//}
//
//func (s *Store) Folder() repo.FolderRepoI {
//	if s.folderRepo == nil {
//		s.folderRepo = NewFolderRepo(s.db.Database(s.cfg.MongoDB))
//	}
//
//	return s.folderRepo
//}
//
//func (s *Store) Query() repo.QueryRepoI {
//	if s.queryRepo == nil {
//		s.queryRepo = NewQueryRepo(s.db.Database(s.cfg.MongoDB))
//	}
//
//	return s.queryRepo
//}
//
//func (s *Store) Log() repo.LogRepoI {
//	if s.logRepo == nil {
//		s.logRepo = NewLogRepo(s.db.Database(s.cfg.MongoDB))
//	}
//
//	return s.logRepo
//}
//
////
////func (s *Store) Project() repo.ProjectStorageI {
////	if s.projectRepo == nil {
////		s.projectRepo = NewProjectRepo(s.db)
////	}
////	return s.projectRepo
////}
////
////func (s *Store) Environment() repo.EnvironmentStorageI {
////	if s.environmentRepo == nil {
////		s.environmentRepo = NewEnvironmentRepo(s.db)
////	}
////	return s.environmentRepo
////}
////
////func (s *Store) Resource() repo.ResourceStorageI {
////	if s.resourceRepo == nil {
////		s.resourceRepo = NewResourceRepo(s.db)
////	}
////	return s.resourceRepo
////}
