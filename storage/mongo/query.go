package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
	"ucode_go_query_service/models"
	"ucode_go_query_service/pkg/helper"
	"ucode_go_query_service/storage/repo"

	pb "ucode_go_query_service/genproto/query_service"

	"go.mongodb.org/mongo-driver/mongo"
)

type QueryRepo struct {
	collections map[string]*mongo.Collection
}

func NewQueryRepo(db *mongo.Database) repo.QueryRepoI {
	return &QueryRepo{
		collections: map[string]*mongo.Collection{
			"query":  db.Collection("queries.query"),
			"commit": db.Collection("queries.commit"),
		},
	}
}

func (q *QueryRepo) Create(ctx context.Context, in *pb.CreateQueryReq) (*pb.Query, error) {
	var (
		variables []*models.VariableSchema
	)
	id := uuid.New().String()

	body, err := helper.ConvertStructToResponse(in.Body)
	if err != nil {
		return nil, err
	}

	err = helper.ConvertStructToStruct(in.GetVariables(), &variables)

	doc := models.QuerySchema{
		Id:            id,
		Title:         in.GetTitle(),
		Description:   in.GetDescription(),
		Timeout:       in.GetTimeout(),
		ProjectId:     in.GetProjectId(),
		EnvironmentId: in.GetEnvironmentId(),
		FolderId:      in.GetFolderId(),
		QueryType:     in.GetQueryType(),
		ResourceId:    in.GetResourceId(),
		Body:          body,
		Variables:     variables,
		CommitId:      in.GetCommitId(),
		VersionId:     in.GetVersionId(),
		CreatedAt:     time.Now().Format(time.RFC3339),
		UpdatedAt:     time.Now().Format(time.RFC3339),
	}

	_, err = q.collections["query"].InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	return &pb.Query{
		Id:            id,
		Title:         in.GetTitle(),
		Description:   in.GetDescription(),
		Timeout:       in.GetTimeout(),
		ProjectId:     in.GetProjectId(),
		EnvironmentId: in.GetEnvironmentId(),
		FolderId:      in.GetFolderId(),
		QueryType:     in.GetQueryType(),
		ResourceId:    in.GetResourceId(),
		Body:          in.GetBody(),
		Variables:     in.GetVariables(),
		CommitId:      in.GetCommitId(),
		VersionId:     in.GetVersionId(),
	}, nil
}

func (q *QueryRepo) Update(ctx context.Context, in *pb.UpdateQueryReq) (*pb.Query, error) {
	var (
		variables []*models.VariableSchema
	)

	body, err := helper.ConvertStructToResponse(in.GetBody())
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	variableMarshal, err := json.Marshal(in.GetVariables())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(variableMarshal, &variables)
	if err != nil {
		return nil, err
	}

	doc := models.QuerySchema{
		Id:            in.GetId(),
		Title:         in.GetTitle(),
		Description:   in.GetDescription(),
		Timeout:       in.GetTimeout(),
		ProjectId:     in.GetProjectId(),
		EnvironmentId: in.GetEnvironmentId(),
		FolderId:      in.GetFolderId(),
		QueryType:     in.GetQueryType(),
		ResourceId:    in.GetResourceId(),
		Body:          body,
		Variables:     variables,
		CommitId:      in.GetCommitId(),
		VersionId:     in.GetVersionId(),
		CreatedAt:     time.Now().Format(time.RFC3339),
		UpdatedAt:     time.Now().Format(time.RFC3339),
	}

	_, err = q.collections["query"].InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	return &pb.Query{
		Id:            in.GetId(),
		Title:         in.GetTitle(),
		Description:   in.GetDescription(),
		Timeout:       in.GetTimeout(),
		ProjectId:     in.GetProjectId(),
		EnvironmentId: in.GetEnvironmentId(),
		FolderId:      in.GetFolderId(),
		QueryType:     in.GetQueryType(),
		ResourceId:    in.GetResourceId(),
		Body:          in.GetBody(),
		Variables:     in.GetVariables(),
		CommitId:      in.GetCommitId(),
		VersionId:     in.GetVersionId(),
	}, nil
}

