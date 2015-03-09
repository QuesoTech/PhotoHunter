//Researcher Interface for PhotoHunter Project

package main

import (
	
	"net/http"
	"database/sql"
	"fmt"
	"log"
	_"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	
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


func signupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)

	uname := r.FormValue("username")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	
	
	if r.FormValue("pword") != r.FormValue("pword2") {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	fmt.Println(uname, fname, lname, email)
	pword := r.FormValue("pword")


	hword, err := bcrypt.GenerateFromPassword([]byte(pword),10)
	check(err)



	stmt, err := DB.Prepare("INSERT INTO researchers (fname, lname, email, pword, uname) VALUES ($1,$2,$3,$4,$5)")
	check(err)
	defer stmt.Close()

	fmt.Println(hword)
	_, err = stmt.Exec(fname, lname, email, hword, uname)
	check(err)
	
	http.Redirect(w, r, "/", http.StatusFound)

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


	//Static file servers
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.ListenAndServe(":8080", nil)
}
