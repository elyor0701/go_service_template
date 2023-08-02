package client

import (
	"ucode_go_query_service/config"
	"ucode_go_query_service/genproto/company_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	ResourceService() company_service.ResourceServiceClient
}

type grpcClients struct {
	resourceService company_service.ResourceServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connCompanyService, err := grpc.Dial(
		cfg.CompanyServiceHost+cfg.CompanyServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		resourceService: company_service.NewResourceServiceClient(connCompanyService),
	}, nil
}

func (g *grpcClients) ResourceService() company_service.ResourceServiceClient {
	return g.resourceService
}
