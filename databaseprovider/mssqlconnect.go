package databaseprovider

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var server = "192.168.1.21"
var port = 1433
var user = "smc"
var password = "smc@123456"
var database = "SellerTest2"

func MsSQLConnection() (db *sql.DB) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable", server, user, password, database)
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	return db
}
