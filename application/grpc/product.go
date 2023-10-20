package grpc

import (
	"context"
	"github.com/lhps/desafio-01/application/grpc/pb"
	"github.com/lhps/desafio-01/application/usecase"
	"google.golang.org/grpc"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest, opts ...grpc.CallOption) (*pb.CreateProductResponse, error) {
	product, err := p.ProductUseCase.CreateProduct(in.Name, in.Description, float64(in.Price))
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          uint32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest, opts ...grpc.CallOption) (*pb.FindProductsResponse, error) {
	products, err := p.ProductUseCase.ListProducts()
	if err != nil {
		return nil, err
	}

	var productsResponse = make([]*pb.Product, 0)

	for _, product := range products {
		productsResponse = append(productsResponse, &pb.Product{
			Id:          uint32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		})
	}

	return &pb.FindProductsResponse{Products: productsResponse}, nil
}

func NexProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase:                    usecase,
		UnimplementedProductServiceServer: pb.UnimplementedProductServiceServer{},
	}
}
