package mongo

import (
	"context"
	"log"
	"time"
	pb "ucode_go_query_service/genproto/query_service"
	"ucode_go_query_service/models"
	"ucode_go_query_service/storage/repo"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/protobuf/types/known/emptypb"
)

type FolderRepo struct {
	collection *mongo.Collection
}

func NewFolderRepo(db *mongo.Database) repo.FolderRepoI {
	return &FolderRepo{
		collection: db.Collection("queries.folder"),
	}
}

func (f *FolderRepo) Create(ctx context.Context, in *pb.CreateFolderReq) (*pb.Folder, error) {
	log.Printf("-->STRG->Create folder---> %+v", in)

	id := uuid.New().String()

	doc := models.FolderSchema{
		Id:        id,
		Title:     in.GetTitle(),
		ProjectId: in.GetProjectId(),
		ParentId:  in.GetParentId(),
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	_, err := f.collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	return &pb.Folder{
		Id:        doc.Id,
		Title:     doc.Title,
		ProjectId: doc.ProjectId,
		ParentId:  doc.ParentId,
	}, nil
}

func (f *FolderRepo) Update(ctx context.Context, in *pb.UpdateFolderReq) (*pb.Folder, error) {
	log.Printf("-->STRG->Update folder---> %+v", in)

	_, err := f.collection.UpdateOne(
		ctx,
		bson.M{
			"id":         in.GetId(),
			"is_deleted": false,
		},
		bson.M{
			"$set": bson.M{
				"title":      in.GetTitle(),
				"updated_at": time.Now().UTC().Format(time.RFC3339),
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &pb.Folder{
		Id:        in.GetId(),
		Title:     in.GetTitle(),
		ProjectId: in.GetProjectId(),
		ParentId:  in.GetParentId(),
	}, nil
}

func (f *FolderRepo) GetSingle(ctx context.Context, in *pb.GetSingleFolderReq) (*pb.GetSingleFolderRes, error) {
	log.Printf("-->STRG->GetSingle folder---> %+v", in)

	// Create a FindOneOptions struct to set the sort and projection fields
	//opts := options.FindOne()
	//opts.SetSort(bson.M{
	//	"updated_at": -1,
	//})
	//
	//// Find the category
	//res := f.collection.FindOne(ctx, bson.M{
	//	"id": in.GetId(),
	//	// "version_id": in.VersionId,
	//}, opts)
	//if res.Err() != nil {
	//	return nil, res.Err()
	//}
	//
	//resDecode := models.FolderSchema{}
	//
	//// Decode the result into a Category struct
	//err := res.Decode(&resDecode)
	//if err != nil {
	//	return nil, errors.Wrap(err, "error decoding category")
	//}
	// @TODO not implemented yet
	return &pb.GetSingleFolderRes{}, nil
}

func (f *FolderRepo) GetList(ctx context.Context, in *pb.GetListFolderReq) (*pb.GetListFolderRes, error) {
	log.Printf("-->STRG->GetList folder---> %+v", in)
	result := pb.GetListFolderRes{}

	// 1. match version_id
	// 2. sort by created_at desc
	// 3. group by guid and get the first one (this is the latest version of the category)

	pipeline := []bson.M{
		{
			"$match": bson.M{
				// "version_id": in.GetVersionId(),
				//"project_id": in.GetProjectId(),
				"is_deleted": false,
			},
		},
		{
			"$sort": bson.M{
				"updated_at": -1,
			},
		},
		// {
		// 	"$group": bson.M{
		// 		"_id": "$guid",
		// 		"commit_id": bson.M{
		// 			"$first": "$commit_id",
		// 		},
		// 	},
		// },
	}

	rows, err := f.collection.Aggregate(
		ctx,
		pipeline,
		options.Aggregate(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error getting aggregate folder list")
	}
	defer rows.Close(ctx)

	filter := []struct {
		Id string `bson:"_id"`
		// CommitId string `bson:"commit_id"`
	}{}

	err = rows.All(ctx, &filter)
	if err != nil {
		return nil, errors.Wrap(err, "error getting folder list")
	}

	// Set total count
	count := int64(len(filter))

	bsonFilter := []bson.D{}
	for _, value := range filter {
		bsonData := bson.D{
			bson.E{Key: "id", Value: value.Id},
			// bson.E{Key: "commit_id", Value: value.CommitId},
		}
		bsonFilter = append(bsonFilter, bsonData)
	}

	// Query

	// var query primitive.M

	if len(bsonFilter) == 0 {
		return &pb.GetListFolderRes{
			Count: int32(count),
		}, nil
	}

	// query = bson.M{
	// 	"$or": bsonFilter,
	// }

	opts := options.Find()
	opts.SetSort(bson.M{
		"updated_at": -1,
	})
	//opts.SetSkip(in.GetOffset())
	//opts.SetLimit(in.GetLimit())

	rows, err = f.collection.Find(
		ctx,
		bson.M{
			//"project_id": in.GetProjectId(),
			"is_deleted": false,
		},
		opts,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while finding category list")
	}
	defer rows.Close(context.Background())

	for rows.Next(ctx) {
		res := models.FolderSchema{}
		err := rows.Decode(&res)
		if err != nil {
			return nil, errors.Wrap(err, "error decoding folder")
		}

		folder := pb.Folder{
			Id:        res.Id,
			Title:     res.Title,
			ProjectId: res.ProjectId,
			ParentId:  res.ParentId,
			// CommitId:   res.CommitId,
			// VersionId:  res.VersionId,
		}

		result.Folders = append(result.Folders, &folder)
	}
	result.Count = int32(count)

	return &result, nil
}

func (f *FolderRepo) Delete(ctx context.Context, in *pb.DeleteFolderReq) (*emptypb.Empty, error) {
	log.Printf("-->STRG->Delete folder---> %+v", in)

	_, err := f.collection.UpdateOne(ctx, bson.M{
		"id": in.GetId(),
	}, bson.M{
		"$set": bson.M{
			"is_deleted": true,
			"updated_at": time.Now().Format(time.RFC3339),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "error deleting folder")
	}

	return &emptypb.Empty{}, nil
}
