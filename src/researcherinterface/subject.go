package main

// Subject represnts the subject being captured by a dataset. A dataset can have
// many of these
type Subject struct {
	Id        int
	DatasetId int
	Target    string
	Dummies   []string
}
