package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductRequest struct {
	Limit  int
	Offset int
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		Product, err := s.GetProductbyId(&req)
		if err != nil {
			panic(err)
		}
		return Product, nil
	}
	return getProductByIdEndPoint
}

func makeGetProductsEndPoints(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductRequest)
		result, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getProductsEndPoint
}
