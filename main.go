package main

import (
	database "GoLandRest/shared"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	databaseConnection := database.InitDB()

	defer databaseConnection.Close()

	fmt.Println(databaseConnection)
	//fmt.Println(connection)
	/*r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)*/

}
