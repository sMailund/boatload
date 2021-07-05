package metService

import (
	"bytes"
	"encoding/json"
	"github.com/sMailund/boatload/src/core/domainEntities"
	"github.com/sMailund/boatload/src/core/domainServices"
	"net/http"
)

type HavvarselMetService struct {
	authenticator   iAuthenticator
	dataTransmitter iDataTransmitter
}

func (h HavvarselMetService) sendToApi(series domainEntities.TimeSeries, url string) error {
	payload, err := json.Marshal(series)
	if err != nil { // TODO improve error handling
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	// authenticate message
	err = h.authenticator.authenticate(req)
	if err != nil {
		return err
	}

	//transmit message
	return h.dataTransmitter.SendData(req)
}

func (h HavvarselMetService) SubmitData(series domainEntities.TimeSeries) error {
	// TODO: this might be correct to only do once, not every request. further research needed
	// create time series
	err := h.sendToApi(series.HeadersOnly(), createTimeSeriesUrl())
	if err != nil {
		return err
	}

	// submit data to time series
	return h.sendToApi(series, putTimeSeriesUrl())
}

func GetDevMetService() domainServices.IMetService {
	return HavvarselMetService{
		authenticator:   devAuthenticator{},
		dataTransmitter: devDataTransmitter{},
	}
}
