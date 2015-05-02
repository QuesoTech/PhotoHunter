package main

// Dataset represents a dataset requested by a researcher
type Dataset struct {
	ID int64

	// ID of owning researcher
	ResearcherID int64

	// name of dataset
	Name string

	// number of images requested
	NumRequest int64
}

// NewDataset creates a new dataset and inserts it into the database
func NewDataset(researcherID int64, name string, numRequest int64) (ds *Dataset, err error) {
	var id int64

	err = DB.QueryRow("INSERT INTO dataset (researcher_id, name,numRequest) VALUES ($1, $2, $3) RETURNING id", researcherID, name,numRequest).Scan(&id)
	if err != nil {
		return
	}

	ds = &Dataset{
		ID:           id,
		ResearcherID: researcherID,
		Name:         name,
		NumRequest: numRequest, 
	}

	return
}

// GetDatasetByID fetches a dataset from the database by its id
func GetDatasetByID(id int64) (ds *Dataset, err error) {
	var researcherID int64
	var name string

	err = DB.QueryRow("SELECT researcher_id, name FROM dataset WHERE id=$1", id).Scan(&researcherID, &name)
	if err != nil {
		return
	}

	ds = &Dataset{
		ID:           id,
		ResearcherID: researcherID,
		Name:         name,
	}

	return
}

// GetAllDatasets fetches a list of datasets from the database
func GetAllDatasets() (sets []*Dataset, err error) {
	rows, err := DB.Query("SELECT id, researcher_id, name FROM dataset")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var researcherID int64
		var name string

		err = rows.Scan(&id, &researcherID, &name)
		if err != nil {
			return
		}

		ds := &Dataset{
			ID:           id,
			ResearcherID: researcherID,
			Name:         name,
		}
		sets = append(sets, ds)
	}

	return
}
