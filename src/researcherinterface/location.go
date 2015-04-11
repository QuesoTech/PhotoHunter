package main

type GeoPoint struct {
	Lat float64
	Lon float64
}

type Location struct {
	Id        int
	DatasetId int
	Target    GeoPoint
}
