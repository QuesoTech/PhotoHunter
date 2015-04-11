package main

type Datapoint struct {
	Id        int
	DatasetId int
	ImageURL  string
}

func GetDatapointById(id int) (dp *Datapoint, err error) {
	stmt, err := DB.Prepare("SELECT id, dataset_id, image_url FROM datapoint WHERE id = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var datasetID int
	var imageURL string

	err = row.Scan(&id)
	err = row.Scan(&datasetID)
	err = row.Scan(&imageURL)

	if err != nil {
		return
	}

	dp.Id = id
	dp.DatasetId = datasetID
	dp.ImageURL = imageURL

	return
}
