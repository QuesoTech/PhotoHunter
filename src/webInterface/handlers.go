package main

import (
	"html/template"
	"net/http"
	"fmt"
         "golang.org/x/crypto/bcrypt"
	"database/sql"
	"log"
	
)

//Generates a template for the page and writes it. 

func genTemplate(w http.ResponseWriter, tmpt string, p *Page ){
	t, _ := template.ParseFiles("templates/"+ tmpt + ".html")
	t.Execute(w, p)
}


//Handles the front page of the site.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	p1 := &Page{Title: "index", Body: []byte(""), User: "Aaron"}		
	
	genTemplate(w, "index", p1)
}

//Handles the account page of the site.
func accountHandler(w http.ResponseWriter, r *http.Request) {
	p1 := &Page{Title: "account", Body: []byte(""), User: "Aaron"}		
	genTemplate(w, "account", p1)
}

//Handles a single file, such as robots, favicon, sitemap. 
//Adapted from http://stackoverflow.com/a/14187941
func singleHandler(path string, fname string) {
	http.HandleFunc(path, func (w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, fname)
	})
}


//Gets the data submitted from sign up. Encrypts the password and
//creates a new user account.
func signupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

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

func signinHandler(w http.ResponseWriter, r*http.Request){
	r.ParseForm()

	uname := string(r.FormValue("username"))
	pword := []byte(r.FormValue("pword"))
	var hword []byte

	//Something's going on with this query. 
	err := DB.QueryRow("SELECT pword FROM researchers WHERE uname=?",uname).Scan(&hword)
	
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that name")
	case err !=nil:
		log.Fatal(err)
	default:
		err = bcrypt.CompareHashAndPassword(hword,pword)
		check(err)

		fmt.Printf("User found")
	}

	
}
