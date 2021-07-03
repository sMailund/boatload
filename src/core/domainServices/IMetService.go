package domainServices

import (
	ent "domainEntities"
)

type MetService interface {
	submitData(ent.TimeSeries) error
}
