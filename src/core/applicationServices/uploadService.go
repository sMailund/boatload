package applicationServices

import (
	"github.com/sMailund/boatload/src/core/domainEntities"
	"github.com/sMailund/boatload/src/core/domainServices"
)

type UploadService struct {
	metService domainServices.IMetService
}

func CreateUploadService(metService domainServices.IMetService) *UploadService {
	s := new(UploadService)
	s.metService = metService
	return s
}

func (us UploadService) UploadTimeSeries(series domainEntities.TimeSeries) error {
	err := us.metService.SubmitData(series)
	return err
}
