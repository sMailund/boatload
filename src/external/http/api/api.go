package api

import (
	"encoding/json"
	"github.com/sMailund/boatload/src/core/applicationServices"
	"github.com/sMailund/boatload/src/core/domainEntities"
	"net/http"
)

var UploadService applicationServices.UploadService

const UploadRoute = "/api/upload"

func RegisterRoutes(us applicationServices.UploadService, mux *http.ServeMux) {
	UploadService = us
	mux.HandleFunc(UploadRoute, UploadTimeSeries)
}

func UploadTimeSeries(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}

	var timeSeries domainEntities.TimeSeries

	err := json.NewDecoder(req.Body).Decode(&timeSeries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = UploadService.UploadTimeSeries(timeSeries)
	if err != nil {
		// TODO: improve error handling
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
