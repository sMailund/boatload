package metService

import "fmt"

const (
	baseUrl              = "havvarsel-frost.met.no/api/v1/obs/glider"
	createTimeSeriesPath = "/ts/create"
	putTimeSeriesPath    = "/put" // counterintuitively defined as a POST endpoint in 3rd party API
	getTimeSeriesPath    = "/get"
)

func createTimeSeriesUrl() string {
	return fmt.Sprintf("%v%v", baseUrl, createTimeSeriesPath)
}

func putTimeSeriesUrl() string {
	return fmt.Sprintf("%v%v", baseUrl, putTimeSeriesPath)
}

func getTimeSeriesUrl() string {
	return fmt.Sprintf("%v%v", baseUrl, getTimeSeriesPath)
}
