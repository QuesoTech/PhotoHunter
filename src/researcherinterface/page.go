package main

// Page is used to construct a webpage using golangs template lib
type Page struct {
	Title string
	Body  []byte
	User  Researcher //Just using this to play around with pipelining
}
