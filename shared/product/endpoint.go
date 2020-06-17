package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
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
