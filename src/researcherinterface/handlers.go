package main

import (
	"database/sql"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"strconv"
	"time"
)

var store = sessions.NewCookieStore([]byte("aaron"))

//Generates a template for the page and writes it.
func genTemplate(w http.ResponseWriter, tmpt string, p *Page) {
	t, _ := template.ParseFiles("templates/" + tmpt + ".html")
	t.Execute(w, p)
}

//Handles the front page of the site.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	var email string

	if name, ok := session.Values["user"].(string); ok {
		email = name
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

//Creates a dataset
func createDatasetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	r.ParseForm()
	var id int64


	if rid, ok := session.Values["id"].(string); ok {
		id,_ = strconv.ParseInt(rid,10,64)
	}
	
	//create dataset
	
	name := r.FormValue("name")

	//create subjects
	target := r.FormValue("target")
	dummies := strings.Split(r.FormValue("dummies"),",")

	//create locations
	lat,err := strconv.ParseFloat(r.FormValue("lat"),64)
	long,err := strconv.ParseFloat(r.FormValue("long"),64)
	
	stime,err := time.Parse("15:04:05",r.FormValue("stime"))
	etime,err := time.Parse("15:04:05",r.FormValue("etime"))
	
	//time of day. HH:MM:SS
	
	ds, err := NewDataset(id, name)
	_,err = NewSubject(ds.ID,target,dummies)
	_,err = NewLocation(ds.ID,GeoPoint{Lat:lat,Lon:long})
	_,err = NewTimeperiod(ds.ID,stime,etime)

	if err != nil {
	http.Redirect(w,r,"/",http.StatusFound)
	}else {
		http.Redirect(w,r,"/",http.StatusNotFound)
	}
	
	//POINT(lat,long)
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

	_, err = NewResearcher(Name{fname, lname}, email, string(hword))
	check(err)

	http.Redirect(w, r, "/", http.StatusFound)
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := string(r.FormValue("email"))
	pword := []byte(r.FormValue("pword"))
	var hword []byte
	var id int64
	var fname, lname string
	err := DB.QueryRow("SELECT pword,id,fname,lname FROM researchers WHERE email=$1", email).Scan(&hword,&id,&fname,&lname)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that email")
	case err != nil:
		log.Fatal(err)
	default:
		err = bcrypt.CompareHashAndPassword(hword, pword)
		check(err)

		session, _ := store.Get(r, "session-name")

		

		session.Values["user"] = email
		session.Values["fname"] = fname
		session.Values["lname"] = lname
		session.Values["id"] = id
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	delete(session.Values, "user")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)

}
