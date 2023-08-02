package service

import (
	"context"
	"ucode_go_query_service/config"
	"ucode_go_query_service/genproto/common_messages"
	"ucode_go_query_service/grpc/client"
	"ucode_go_query_service/storage"
	"ucode_go_query_service/storage/clickhouse"
	"ucode_go_query_service/storage/mongo"

	"ucode_go_query_service/pkg/logger"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "ucode_go_query_service/genproto/query_service"
)

type ProjectService struct {
	cfg      config.Config
	log      logger.Logger
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedProjectServiceServer
}

func NewProjectService(cfg config.Config, log logger.Logger, svcs client.ServiceManagerI, strg storage.StorageI) *ProjectService {
	return &ProjectService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *ProjectService) Register(ctx context.Context, req *pb.RegisterProjectRequest) (response *empty.Empty, err error) {
	s.log.Info("---Register--->", logger.Any("req", req))

	dbConn, err := mongo.NewMongoConn(&common_messages.MongodbCredentials{
		Host:     req.GetCredentials().GetHost(),
		Port:     req.GetCredentials().GetPort(),
		Username: req.GetCredentials().GetUsername(),
		Password: req.GetCredentials().GetPassword(),
		Database: req.GetCredentials().GetDatabase(),
	})
	if err != nil {
		s.log.Error("--Register--", logger.Error(err))
		return nil, err
	}

	strg := mongo.NewStorageM(dbConn.Database(req.GetCredentials().GetDatabase()))
	err = s.strg.PoolM().Add(strg, req.ResourceId)
	if err != nil {
		s.log.Error("--Register--", logger.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ProjectService) Deregister(ctx context.Context, req *pb.DeregisterProjectRequest) (response *empty.Empty, err error) {
	s.log.Info("---Deregister--->", logger.Any("req", req))

	err = s.strg.PoolM().Remove(req.ProjectId)
	if err != nil {
		s.log.Error("--Deregister--", logger.Error(err))
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *ProjectService) RegisterMany(ctx context.Context, req *pb.RegisterManyProjectsRequest) (response *pb.RegisterManyProjectsResponse, err error) {
	s.log.Info("---RegisterMany--->", logger.Any("req", req))

	return &pb.RegisterManyProjectsResponse{}, status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *ProjectService) DeregisterMany(ctx context.Context, req *pb.DeregisterManyProjectsRequest) (response *pb.DeregisterManyProjectsResponse, err error) {
	s.log.Info("---DeregisterMany--->", logger.Any("req", req))

	return &pb.DeregisterManyProjectsResponse{}, status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *ProjectService) Reconnect(ctx context.Context, req *pb.RegisterProjectRequest) (response *empty.Empty, err error) {
	s.log.Info("---Reconnect--->", logger.Any("req", req))

	if req.ResourceType == pb.ResourceType_MONGODB {
		dbConn, err := mongo.NewMongoConn(&common_messages.MongodbCredentials{
			Host:     req.GetCredentials().GetHost(),
			Port:     req.GetCredentials().GetPort(),
			Username: req.GetCredentials().GetUsername(),
			Password: req.GetCredentials().GetPassword(),
			Database: req.GetCredentials().GetDatabase(),
		})
		if err != nil {
			s.log.Error("--Reconnect--", logger.Error(err))
			return nil, err
		}

		strg := mongo.NewStorageM(dbConn.Database(req.GetCredentials().GetDatabase()))
		err = s.strg.PoolM().Add(strg, req.ResourceId)
		if err != nil {
			s.log.Error("--Reconnect--", logger.Error(err))
			return nil, err
		}
	} else if req.ResourceType == pb.ResourceType_CLICKHOUSE {
		dbConn, err := clickhouse.NewClickhouseConn(&common_messages.ClickhouseCredentials{
			Host:     req.GetCredentials().GetHost(),
			Port:     req.GetCredentials().GetPort(),
			Username: req.GetCredentials().GetUsername(),
			Password: req.GetCredentials().GetPassword(),
			Database: req.GetCredentials().GetDatabase(),
		})
		if err != nil {
			s.log.Error("--Reconnect--", logger.Error(err))
			return nil, err
		}

		strg := clickhouse.NewStorageCH(dbConn)
		err = s.strg.PoolCH().Add(strg, req.ResourceId)
		if err != nil {
			s.log.Error("--Reconnect--", logger.Error(err))
			return nil, err
		}
	}

	return &emptypb.Empty{}, nil
}
