package apiTest

import (
	"bytes"
	"github.com/sMailund/boatload/src/core/applicationServices"
	"github.com/sMailund/boatload/src/core/domainEntities"
	"github.com/sMailund/boatload/src/external/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

type metServiceStub struct {
	submitMethod func() error
}

func (ms metServiceStub) SubmitData(_ domainEntities.TimeSeries) error {
	return ms.submitMethod()
}

func TestShouldMarshallPostBody(t *testing.T) {
	serviceStub := struct{ metServiceStub }{}
	serviceStub.submitMethod = func() error {
		return nil
	}

	body := []byte(testPayload)
	api.UploadService = *applicationServices.CreateUploadService(serviceStub)

	req := httptest.NewRequest(http.MethodPost, api.UploadRoute, bytes.NewReader(body))
	res := httptest.NewRecorder()

	api.UploadTimeSeries(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expexted 200, got %v: %v\n", res.Code, res.Body)
	}
}

func TestShouldOnlyAcceptPostMethod(t *testing.T) {

}

const testPayload = `{
  "tstype": "test",
  "tseries": [
    {
      "header": {
        "id": {
          "gliderID": "testID",
          "paramter": "testparam"
        },
        "extra": {
          "source": "sadfasdf",
          "name": "stuff"
        }
      },
      "observations": [
        {
          "time": "2020-06-16T06:00:00Z",
          "body": {
            "pos": {
              "lon": "1",
              "lat": "2",
              "depth": "3",
              "qc_flag": "test"
            },
            "value": "123",
            "qc_flag": "test"
          }
        }
      ]
    }
  ]
}
`
