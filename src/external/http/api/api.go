package api

import (
	"bytes"
	"encoding/json"
	"github.com/sMailund/boatload/src/core/applicationServices"
	"github.com/sMailund/boatload/src/core/domainEntities"
	"io"
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

	timeSeries, err := readTimeSeriesFromRequest(req)
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

func readTimeSeriesFromRequest(req *http.Request) (domainEntities.TimeSeries, error) {

	req.ParseMultipartForm(32 << 20) // limit your max input length!
	var buf bytes.Buffer
	// in your case file would be fileupload
	file, _, err := req.FormFile("file")
	if err != nil {
		return domainEntities.TimeSeries{}, err
	}
	defer file.Close()
	// Copy the file data to my buffer
	io.Copy(&buf, file)

	var timeSeries domainEntities.TimeSeries
	err = json.NewDecoder(&buf).Decode(&timeSeries)
	if err != nil {
		return domainEntities.TimeSeries{}, err
	}

	return timeSeries, nil
}
