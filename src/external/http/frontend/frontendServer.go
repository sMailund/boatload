package frontend

import (
	"net/http"
)

var frontendRetriever IFrontendRetriever

const route = "/uploader"

func RegisterRoutes(retriever IFrontendRetriever, mux *http.ServeMux) {
	frontendRetriever = retriever
	mux.HandleFunc(route, serveFrontend)
}

func serveFrontend(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}

	frontend, err := frontendRetriever.getFrontend()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(frontend)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
