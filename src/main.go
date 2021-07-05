package main

import (
	"fmt"
	"github.com/sMailund/boatload/src/core/applicationServices"
	"github.com/sMailund/boatload/src/core/domainEntities"
	"github.com/sMailund/boatload/src/external/http/api"
	"github.com/sMailund/boatload/src/external/http/frontend"
	"net/http"
	"os"
)

const port = ":3000"

type metServiceStub struct{}

func (m metServiceStub) SubmitData(_ domainEntities.TimeSeries) error {
	fmt.Fprint(os.Stderr, "WARNING: attempting to communicate with unimplemented metservice")
	return nil
}

func main() {
	env := getDeploymentEnvironment()

	mux := http.NewServeMux()

	uploadService := applicationServices.CreateUploadService(metServiceStub{})

	api.RegisterRoutes(*uploadService, mux)
	frontend.RegisterRoutes(frontend.InMemoryHtmlRetriever{}, mux)

	fmt.Printf("running as %v, serving from %v...\n", env, port)
	err := http.ListenAndServe(port, mux)
	fmt.Fprintf(os.Stderr, "%v\n", err)
}

type environment int
const (
	DEV = iota
	PROD
)

// getDeploymentEnvironment gets the current deployment environment (i.e. dev or prod).
// Environement is configured through the BOATLOAD_ENV environment variable.
// Defaults to DEV.
func getDeploymentEnvironment() environment {
	switch env := os.Getenv("BOATLOAD_ENV"); env {
	case "PROD":
		return PROD
	case "DEV":
		return DEV
	default:
		fmt.Fprintf(os.Stderr, "WARNING: unrecognized BOATLOAD_ENV: '%v', defaulting to DEV\n", env) // TODO use logger
		return DEV
	}
}