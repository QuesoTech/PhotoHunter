package main

import (
	"encoding/json"
	"net/http"
)

// getDatapointHandler handles requests from quickpic for data to evaluate
func getDatapointHandler(w http.ResponseWriter, r *http.Request) {
	query := `SELECT dp.id, dp.dataset_id, dp.image_url, s.target, s.dummies
	FROM datapoints dp 
	JOIN datasets ds, subjects s ON (dp.dataset_id = ds.id, s.ds_id = ds.id)
	ORDER BY random() limit 1`

	var id, datasetID int64
	var url, target string
	var dummies []string
	DB.QueryRow(query).Scan(id, datasetID, url, target, dummies)

	out := json.NewEncoder(w)
	out.Encode(struct {
		ID        int64
		DatasetID int64
		ImageURL  string
		Target    string
		Dummies   []string
	}{
		id,
		datasetID,
		url,
		target,
		dummies,
	})
}

// evalDatapointHandler handles requests from quickpic to submit datapoint evaluation
// and return datapoint statistics in return
func evalDatapointHandler(w http.ResponseWriter, r *http.Request) {}
