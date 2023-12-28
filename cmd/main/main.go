package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"gorm.io/driver/sqlserver"
	//_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/qadrina/go-bookstore-simple/pkg/routes"
)

func main() {
	// initialize router
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
