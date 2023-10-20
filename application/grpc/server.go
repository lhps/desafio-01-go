package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/lhps/desafio-01/application/grpc/pb"
	"github.com/lhps/desafio-01/application/usecase"
	"github.com/lhps/desafio-01/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	productRepository := repository.ProductRepositoryDb{Db: database}
	productUseCase := usecase.ProductUseCase{ProductRepository: &productRepository}

	productGrpcService := NexProductGrpcService(productUseCase)
	pb.RegisterProductServiceServer(grpcServer, productGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start gRPC server", err)
	}

	log.Printf("gRPC server has been started gracefully on port %d", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot serve gRPC server")
	}
}
