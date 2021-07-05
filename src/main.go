package main

import (
	"fmt"
	"github.com/sMailund/boatload/src/core/applicationServices"
	"github.com/sMailund/boatload/src/core/domainServices"
	"github.com/sMailund/boatload/src/external/havvarsel-frost/metService"
	"github.com/sMailund/boatload/src/external/http/api"
	"github.com/sMailund/boatload/src/external/http/frontend"
	"net/http"
	"os"
)

const port = ":3000"

func main() {
	env := getDeploymentEnvironment()
	metServiceImpl := getMetService(env)

	uploadService := applicationServices.CreateUploadService(metServiceImpl)

	mux := http.NewServeMux()
	api.RegisterRoutes(*uploadService, mux)
	frontend.RegisterRoutes(frontend.InMemoryHtmlRetriever{}, mux)

	fmt.Printf("running as %v, serving from %v...\n", env, port)
	err := http.ListenAndServe(port, mux)
	fmt.Fprintf(os.Stderr, "%v\n", err)
}

func getMetService(env environment) domainServices.IMetService {
	switch env {
	case DEV:
		return metService.GetDevMetService()
	case PROD:
		panic("PROD METSERVICE NOT IMPLEMENTED")
	}
	return nil
}

type environment int

const (
	DEV environment = iota
	PROD
)

func (e environment) String() string {
	return [...]string{"dev", "prod"}[e]
}

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
