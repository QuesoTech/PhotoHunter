package main

type Datapoint struct {
	Id        int64
	DatasetId int64
	ImageURL  string
}

func NewDatapoint(datasetId int64, imageURL string) (dp *Datapoint, err error) {
	var id int64
	err = DB.QueryRow("INSERT INTO datapoint (dataset_id, image_url) VALUES ($1, $2) RETURNING id", datasetId, imageURL).Scan(&id)
	if err != nil {
		return
	}

	dp = &Datapoint{
		Id:        id,
		DatasetId: datasetId,
		ImageURL:  imageURL,
	}

	return
}

func GetDatapointById(id int64) (dp *Datapoint, err error) {
	var datasetID int64
	var imageURL string

	err = DB.QueryRow("SELECT id, dataset_id, image_url FROM datapoint WHERE id=$1", id).Scan(&id, &datasetID, &imageURL)

	if err != nil {
		return
	}

	dp = &Datapoint{
		Id:        id,
		DatasetId: datasetID,
		ImageURL:  imageURL,
	}

	return
}
