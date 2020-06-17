package main

import (
	"net/http"

	database "github.com/GoLandRest/shared"
	"github.com/GoLandRest/shared/product"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	var productRepository = product.NewRepository(databaseConnection)
	var productService product.Service
	productService = product.NewService(productRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))

	http.ListenAndServe(":3000", r)

	//	fmt.Println(databaseConnection)
	//fmt.Println(connection)
	/*r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)*/

}
