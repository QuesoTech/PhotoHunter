package main

// Datapoint represents one datapoint submitted to a dataset. It belongs to DatasetID
// and the associated image is located at ImageURL
type Datapoint struct {
	ID        int64
	DatasetID int64
	ImageURL  string
}

// NewDatapoint creates a new Datapoint, inserts it to the database and returns it
func NewDatapoint(datasetID int64, imageURL string) (dp *Datapoint, err error) {
	var id int64
	err = DB.QueryRow("INSERT INTO datapoint (dataset_id, image_url) VALUES ($1, $2) RETURNING id", datasetID, imageURL).Scan(&id)
	if err != nil {
		return
	}

	dp = &Datapoint{
		ID:        id,
		DatasetID: datasetID,
		ImageURL:  imageURL,
	}

	return
}

// GetDatapointByID fetches a datapoint from the database and returns it
func GetDatapointByID(id int64) (dp *Datapoint, err error) {
	var datasetID int64
	var imageURL string

	err = DB.QueryRow("SELECT id, dataset_id, image_url FROM datapoint WHERE id=$1", id).Scan(&id, &datasetID, &imageURL)

	if err != nil {
		return
	}

	dp = &Datapoint{
		ID:        id,
		DatasetID: datasetID,
		ImageURL:  imageURL,
	}

	return
}
