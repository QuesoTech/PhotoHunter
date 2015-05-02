package main

import (
	"encoding/json"
	"net/http"
)

// getDatasetsHandler handles requests from photohunter for lists of available datasets
func getDatasetsHandler(w http.ResponseWriter, r *http.Request) {
	datasets, _ := GetAllDatasets()
	out := json.NewEncoder(w)
	out.Encode(datasets)
}

// submitDatapointHandler handles requests from photohunter for submitting data to a dataset
func submitDataHandler(w http.ResponseWriter, r *http.Request) {}
