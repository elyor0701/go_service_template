package mongo

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	pb "ucode_go_query_service/genproto/query_service"
	"ucode_go_query_service/models"
	"ucode_go_query_service/storage/repo"
)

type LogRepo struct {
	collection *mongo.Collection
}

func NewLogRepo(db *mongo.Database) repo.LogRepoI {
	return &LogRepo{
		collection: db.Collection("queries.log"),
	}
}

func (l *LogRepo) Create(ctx context.Context, in models.LogSchema) error {

	id := uuid.New().String()
	in.Id = id
	in.CreatedAt = time.Now().Format(time.RFC3339)

	_, err := l.collection.InsertOne(ctx, in)
	if err != nil {
		return err
	}
	return nil
}

func (l *LogRepo) GetSingle(ctx context.Context, in *pb.GetSingleLogReq) (*pb.Log, error) {
	var (
		logModel models.LogSchema
		log      pb.Log
	)

	err := l.collection.FindOne(ctx,
		bson.M{
			"id": in.GetId(),
			//"project_id":     in.GetProjectId(),
			//"environment_id": in.EnvironmentId,
		},
	).Decode(&logModel)
	if err != nil {
		return nil, err
	}

	jsonLog, err := json.Marshal(logModel)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonLog, &log)
	if err != nil {
		return nil, err
	}

	return &log, nil
}

func (l *LogRepo) GetList(ctx context.Context, in *pb.GetListLogReq) (*pb.GetListLogRes, error) {
	var (
		res pb.GetListLogRes
	)
	pipeline := []bson.M{
		{
			"$match": bson.M{
				// "version_id": in.GetVersionId(),
				//"project_id":     in.GetProjectId(),
				"query_id": in.GetQueryId(),
				//"environment_id": in.GetEnvironmentId(),
			},
		},
		{
			"$sort": bson.M{
				"created_at": -1,
			},
		},
	}

	rows, err := l.collection.Aggregate(
		ctx,
		pipeline,
		options.Aggregate(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error getting aggregate log list")
	}
	defer rows.Close(ctx)

	var filter []struct {
		Id string `bson:"_id"`
	}

	err = rows.All(ctx, &filter)
	if err != nil {
		return nil, errors.Wrap(err, "error getting log list")
	}

	// Set total count
	count := int64(len(filter))

	var bsonFilter []bson.D
	for _, value := range filter {
		bsonData := bson.D{
			bson.E{Key: "id", Value: value.Id},
			// bson.E{Key: "commit_id", Value: value.CommitId},
		}
		bsonFilter = append(bsonFilter, bsonData)
	}

	if len(bsonFilter) == 0 {
		return &pb.GetListLogRes{
			Count: count,
		}, nil
	}

	opts := options.Find()
	opts.SetSort(bson.M{
		"updated_at": -1,
	})
	opts.SetLimit(in.GetLimit())
	opts.SetSkip(in.GetOffset())

	rows, err = l.collection.Find(
		ctx,
		bson.M{
			"query_id": in.GetQueryId(),
			//"project_id":     in.GetProjectId(),
			//"environment_id": in.GetEnvironmentId(),
		},
		opts,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close(context.Background())

	for rows.Next(ctx) {
		var (
			row   models.LogSchema
			rowPb pb.Log
		)
		err := rows.Decode(&row)
		if err != nil {
			return nil, err
		}

		rowJson, err := json.Marshal(row)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(rowJson, &rowPb)

		res.Log = append(res.Log, &rowPb)
	}
	res.Count = count

	return &res, nil
}
