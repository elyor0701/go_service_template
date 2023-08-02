package service

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"ucode_go_query_service/config"
	pb "ucode_go_query_service/genproto/query_service"
	"ucode_go_query_service/pkg/logger"
	l "ucode_go_query_service/pkg/logger"
	"ucode_go_query_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FolderService ...
type FolderService struct {
	cfg    config.Config
	logger l.Logger
	strg   storage.StorageI
	pb.UnimplementedFolderServiceServer
}

// NewFolderService ...
func NewFolderService(cfg config.Config, log l.Logger, strg storage.StorageI) *FolderService {
	return &FolderService{
		cfg:    cfg,
		logger: log,
		strg:   strg,
	}
}

func (cs *FolderService) CreateFolder(ctx context.Context, in *pb.CreateFolderReq) (*pb.Folder, error) {
	cs.logger.Info("--Create Folder-- requested", logger.Any("req", in))

	strg, err := cs.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		cs.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := strg.Folder().Create(ctx, in)
	if err != nil {
		errCreate := errors.New("error on creating folder")
		cs.logger.Error("Error while creating Folder in service", l.Error(err))
		return nil, status.Error(codes.Internal, errCreate.Error())
	}

	return response, nil
}

func (cs *FolderService) UpdateFolder(ctx context.Context, in *pb.UpdateFolderReq) (*pb.Folder, error) {
	cs.logger.Info("--Update Folder-- requested", logger.Any("req", in))

	strg, err := cs.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		cs.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := strg.Folder().Update(ctx, in)
	if err != nil {
		cs.logger.Error("Error while updating Folder in service", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (cs *FolderService) GetSingleFolder(ctx context.Context, in *pb.GetSingleFolderReq) (*pb.GetSingleFolderRes, error) {
	cs.logger.Info("--Get Single Folder-- requested", logger.Any("req", in))

	strg, err := cs.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		cs.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := strg.Folder().GetSingle(ctx, in)
	if err != nil {
		cs.logger.Error("Error while getting single Folder in service", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (cs *FolderService) GetListFolder(ctx context.Context, in *pb.GetListFolderReq) (*pb.GetListFolderRes, error) {
	cs.logger.Info("--Get All Folder-- requested", logger.Any("req", in))

	strg, err := cs.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		cs.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.ProjectId == "" {
		cs.logger.Info("--Get All Folder-- Project ID required", logger.Any("req", in))
		return nil, status.Error(codes.Internal, "Project id is required")
	}

	response, err := strg.Folder().GetList(ctx, in)
	if err != nil {
		cs.logger.Error("Error while getting list Folder in service", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (cs *FolderService) DeleteFolder(ctx context.Context, in *pb.DeleteFolderReq) (*emptypb.Empty, error) {
	cs.logger.Info("--Delete Folder-- requested", logger.Any("req", in))

	strg, err := cs.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		cs.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := strg.Folder().Delete(ctx, in)
	if err != nil {
		cs.logger.Error("Error while deleting Folder in service", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}
