package main

import (
	"fmt"
	"foodways/database"
	"foodways/pkg/mysql"
	"foodways/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	mysql.DatabaseInit()
	database.RunMigration()
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}