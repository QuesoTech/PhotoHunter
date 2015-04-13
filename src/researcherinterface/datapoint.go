package main

// Datapoint represents one datapoint submitted to a dataset. It belongs to DatasetID
// and the associated image is located at ImageURL
type Datapoint struct {
	ID int64

	// ID of owning dataset
	DatasetID int64

	// URL of contributed image
	ImageURL string
}

// NewDatapoint creates a new Datapoint, inserts it to the database and returns it
func NewDatapoint(datasetID int64, imageURL string) (dp *Datapoint, err error) {
	// initialize id since we dont know what it is yet
	var id int64

	// reach out to the database
	err = DB.QueryRow("INSERT INTO datapoint (dataset_id, image_url) VALUES ($1, $2) RETURNING id", datasetID, imageURL).Scan(&id)

	// did something go wrong?
	if err != nil {
		return // something went wrong! lets get out of here!
	}

	//blank space for readability

	// create and populate datapoint
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
