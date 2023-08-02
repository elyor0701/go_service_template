package service

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ucode_go_query_service/config"
	pb "ucode_go_query_service/genproto/query_service"
	"ucode_go_query_service/pkg/logger"
	"ucode_go_query_service/storage"
)

// LogService ...
type LogService struct {
	cfg    config.Config
	logger logger.Logger
	strg   storage.StorageI
	pb.UnimplementedLogServiceServer
}

// NewLogService ...
func NewLogService(cfg config.Config, log logger.Logger, strg storage.StorageI) *LogService {
	return &LogService{
		cfg:    cfg,
		logger: log,
		strg:   strg,
	}
}

func (s *LogService) GetSingleLog(ctx context.Context, in *pb.GetSingleLogReq) (*pb.Log, error) {
	s.logger.Info("--GetSingleLog-- requested", logger.Any("req", in))

	strg, err := s.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		s.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := strg.Log().GetSingle(ctx, in)
	if err != nil {
		errCreate := errors.New("error on getting query log")
		s.logger.Error("Error while getting query log", logger.Error(err))
		return nil, status.Error(codes.Internal, errCreate.Error())
	}

	return response, nil
}

func (s *LogService) GetListLog(ctx context.Context, in *pb.GetListLogReq) (*pb.GetListLogRes, error) {
	s.logger.Info("--GetListLog-- requested", logger.Any("req", in))

	strg, err := s.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		s.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := strg.Log().GetList(ctx, in)
	if err != nil {
		errCreate := errors.New("error on getting query logs")
		s.logger.Error("Error while getting query logs", logger.Error(err))
		return nil, status.Error(codes.Internal, errCreate.Error())
	}

	return response, nil
}
