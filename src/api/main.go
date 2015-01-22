package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var MDB_ADDR = "localhost"

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
