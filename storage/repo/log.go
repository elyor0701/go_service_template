package repo

import (
	"context"
	pb "ucode_go_query_service/genproto/query_service"
	"ucode_go_query_service/models"
)

type LogRepoI interface {
	Create(ctx context.Context, in models.LogSchema) error
	GetSingle(ctx context.Context, in *pb.GetSingleLogReq) (*pb.Log, error)
	GetList(ctx context.Context, in *pb.GetListLogReq) (*pb.GetListLogRes, error)
}
