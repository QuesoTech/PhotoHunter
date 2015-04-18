package main

// GeoPoint represent the PostGIS GEOGRAPHY(POINT, 4326) found in the location table
type GeoPoint struct {
	Lat float64
	Lon float64
}

// Location represents a location restriction for a dataset
type Location struct {
	ID        int64
	DatasetID int64
	Target    GeoPoint
}

func NewLocation(datasetID int64, point GeoPoint) (l *Location, err error) {
	var id int64

	err = DB.QueryRow("INSERT INTO locations (ds_id, target) VALUES ($1, 'Point($2 $3)') RETURNING id", datasetID, point.Lat, point.Lon).Scan(&id)
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
