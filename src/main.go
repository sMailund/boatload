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
	mux := http.NewServeMux()

	uploadService := applicationServices.CreateUploadService(metServiceStub{})

	api.RegisterRoutes(*uploadService, mux)
	frontend.RegisterRoutes(frontend.InMemoryHtmlRetriever{}, mux)

	fmt.Printf("serving from %v...\n", port)
	err := http.ListenAndServe(port, mux)
	fmt.Fprintf(os.Stderr, "%v\n", err)
}
