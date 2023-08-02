package service

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
	"ucode_go_query_service/config"
	pb "ucode_go_query_service/genproto/query_service"
	"ucode_go_query_service/grpc/client"
	"ucode_go_query_service/models"
	"ucode_go_query_service/pkg/helper"
	l "ucode_go_query_service/pkg/logger"
	"ucode_go_query_service/storage"

	"ucode_go_query_service/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// QueryService ...
type QueryService struct {
	cfg      config.Config
	strg     storage.StorageI
	logger   l.Logger
	services client.ServiceManagerI
	pb.UnimplementedQueryServiceServer
}

// NewQueryService ...
func NewQueryService(cfg config.Config, log l.Logger, services client.ServiceManagerI, strg storage.StorageI) *QueryService {
	return &QueryService{
		cfg:      cfg,
		strg:     strg,
		logger:   log,
		services: services,
	}
}

func (s *QueryService) CreateQuery(ctx context.Context, in *pb.CreateQueryReq) (*pb.Query, error) {
	s.logger.Info("--Create Query -- requested", logger.Any("req", in))

	strg, err := s.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		s.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resCommitID, err := strg.Commit().Create(ctx, &models.CommitSchema{
		ProjectId:  in.GetProjectId(),
		VersionIds: []string{in.GetVersionId()}, //this is current version id
		CommitType: in.GetCommitInfo().GetCommitType(),
		AuthorId:   in.GetCommitInfo().GetAuthorId(),
		Name:       in.GetCommitInfo().GetName(),
	})

	// Set Commit ID For Data
	in.CommitId = resCommitID.Id
	response, err := strg.Query().Create(ctx, in)
	if err != nil {
		s.logger.Error("Error while creating Query in service", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (s *QueryService) UpdateQuery(ctx context.Context, in *pb.UpdateQueryReq) (*pb.Query, error) {
	s.logger.Info("--Update Query-- requested", logger.Any("req", in))

	strg, err := s.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		s.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resCommitID, err := strg.Commit().Create(ctx, &models.CommitSchema{
		ProjectId:  in.GetProjectId(),
		VersionIds: []string{in.GetVersionId()},
		CommitType: in.GetCommitInfo().GetCommitType(),
		AuthorId:   in.GetCommitInfo().GetAuthorId(),
		Name:       in.GetCommitInfo().GetName(),
	})

	// Set Commit ID For Data
	in.CommitId = resCommitID.Id
	response, err := strg.Query().Update(ctx, in)
	if err != nil {
		s.logger.Error("Error while updating Query in service", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (s *QueryService) GetSingleQuery(ctx context.Context, in *pb.GetSingleQueryReq) (*pb.Query, error) {
	s.logger.Info("--Get Single Query-- requested", logger.Any("req", in))

	strg, err := s.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		s.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetCommitId() == "" && in.GetVersionId() != "" {
		commitInfo, err := strg.Commit().GetCommitByVersionId(ctx, in.GetVersionId())
		if err != nil {
			s.logger.Error("Error while getting single query in service", l.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		in.CommitId = commitInfo.GetId()
	}

	response, err := strg.Query().GetSingle(ctx, in)
	if err != nil {
		s.logger.Error("Error while getting single query in service", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	fmt.Println(response.GetCommitId())
	commitInfo, err := strg.Commit().GetByID(ctx, &pb.CommitPrimaryKey{
		Id: response.GetCommitId(),
	})
	if err != nil {
		s.logger.Error("Error while getting single query in service", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = helper.ConvertStructToStruct(commitInfo, &response.CommitInfo)
	if err != nil {
		s.logger.Error("Error while getting single query in service [COMMIT]", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (s *QueryService) GetListQuery(ctx context.Context, in *pb.GetListQueryReq) (*pb.GetListQueryRes, error) {
	s.logger.Info("--Get All Query-- requested", logger.Any("req", in))

	strg, err := s.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		s.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.ProjectId == "" {
		s.logger.Info("--Get All Query-- Project ID required", logger.Any("req", in))
		return nil, status.Error(codes.Internal, "Project id is required")
	}

	response, err := strg.Query().GetList(ctx, in)
	if err != nil {
		s.logger.Error("Error while getting list Query in service", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (s *QueryService) DeleteQuery(ctx context.Context, in *pb.DeleteQueryReq) (*emptypb.Empty, error) {
	s.logger.Info("--Delete Query-- requested", logger.Any("req", in))

	strg, err := s.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		s.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := strg.Query().Delete(ctx, in)
	if err != nil {
		s.logger.Error("Error while getting single Category in service", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (s *QueryService) RunQuery(ctx context.Context, in *pb.Query) (*pb.RunQueryRes, error) {
	var (
		res          pb.RunQueryRes
		variables    []models.VariableSchema
		bodyReplaced map[string]interface{}
	)
	/*
		body:
			body [data or select query]
			body_type
			action_type
			base_url
			cookies object
			headers object
			params object
			path
			resource
			object [key, value]
	*/
	s.logger.Info("--Run Query-- requested", logger.Any("req", in))

	strg, err := s.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		s.logger.Error("!!!RunQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	body := in.Body.AsMap()
	if len(body) == 0 {
		query, err := strg.Query().GetSingle(ctx, &pb.GetSingleQueryReq{
			Id: in.Id,
		})
		if err != nil {
			s.logger.Error("!!!RunQuery--->", logger.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		body = query.Body.AsMap()
	}
	fmt.Println("body", body)

	switch strings.ToUpper(in.GetQueryType()) {
	case "REST":
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		httpClient := &http.Client{Transport: tr}

		bodyJson, err := json.Marshal(body)
		if err != nil {
			s.logger.Error("Error while running query in service", l.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}

		variableJson, err := json.Marshal(in.GetVariables())
		if err != nil {
			s.logger.Error("Error while running query in service", l.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}

		err = json.Unmarshal(variableJson, &variables)
		if err != nil {
			s.logger.Error("Error while running query in service", l.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		fmt.Println("before replacing::::::", bodyJson, variables)
		bodyReplacedJson, err := helper.ReplaceVariable(bodyJson, variables)
		if err != nil {
			s.logger.Error("Error while running query in service", l.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}

		err = json.Unmarshal(bodyReplacedJson, &bodyReplaced)
		if err != nil {
			s.logger.Error("Error while running query in service", l.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		fmt.Println("bodyReplaced", bodyReplaced)
		req, err := helper.ParseQueryBody(bodyReplaced)
		if err != nil {
			s.logger.Error("Error while running query in service", l.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		fmt.Println("request::::", req)
		fmt.Println("request::::", req.Header)

		resRequest, err := httpClient.Do(req)
		if err != nil {
			s.logger.Error("Error while running query in service", l.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}

		defer resRequest.Body.Close()

		resJson, err := io.ReadAll(resRequest.Body)
		if err != nil {
			s.logger.Error("Error while running query in service", l.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		fmt.Println("res::::", string(resJson))
		res.Res = string(resJson)
		res.Status = resRequest.Status
		res.StatusCode = int32(resRequest.StatusCode)
	case "CLICKHOUSE":
		strg, err := s.strg.PoolCH().Get(in.GetResourceId())
		if err != nil {
			err = config.ErrProjectNotExists
			s.logger.Error("!!!RunQuery--->", logger.Error(err))
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		query, ok := body[helper.BODY]
		if !ok {
			s.logger.Error("!!!RunQuery--->", logger.Error(err))
			return nil, status.Error(codes.Internal, "query not found")
		}
		vars := map[string]interface{}{}
		// replace variables {{some_field}} to some_field =  :some_field
		for _, variable := range in.GetVariables() {
			var val string
			if variable.GetValue() != "" {
				val = variable.GetValue()
			} else {
				val = variable.GetDefault()
			}
			query = strings.ReplaceAll(query.(string), "{{"+variable.GetKey()+"}}", ":"+variable.GetKey())
			vars[variable.GetKey()] = val
		}

		resp, err := strg.Object().GetQuery(models.CommonInput{
			Query: query.(string),
			Data:  vars,
		})
		if err != nil {
			s.logger.Error("!!!RunQuery--->", logger.Error(err))
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		marshalledResponse, _ := json.Marshal(resp)
		fmt.Println("string ::", string(marshalledResponse))
		res.Res = string(marshalledResponse)
		res.Status = "SUCCESSFUL"
		res.StatusCode = http.StatusOK
	default:
		s.logger.Error("!!!wrong service type!!!")
	}

	err = strg.Log().Create(ctx, models.LogSchema{
		ProjectId:     in.GetProjectId(),
		EnvironmentId: in.GetEnvironmentId(),
		QueryId:       in.GetId(),
		UserId:        "", // @TODO
		Request:       body,
		Response:      res.GetRes(),
	})
	if err != nil {
		s.logger.Error("!!!log service not working!!!", logger.Error(err))
	}
	//fmt.Println("last", res)
	return &res, nil
}

func (s *QueryService) GetQueryHistory(ctx context.Context, in *pb.GetQueryHistoryReq) (resp *pb.GetQueryHistoryRes, err error) {
	s.logger.Info("--GetQueryHistory-- requested", logger.Any("req", in))

	strg, err := s.strg.PoolM().Get(in.GetResourceId())
	if err != nil {
		s.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetId()); err != nil {
		s.logger.Error("invalid guid", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	response, err := strg.Query().GetQueryHistory(ctx, in)
	if err != nil {
		s.logger.Error("Error while getting query history", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (s *QueryService) RevertQuery(ctx context.Context, req *pb.RevertQueryReq) (resp *pb.Query, err error) {
	s.logger.Info("--RevertQuery-- requested", logger.Any("req", req))

	strg, err := s.strg.PoolM().Get(req.GetResourceId())
	if err != nil {
		s.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resCommitID, err := strg.Commit().Create(ctx, &models.CommitSchema{
		ProjectId:  req.GetProjectId(),
		CommitType: req.GetCommitInfo().GetCommitType(),
		AuthorId:   req.GetCommitInfo().GetAuthorId(),
		Name:       req.GetCommitInfo().GetName(),
		VersionIds: []string{req.GetVersionId()},
	})

	req.NewCommitId = resCommitID.Id
	res, err := strg.Query().Revert(ctx, req)
	if err != nil {
		s.logger.Error("---ERR->RevertQuery->Create--->", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *QueryService) CreateManyQuery(ctx context.Context, req *pb.ManyVersions) (resp *emptypb.Empty, err error) {
	s.logger.Info("--CreateManyQuery-- requested", logger.Any("req", req))

	strg, err := s.strg.PoolM().Get(req.GetResourceId())
	if err != nil {
		s.logger.Error("!!!CreateQuery--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = strg.Commit().RemoveVersions(ctx, &pb.CommitWithRelease{
		ProjectId:     req.GetProjectId(),
		EnvironmentId: req.GetEnvironmentId(),
		ReleaseInfo: &pb.CommitWithRelease_Release{
			Ids: req.GetVersionIds(),
		},
		CommitType: req.GetCommitInfo().GetCommitType(),
		AuthorId:   req.GetCommitInfo().GetAuthorId(),
		Name:       req.GetCommitInfo().GetName(),
		Id:         req.GetOldCommitId(),
	})
	if err != nil {
		s.logger.Error("---ERR->CreateManyQuery--->", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = strg.Commit().UpdateVersions(ctx, &pb.CommitWithRelease{
		ProjectId:     req.GetProjectId(),
		EnvironmentId: req.GetEnvironmentId(),
		ReleaseInfo: &pb.CommitWithRelease_Release{
			Ids: req.GetVersionIds(),
		},
		CommitType: req.GetCommitInfo().GetCommitType(),
		AuthorId:   req.GetCommitInfo().GetAuthorId(),
		Name:       req.GetCommitInfo().GetName(),
		Id:         req.GetOldCommitId(),
	})
	if err != nil {
		s.logger.Error("---ERR->CreateManyQuery--->", l.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
