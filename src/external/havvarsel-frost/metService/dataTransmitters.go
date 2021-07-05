package metService

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type iDataTransmitter interface {
	SendData(r *http.Request) error
}

type devDataTransmitter struct {
}

func (d devDataTransmitter) SendData(r *http.Request) error {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	fmt.Printf("intercepted message:\n%v\n\n", string(requestDump))
	return nil
}
