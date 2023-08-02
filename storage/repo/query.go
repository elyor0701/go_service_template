package repo

import (
	"context"
	pb "ucode_go_query_service/genproto/query_service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type QueryRepoI interface {
	Create(ctx context.Context, in *pb.CreateQueryReq) (*pb.Query, error)
	Update(ctx context.Context, in *pb.UpdateQueryReq) (*pb.Query, error)
	GetSingle(ctx context.Context, in *pb.GetSingleQueryReq) (*pb.Query, error)
	GetList(ctx context.Context, in *pb.GetListQueryReq) (*pb.GetListQueryRes, error)
	Delete(ctx context.Context, in *pb.DeleteQueryReq) (*emptypb.Empty, error)
	GetQueryHistory(context.Context, *pb.GetQueryHistoryReq) (*pb.GetQueryHistoryRes, error)
	Revert(context.Context, *pb.RevertQueryReq) (*pb.Query, error)
	//CreateManyApiReference(ctx context.Context, req *pb.ManyVersions) (resp *emptypb.Empty, err error)
}
