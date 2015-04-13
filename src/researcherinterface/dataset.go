package main

type Dataset struct {
	Id           int64
	ResearcherId int64
	Name         string
}

func NewDataset(researcherId int64, name string) (ds *Dataset, err error) {
	stmt, err := DB.Prepare("INSERT INTO dataset (researcher_id, name) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(researcherId, name).Scan(&id)
	if err != nil {
		return
	}

	ds = &Dataset{
		Id:           id,
		ResearcherId: researcherId,
		Name:         name,
	}

	return
}

func GetDatasetById(id int64) (ds *Dataset, err error) {
	var researcherId int64
	var name string

	err = DB.QueryRow("SELECT researcher_id, name FROM dataset WHERE id=$1", id).Scan(&researcherId, &name)

	ds = &Dataset{
		Id:           id,
		ResearcherId: researcherId,
		Name:         name,
	}

	return
}
