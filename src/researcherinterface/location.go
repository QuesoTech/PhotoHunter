package main

// GeoPoint represent the PostGIS GEOGRAPHY(POINT, 4326) found in the location table
type GeoPoint struct {
	Lat float64
	Lon float64
}

// Location represents a location restriction for a dataset
type Location struct {
	Id        int64
	DatasetId int64
	Target    GeoPoint
}
