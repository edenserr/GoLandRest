package database

import "database/sql"

func InitDB() *sql.DB {
	connection := "root:admin@tcp(localhost::3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connection)
	// defer: ejecuta funciones una vez ka función main() termina
	defer databaseConnection.Close()
	if err != nil {
		panic(err.Error())

	}
	return databaseConnection
}