func (q *QueryRepo) GetSingle(ctx context.Context, in *pb.GetSingleQueryReq) (*pb.Query, error) {
	log.Printf("-->STRG:Get---> %+v", in)
	var (
		resSchema models.QuerySchema
		res       pb.Query
	)

	filter := bson.M{
		"id": in.GetId(),
		"is_deleted": bson.M{
			"$ne": true,
		},
	}

	opts := options.FindOne()
	opts.SetSort(
		bson.M{"created_at": -1},
	)

	if in.GetCommitId() != "" {
		filter["commit_id"] = in.GetCommitId()
	}

	singleResult := q.collections["query"].FindOne(
		ctx,
		filter,
		opts,
	)

	err := singleResult.Decode(&resSchema)
	if err != nil {
		return nil, err
	}

	err = helper.ConvertStructToStruct(resSchema, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (q *QueryRepo) GetList(ctx context.Context, in *pb.GetListQueryReq) (*pb.GetListQueryRes, error) {
	log.Printf("-->STRG:GetList---> %+v", in)
	var (
		queriesPb []*pb.Query
		queries   []models.QuerySchema
	)
	fmt.Println("TEST::::::::1")
	/*
		1. match: with version_id
		2. sort: with created_at descending order
		3. group: with id and select first element each group
	*/

	filterMatchStep1 := bson.M{
		//"project_id": in.GetProjectId(),
		"is_deleted": bson.M{
			"$ne": true,
		},
	}

	filterMatchStep2 := bson.M{
		"folder_id": in.GetFolderId(),
		// @TODO add search
	}
	fmt.Println("TEST::::::::2")
	pipeline := []bson.M{
		{
			"$match": filterMatchStep1,
		},
		{
			"$sort": bson.M{
				"created_at": -1,
			},
		},
		{
			"$group": bson.M{
				"_id": "$id",
				"id": bson.M{
					"$first": "$id",
				},
				"commit_id": bson.M{
					"$first": "$commit_id",
				},
				"title": bson.M{
					"$first": "$title",
				},
				"project_id": bson.M{
					"$first": "$project_id",
				},
				"folder_id": bson.M{
					"$first": "$folder_id",
				},
				"query_type": bson.M{
					"$first": "$query_type",
				},
				"description": bson.M{
					"$first": "$description",
				},
				"resource_id": bson.M{
					"$first": "$resource_id",
				},
			},
		},
		{
			"$match": filterMatchStep2,
		},
	}

	rows, err := q.collections["query"].Aggregate(
		ctx,
		pipeline,
		options.Aggregate(),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close(ctx)

	err = rows.All(ctx, &queries)
	if err != nil {
		return nil, err
	}

	err = helper.ConvertStructToStruct(queries, &queriesPb)
	if err != nil {
		return nil, err
	}

	return &pb.GetListQueryRes{
		Count:   int32(len(queriesPb)),
		Queries: queriesPb,
	}, nil
}

func (q *QueryRepo) Delete(ctx context.Context, in *pb.DeleteQueryReq) (*emptypb.Empty, error) {

	_, err := q.collections["query"].UpdateMany(ctx, bson.M{
		"id": in.GetId(),
	},
		bson.M{
			"$set": bson.M{
				"is_deleted": true,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (q *QueryRepo) GetQueryHistory(ctx context.Context, in *pb.GetQueryHistoryReq) (resp *pb.GetQueryHistoryRes, err error) {
	log.Printf("---STRG:GetApiReferenceChanges---> %+v", in)

	var filter []struct {
		CommitID   string   `bson:"_id"`
		Id         string   `bson:"id"`
		VersionIds []string `bson:"version_ids"`
		UpdatedAt  string   `bson:"updated_at"`
		CreatedAt  string   `bson:"created_at"`
	}

	aggResp, err := q.collections["query"].Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{
				"id": in.GetId(),
				//"project_id": in.GetProjectId(),
			},
		},
		{
			"$group": bson.M{
				"_id":        "$commit_id",
				"id":         bson.M{"$first": "$id"},
				"updated_at": bson.M{"$first": "$updated_at"},
				"created_at": bson.M{"$first": "$created_at"},
			},
		},
		{
			"$sort": bson.M{
				"created_at": -1,
			},
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "error getting query history")
	}
	defer aggResp.Close(ctx)

	err = aggResp.All(ctx, &filter)
	if err != nil {
		return nil, errors.Wrap(err, "aggregations.All error")
	}

	// Set count
	count := int64(len(filter))

	var queries = make([]*pb.GetQueryHistoryRes_QueryHistory, 0, count)
	for _, value := range filter {

		var query pb.GetQueryHistoryRes_QueryHistory
		// var mongoRes mongo_models.ApiReferenceChangesMongo

		commitInfo := models.CommitSchema{}

		err := q.collections["commit"].FindOne(ctx, bson.M{
			"id": value.CommitID,
		},
			&options.FindOneOptions{
				Sort: bson.M{
					"created_at": -1,
				},
			},
		).Decode(&commitInfo)
		if err != nil {
			return nil, errors.Wrap(err, "error while decoding commit")
		}

		query.CommitInfo = &pb.CommitInfo{
			Name:       commitInfo.Name,
			CommitType: commitInfo.CommitType,
			AuthorId:   commitInfo.AuthorId,
			Id:         commitInfo.Id,
			CreatedAt:  commitInfo.CreatedAt,
			UpdatedAt:  commitInfo.UpdatedAt,
		}
		query.VersionInfos = map[string]*pb.VersionInfo{}

		for _, versionId := range commitInfo.VersionIds {
			query.VersionInfos[versionId] = &pb.VersionInfo{
				VersionId: versionId,
			}
		}

		query.CreatedAt = value.CreatedAt
		query.UpdatedAt = value.UpdatedAt

		queries = append(queries, &query)
	}
	log.Printf("---STRG:==>GetQueryHistory---> %+v", filter)

	return &pb.GetQueryHistoryRes{
		Count:   int32(count),
		Queries: queries,
		Id:      in.GetId(),
	}, nil
}

func (q *QueryRepo) Revert(ctx context.Context, req *pb.RevertQueryReq) (resp *pb.Query, err error) {
	log.Printf("---STRG:Revert---> %+v", resp)

	var (
		query    models.QuerySchema
		queryRes pb.Query
	)

	err = q.collections["query"].FindOne(ctx, bson.M{
		"id":        req.GetId(),
		"commit_id": req.GetOldCommitId(),
	}).Decode(&query)
	if err != nil {
		return nil, errors.Wrap(err, "error while finding query")
	}

	query.CommitId = req.GetNewCommitId()
	query.VersionId = req.GetVersionId()
	query.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	query.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	_, err = q.collections["query"].InsertOne(ctx, &query)
	if err != nil {
		return nil, errors.Wrap(err, "error while inserting query")
	}

	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonQuery, &queryRes)
	if err != nil {
		return nil, err
	}

	return &queryRes, nil
}

func (q *QueryRepo) CreateManyQuery(ctx context.Context, req *pb.ManyVersions) (resp *emptypb.Empty, err error) {
	log.Printf("-->STRG:CreateMany---> %+v", req)

	var query models.QuerySchema

	err = q.collections["query"].FindOne(ctx, bson.M{
		"id":        req.GetId(),
		"commit_id": req.GetOldCommitId(),
	}).Decode(&query)
	if err != nil {
		return nil, errors.Wrap(err, "error while find query")
	}

	var arr []interface{}

	for _, version := range req.GetVersionIds() {

		queryInit := query
		queryInit.CommitId = req.NewCommitId
		queryInit.VersionId = version
		queryInit.CreatedAt = time.Now().UTC().Format(time.RFC3339)
		queryInit.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

		arr = append(arr, &query)
	}

	_, err = q.collections["query"].InsertMany(ctx, arr)
	if err != nil {
		return nil, errors.Wrap(err, "error while inserting queries")
	}

	return
}
