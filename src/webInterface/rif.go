//Researcher Interface for PhotoHunter Project

package main

import (
	
	"net/http"
	"database/sql"
	"fmt"
	"log"
	_"github.com/gorilla/sessions"
	
	_ "github.com/lib/pq"
)

var DB *sql.DB


func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getDbHandle() (db *sql.DB, err error) {
	user := "tester"
	password := "test"
	host:= "localhost"
	dbname:= "mydb"

	conn_string := fmt.Sprintf("postgres://%s:%s@%s/%s",user,password,host,dbname)
	db,err = sql.Open("postgres", conn_string)
	return
}


func main() {
	
	var err error
	//Database handlers
	DB,err = getDbHandle()
	check(err)
	defer DB.Close()
	check(DB.Ping())
	


	//Templates and specific pages
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/account",accountHandler)
	http.HandleFunc("/signup",signupHandler)
	http.HandleFunc("/signin",signinHandler)


	//Static file servers
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.ListenAndServe(":8080", nil)
}
