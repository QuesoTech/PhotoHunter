package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var DB *sql.DB

func getDbHandle() (db *sql.DB, err error) {
	// credentials and sensitive information are stred in the
	// environment so that we dont have to keep a config file that might
	// accidentally get committed...
	user := os.Getenv("$DBUSER")
	password := os.Getenv("$DBPASSWORD")
	host := os.Getenv("$DBHOST")
	dbname := os.Getenv("$DBNAME")

	// postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	conn_string := fmt.Sprintf("postgres://%s:%s@%s/%s", user, password, host, dbname)

	db, err = sql.Open("postgres", conn_string)
	return
}

func main() {
	DB, err := getDbHandle()
	check(err)
	defer DB.Close()
}
