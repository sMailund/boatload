package applicationServices

import (
	"github.com/sMailund/boatload/src/core/domainEntities"
	"github.com/sMailund/boatload/src/core/domainServices"
)

type uploadService struct {
	metService domainServices.IMetService
}

func (us uploadService) UploadTimeSeries(series domainEntities.TimeSeries) error {
	err := us.metService.SubmitData(series)
	return err
}
