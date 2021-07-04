package domainServices

import (
	"github.com/sMailund/boatload/src/core/domainEntities"
)

type IMetService interface {
	SubmitData(domainEntities.TimeSeries) error
}
