package main

import (
	"fmt"
	"log"
	"net"

	"github.com/LavaJover/shvark-imageboard-service/internal/config"
	"github.com/LavaJover/shvark-imageboard-service/internal/delivery/grpcapi"
	"github.com/LavaJover/shvark-imageboard-service/internal/infrastructure/postgres"
	"github.com/LavaJover/shvark-imageboard-service/internal/usecase"
	imagepb "github.com/LavaJover/shvark-imageboard-service/proto/gen"
	"google.golang.org/grpc"
)

func main() {
	// Reading config
	cfg := config.MustLoad()
	// Init database
	db := postgres.MustInitDB(cfg)

	// Init image repo
	imageRepo := postgres.DefaultImageRepository{db}

	// Init user usecase
	uc := usecase.DefaultImageUsecase{&imageRepo}

	// Creating gRPC server
	grpcServer := grpc.NewServer()
	userHandler := grpcapi.ImageHandler{Uc: uc}

	imagepb.RegisterImageServiceServer(grpcServer, &userHandler)

	// Start
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("gRPC server started on %s:%s\n", cfg.Host, cfg.Port)
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v\n", err)
	}
}