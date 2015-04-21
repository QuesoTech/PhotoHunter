package main

import (
	"fmt"
)

// GeoPoint represent the PostGIS GEOGRAPHY(POINT, 4326) found in the location table
type GeoPoint struct {
	Lat float64
	Lon float64
}

func (g GeoPoint) String() string {
	return fmt.Sprintf("ST_GeographyFromText('Point(%f %f)')", g.Lat, g.Lon)
}

// Location represents a location restriction for a dataset
type Location struct {
	ID        int64
	DatasetID int64
	Target    GeoPoint
}

func NewLocation(datasetID int64, point GeoPoint) (l *Location, err error) {
	var id int64

	err = DB.QueryRow("INSERT INTO locations (ds_id, target) VALUES ($1, $2) RETURNING id", datasetID, point).Scan(&id)
	if err != nil {
		return
	}

	l = &Location{
		ID:        id,
		DatasetID: datasetID,
		Target:    point,
	}

	return
}
