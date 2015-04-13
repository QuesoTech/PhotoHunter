package main

type GeoPoint struct {
	Lat float64
	Lon float64
}

type Location struct {
	Id        int64
	DatasetId int64
	Target    GeoPoint
}
