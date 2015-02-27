package main

type Name struct {
	First string
	Last  string
}

type Researcher struct {
	Id int
	Name
	Email string
}
