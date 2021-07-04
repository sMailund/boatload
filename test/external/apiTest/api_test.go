package apiTest

import (
	"bytes"
	"errors"
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
	submitMethod := func() error {
		return nil
	}
	api.UploadService = createUploadStub(submitMethod)

	res, req := createResponseAndRequestStubs(testPayload)

	api.UploadTimeSeries(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expexted 200, got %v: %v\n", res.Code, res.Body)
	}
}

func TestShouldOnlyAcceptPostMethod(t *testing.T) {
	methods := []string{
		http.MethodConnect,
		http.MethodDelete,
		http.MethodGet,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPut,
		http.MethodTrace,
	}

	submitMethod := func() error {
		return nil
	}
	api.UploadService = createUploadStub(submitMethod)


	for _, method := range methods {
		res, req := createResponseAndRequestStubsWithMethod(testPayload, method)

		api.UploadTimeSeries(res, req)
		if res.Code != http.StatusMethodNotAllowed {
			t.Errorf("expexted 405 for method %v, got %v: %v\n", method, res.Code, res.Body)
		}
	}

}

func TestShouldOnlyRespond500OnMetServiceError(t *testing.T) {
	submitMethod := func() error {
		return errors.New("sample error")
	}
	api.UploadService = createUploadStub(submitMethod)

	res, req := createResponseAndRequestStubs(testPayload)

	api.UploadTimeSeries(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Errorf("expexted 500, got %v: %v\n", res.Code, res.Body)
	}
}

func createUploadStub(submitMethod func() error) applicationServices.UploadService {
	serviceStub := struct{ metServiceStub }{}
	serviceStub.submitMethod = submitMethod

	return *applicationServices.CreateUploadService(serviceStub)
}

func createResponseAndRequestStubs(payload string) (*httptest.ResponseRecorder, *http.Request) {
	return createResponseAndRequestStubsWithMethod(payload, http.MethodPost)
}

func createResponseAndRequestStubsWithMethod(payload string, method string) (*httptest.ResponseRecorder, *http.Request) {
	body := []byte(payload)
	res := httptest.NewRecorder()
	req := httptest.NewRequest(method, api.UploadRoute, bytes.NewReader(body))
	return res, req
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
