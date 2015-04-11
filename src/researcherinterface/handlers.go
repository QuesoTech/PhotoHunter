package main

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/sessions"
)


var store = sessions.NewCookieStore([]byte("aaron"))


//Generates a template for the page and writes it.
func genTemplate(w http.ResponseWriter, tmpt string, p *Page) {
	t, _ := template.ParseFiles("templates/" + tmpt + ".html")
	t.Execute(w, p)
}

//Handles the front page of the site.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ :=store.Get(r,"session-name")
	var email string;

	if name,ok := session.Values["user"].(string); ok {
		email = name
	}else {
		email = "Nobody"
	}

	p1 := &Page{Title: "index", Body: []byte(""), User: email}

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
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fname)
	})
}

//Gets the data submitted from sign up. Encrypts the password and
//creates a new user account.
func signupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()


	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")

	if r.FormValue("pword") != r.FormValue("pword2") {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	pword := r.FormValue("pword")
	hword, err := bcrypt.GenerateFromPassword([]byte(pword), 10)
	check(err)

	stmt, err := DB.Prepare("INSERT INTO researchers (fname, lname, email, pword) VALUES ($1,$2,$3,$4)")
	check(err)

	defer stmt.Close()


	_, err = stmt.Exec(fname, lname, email, hword)
	check(err)

	http.Redirect(w, r, "/", http.StatusFound)

}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := string(r.FormValue("email"))
	
	pword := []byte(r.FormValue("pword"))
	var hword []byte

	err := DB.QueryRow("SELECT pword FROM researchers WHERE email=$1", email).Scan(&hword)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that email")
	case err != nil:
		log.Fatal(err)
	default:
		err = bcrypt.CompareHashAndPassword(hword, pword)
		check(err)
		
		session, _:= store.Get(r,"session-name")
		session.Values["user"] = email
		session.Save(r,w)
		http.Redirect(w, r, "/", http.StatusFound)
	}

}


func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r,"session-name")
	delete(session.Values,"user")
	session.Save(r,w)

	http.Redirect(w,r,"/",http.StatusFound)


}
