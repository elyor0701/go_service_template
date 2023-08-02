package mongo

import (
	"context"
	"fmt"
	"log"
	"time"
	"ucode_go_query_service/models"
	"ucode_go_query_service/storage/repo"

	pb "ucode_go_query_service/genproto/query_service"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommitRepo struct {
	collection *mongo.Collection
}

func NewCommitRepo(db *mongo.Database) repo.CommitRepoI {
	return &CommitRepo{
		collection: db.Collection("queries.commit"),
	}
}

func (asr *CommitRepo) Create(ctx context.Context, in *models.CommitSchema) (resp *models.CommitSchema, err error) {
	log.Printf("-->STRG->Create commit---> %+v", in)

	id := uuid.New().String()

	in.Id = id
	in.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	in.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	_, err = asr.collection.InsertOne(ctx, in)
	if err != nil {
		return nil, errors.Wrap(err, "error inserting commit")
	}

	return in, nil
}

func (asr *CommitRepo) UpdateVersions(ctx context.Context, in *pb.CommitWithRelease) (resp *pb.CommitWithRelease, err error) {
	log.Printf("-->STRG->Update commit---> %+v", in)

	_, err = asr.collection.UpdateOne(ctx, bson.M{"id": in.GetId()}, bson.M{"$set": bson.M{"version_ids": in.GetReleaseInfo().GetIds()}})
	if err != nil {
		return nil, errors.Wrap(err, "error updating commit")
	}

	return in, nil
}

func (asr *CommitRepo) RemoveVersions(ctx context.Context, in *pb.CommitWithRelease) (resp *pb.CommitWithRelease, err error) {
	log.Printf("-->STRG->RemoveVersions commit---> %+v", in)

	res, err := asr.collection.UpdateMany(
		ctx,
		bson.M{"version_ids": bson.M{"$in": in.GetReleaseInfo().GetIds()}},
		bson.M{"$pull": bson.M{"version_ids": bson.M{"$in": in.GetReleaseInfo().GetIds()}}},
	)
	if err != nil {
		return nil, errors.Wrap(err, "error updating commit")
	}

	fmt.Println("res:::::", res.MatchedCount, res.UpsertedCount, res.ModifiedCount)

	return in, nil
}

func (asr *CommitRepo) GetByID(ctx context.Context, in *pb.CommitPrimaryKey) (resp *models.CommitSchema, err error) {
	log.Printf("-->STRG->Get commit by id---> %+v", in)

	err = asr.collection.FindOne(ctx, bson.M{
		"id": in.GetId(),
	}).Decode(&resp)
	if err != nil {
		return nil, errors.Wrap(err, "error getting commit")
	}

	return resp, nil
}

func (asr *CommitRepo) GetList(ctx context.Context, in *pb.GetCommitListRequest) (resp *pb.GetCommitListResponse, err error) {
	log.Printf("-->STRG->Get commit list---> %+v", in)

	return
}

func (asr *CommitRepo) Restore(ctx context.Context, in *pb.RestoreCommitRequest) (resp *emptypb.Empty, err error) {
	log.Printf("-->STRG->Restore commit---> %+v", in)

	return
}

func (asr *CommitRepo) Insert(ctx context.Context, in *pb.CreateCommitRequest) (resp *pb.InsertCommitResponse, err error) {
	log.Printf("-->STRG->Insert commit---> %+v", in)

	return
}

func (asr *CommitRepo) GetMultipleCommitInfo(ctx context.Context, in *pb.GetMultipleCommitInfoRequest) (resp *pb.GetMultipleCommitInfoResponse, err error) {
	log.Printf("-->STRG->Get multiple commit info---> %+v", in)

	return
}

func (asr *CommitRepo) GetCommitByVersionId(ctx context.Context, in string) (resp *pb.CommitWithRelease, err error) {
	log.Printf("-->STRG->GetCommitByVersionId---> %+v", in)

	var (
		commitData models.CommitSchema
	)

	err = asr.collection.FindOne(ctx, bson.M{"version_ids": bson.M{"$in": []string{in}}}).Decode(&commitData)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, errors.Wrap(err, "error getting commit")
	}

	return &pb.CommitWithRelease{
		Id:         commitData.Id,
		ProjectId:  commitData.ProjectId,
		CommitType: commitData.CommitType,
		AuthorId:   commitData.AuthorId,
		Name:       commitData.Name,
		CreatedAt:  commitData.CreatedAt,
		ReleaseInfo: &pb.CommitWithRelease_Release{
			Ids: commitData.VersionIds,
		},
	}, nil
}
