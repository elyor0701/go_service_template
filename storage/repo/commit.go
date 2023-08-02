package repo

import (
	"context"
	pb "ucode_go_query_service/genproto/query_service"
	"ucode_go_query_service/models"

	"github.com/golang/protobuf/ptypes/empty"
)

type CommitRepoI interface {
	Create(ctx context.Context, in *models.CommitSchema) (resp *models.CommitSchema, err error)
	GetByID(ctx context.Context, in *pb.CommitPrimaryKey) (resp *models.CommitSchema, err error)
	UpdateVersions(ctx context.Context, in *pb.CommitWithRelease) (*pb.CommitWithRelease, error)
	RemoveVersions(ctx context.Context, in *pb.CommitWithRelease) (*pb.CommitWithRelease, error)
	GetList(ctx context.Context, in *pb.GetCommitListRequest) (*pb.GetCommitListResponse, error)
	Restore(ctx context.Context, in *pb.RestoreCommitRequest) (*empty.Empty, error)
	Insert(ctx context.Context, in *pb.CreateCommitRequest) (*pb.InsertCommitResponse, error)
	GetMultipleCommitInfo(context.Context, *pb.GetMultipleCommitInfoRequest) (*pb.GetMultipleCommitInfoResponse, error)
	GetCommitByVersionId(ctx context.Context, in string) (resp *pb.CommitWithRelease, err error)
}
