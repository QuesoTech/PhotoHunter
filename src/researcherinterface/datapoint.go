package main

type Datapoint struct {
	Id        int64
	DatasetId int64
	ImageURL  string
}

func NewDatapoint(datasetId int64, imageURL string) (dp *Datapoint, err error) {
	stmt, err := DB.Prepare("INSERT INTO datapoint (dataset_id, image_url) VALUES ($1, $2)")
	if err != nil {
		return
	}
	defer stmt.Close()

	var id int64
	err := stmt.QueryRow(datasetId, imageURL).Scan(&id)
	if err != nil {
		return
	}

	dp = &Datapoint{
		Id:        id,
		DatasetID: datasetId,
		ImageURL:  imageURL,
	}

	return
}

func GetDatapointById(id int64) (dp *Datapoint, err error) {
	stmt, err := DB.Prepare("SELECT id, dataset_id, image_url FROM datapoint WHERE id = $1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var datasetID int64
	var imageURL string

	err = row.Scan(&id)
	err = row.Scan(&datasetID)
	err = row.Scan(&imageURL)

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
