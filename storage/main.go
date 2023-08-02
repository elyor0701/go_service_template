package storage

import (
	"context"
	"fmt"
	"time"
	"ucode_go_query_service/config"
	"ucode_go_query_service/genproto/common_messages"
	"ucode_go_query_service/genproto/company_service"
	"ucode_go_query_service/grpc/client"
	"ucode_go_query_service/pkg/logger"
	"ucode_go_query_service/storage/clickhouse"
	"ucode_go_query_service/storage/mongo"
)

//type StorageI interface {
//	CloseDB()
//	Folder() repo.FolderRepoI
//	Query() repo.QueryRepoI
//	Commit() repo.CommitRepoI
//	Log() repo.LogRepoI
//}

type StorageI interface {
	PoolM() mongo.ResourceStoragesI
	PoolCH() clickhouse.ProjectStoragesI
}

type storage struct {
	poolM  mongo.ResourceStoragesI
	poolCH clickhouse.ProjectStoragesI
}

func NewStorage() StorageI {
	return &storage{
		poolM:  mongo.NewResourceStorages(),
		poolCH: clickhouse.NewProjectStorages(),
	}
}

func (s *storage) PoolM() mongo.ResourceStoragesI {
	return s.poolM
}

func (s *storage) PoolCH() clickhouse.ProjectStoragesI {
	return s.poolCH
}

func AutoConnect(svcs client.ServiceManagerI, strg StorageI) error {
	fmt.Println("---AutoConnect--->")

	ctx, finish := context.WithTimeout(context.Background(), 60*time.Second)
	defer finish()

	res, err := svcs.ResourceService().AutoConnect(ctx, &company_service.GetProjectsRequest{
		K8SNamespace: config.Load().K8sNamespace,
	})
	if err != nil {
		fmt.Println("--AutoConnect--", logger.Error(err))
		return err
	}

	for _, value := range res.GetRes() {
		if value.GetResourceType() == company_service.ResourceType_CLICKHOUSE {
			//fmt.Println("cred::::::", value)
			clickhouseDB, err := clickhouse.NewClickhouseConn(&common_messages.ClickhouseCredentials{
				Host:     value.GetCredentials().GetHost(),
				Port:     value.GetCredentials().GetPort(),
				Username: value.GetCredentials().GetUsername(),
				Password: value.GetCredentials().GetPassword(),
				Database: value.GetCredentials().GetDatabase(),
			})
			if err != nil {
				fmt.Println("--AutoConnect--", logger.Error(err), logger.Any("cred:::", value))
				continue
			}

			storageCH := clickhouse.NewStorageCH(clickhouseDB)
			_ = strg.PoolCH().Add(storageCH, value.GetId())
		} else if value.ResourceType == company_service.ResourceType_MONGODB {
			//fmt.Println("cred::::::", value)
			//fmt.Println("lalalaalalalalalalalaall")
			dbConn, err := mongo.NewMongoConn(&common_messages.MongodbCredentials{
				Host:     value.GetCredentials().GetHost(),
				Port:     value.GetCredentials().GetPort(),
				Username: value.GetCredentials().GetUsername(),
				Password: value.GetCredentials().GetPassword(),
				Database: value.GetCredentials().GetDatabase(),
			})
			if err != nil {
				fmt.Println("--AutoConnect--", logger.Error(err), logger.Any("cred:::", value))
				continue
			}

			storageM := mongo.NewStorageM(dbConn.Database(value.GetCredentials().GetDatabase()))
			_ = strg.PoolM().Add(storageM, value.GetId())
			//test, err := strg.PoolM().Get(value.GetId())
			//fmt.Println("lalalalaalalalalalal", test, err)
		}
	}

	return nil
}
