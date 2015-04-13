//Researcher Interface for PhotoHunter Project

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

// global connection to the database. safe for concurrent connections
var DB *sql.DB

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getDbHandle() (db *sql.DB, err error) {
	// credentials and sensitive information are stred in the
	// environment so that we dont have to keep a config file that might
	// accidentally get committed...
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	host := os.Getenv("DBHOST")
	dbname := os.Getenv("DBNAME")

	connString := fmt.Sprintf("postgres://%s:%s@%s/%s", user, password, host, dbname)
	db, err = sql.Open("postgres", connString)
	return
}

func main() {

	var err error
	//Database handlers
	DB, err = getDbHandle()
	check(err)
	defer DB.Close()
	check(DB.Ping())

	//Templates and specific pages
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/account", accountHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/signin", signinHandler)
	http.HandleFunc("/logout", logoutHandler)

	//Static file servers
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.ListenAndServe(":8080", nil)
}
