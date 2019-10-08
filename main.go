package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/ugizashinje/epoc/api"
	"github.com/ugizashinje/epoc/service"
)

var connStr string
var db *sql.DB

func main() {
	fmt.Println("Start")
	var err error
	connStr = "user=postgres dbname=postgres host=localhost sslmode=disable"
	service.Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database [%v]", err)
	}
	api.InitRoutes()
}
