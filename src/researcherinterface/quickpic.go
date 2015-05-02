package main

import (
	"net/http"
)

// getDatapointHandler handles requests from quickpic for data to evaluate
func getDatapointHandler(w http.ResponseWriter, r *http.Request) {}

// evalDatapointHandler handles requests from quickpic to submit datapoint evaluation
// and return datapoint statistics in return
func evalDatapointHandler(w http.ResponseWriter, r *http.Request) {}
