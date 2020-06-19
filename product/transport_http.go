package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {

	r := chi.NewRouter()
	getProductByIDHandler := kithttp.NewServer(makeGetProductByIdEndPoint(s),
		getProductByIDRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIDHandler)

	getProductHandler := kithttp.NewServer(makeGetProductsEndPoints(s),
		getProductRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductHandler)
	return r
}

func getProductByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: productId,
	}, nil

}

func getProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	return request, nil

}
