package main

// Dataset represents a dataset requested by a researcher
type Dataset struct {
	ID int64

	// ID of owning researcher
	ResearcherID int64

	// name of dataset
	Name string
}

// NewDataset creates a new dataset and inserts it into the database
func NewDataset(researcherID int64, name string) (ds *Dataset, err error) {
	var id int64

	err = DB.QueryRow("INSERT INTO dataset (researcher_id, name) VALUES ($1, $2) RETURNING id", researcherID, name).Scan(&id)
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
