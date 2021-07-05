package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sMailund/boatload/src/core/applicationServices"
	"github.com/sMailund/boatload/src/core/domainEntities"
	"io"
	"net/http"
	"strings"
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
	file, header, err := req.FormFile("file")
	if err != nil {
		return domainEntities.TimeSeries{}, err
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	// Copy the file data to my buffer
	io.Copy(&buf, file)
	// do something with the contents...
	// I normally have a struct defined and unmarshal into a struct, but this will
	// work as an example
	contents := buf.String()
	fmt.Println(contents)

	var timeSeries domainEntities.TimeSeries
	err = json.NewDecoder(&buf).Decode(&timeSeries)
	if err != nil {
		return domainEntities.TimeSeries{}, err
	}

	return timeSeries, nil
}
