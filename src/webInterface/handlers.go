package main

import (
	"html/template"
	"net/http"
	
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
