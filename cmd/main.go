package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
	"net"
	"ucode_go_query_service/config"
	"ucode_go_query_service/grpc"
	"ucode_go_query_service/grpc/client"
	"ucode_go_query_service/pkg/logger"
	"ucode_go_query_service/storage"
)

func main() {
	godotenv.Load(".env")
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, cfg.ServiceName)
	defer logger.Cleanup(log)

	//log.Info("main: mongoDbConfig",
	//	logger.String("host", cfg.MongoHost),
	//	logger.Int("port", cfg.MongoPort),
	//	logger.String("database", cfg.MongoDB))
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	//mongoStore, err := mongo.NewMongo(ctx, cfg)
	//if err != nil {
	//	log.Fatal("mongo.NewMongo", logger.Error(err))
	//	return
	//}
	//defer mongoStore.CloseDB()

	strg := storage.NewStorage()

	group, _ := errgroup.WithContext(context.Background())

	//grpcClients, err := client.NewGrpcClients(cfg)
	//if err != nil {
	//	log.Error("Error while connecting to grpc clients", logger.Error(err))
	//	return
	//}

	grpcClients, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Error("Error while connecting to grpc clients", logger.Error(err))
		return
	}

	go func() {
		err = storage.AutoConnect(grpcClients, strg)
		if err != nil {
			log.Error("Error while auto connecting to resources", logger.Error(err))
			return
		}
		log.Info("--AutoConnect-- completed successfully")
	}()

	//fmt.Println("strg.PoolM()::::::::", strg.PoolCH())

	grpcServer := grpc.SetUpServer(cfg, log, grpcClients, strg)

	group.Go(func() error {
		lis, err := net.Listen("tcp", cfg.RPCPort)
		if err != nil {
			log.Fatal("net.Listen", logger.Error(err))
		}

		log.Info("GRPC: Server being started...", logger.String("port", cfg.RPCPort))

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("grpcServer.Serve", logger.Error(err))
		}
		return err
	})

	err = group.Wait()
	if err != nil {
		log.Fatal("error on the server", logger.Error(err))
	}

	log.Info("server shutdown successfully")
}
