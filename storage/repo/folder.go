package repo

import (
	"context"
	pb "ucode_go_query_service/genproto/query_service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type FolderRepoI interface {
	Create(ctx context.Context, in *pb.CreateFolderReq) (*pb.Folder, error)
	Update(ctx context.Context, in *pb.UpdateFolderReq) (*pb.Folder, error)
	GetSingle(ctx context.Context, in *pb.GetSingleFolderReq) (*pb.GetSingleFolderRes, error)
	GetList(ctx context.Context, in *pb.GetListFolderReq) (*pb.GetListFolderRes, error)
	Delete(ctx context.Context, in *pb.DeleteFolderReq) (*emptypb.Empty, error)
}
