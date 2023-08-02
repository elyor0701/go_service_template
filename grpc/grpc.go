package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"ucode_go_query_service/config"
	"ucode_go_query_service/genproto/query_service"
	"ucode_go_query_service/grpc/client"
	"ucode_go_query_service/grpc/service"
	"ucode_go_query_service/pkg/logger"
	"ucode_go_query_service/storage"
)

func SetUpServer(cfg config.Config, log logger.Logger, svcs client.ServiceManagerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	query_service.RegisterQueryServiceServer(grpcServer, service.NewQueryService(cfg, log, svcs, strg))
	query_service.RegisterProjectServiceServer(grpcServer, service.NewProjectService(cfg, log, svcs, strg))
	query_service.RegisterFolderServiceServer(grpcServer, service.NewFolderService(cfg, log, strg))
	query_service.RegisterLogServiceServer(grpcServer, service.NewLogService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
