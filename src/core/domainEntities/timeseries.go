package domainEntities

type TimeSeries struct {
	TimeSeriesType string        `json:"tstype"`
	TimeSeries     []seriesEntry `json:"tseries"`
}

type seriesEntry struct {
	Header       seriesHeader
	Observations []seriesObservation
}

type seriesHeader struct {
	Id    headerId
	Extra headerExtra
}

type headerId struct {
	GliderId  string
	Parameter string
}

type headerExtra struct {
	Source string
	Name   string
}

type seriesObservation struct {
	Time string // TODO: switch to date datatype??
	Body observationBody
}

type observationBody struct {
	Pos    observationPosition
	Value  string
	QcFlag string
}

type observationPosition struct {
	Lon    string // TODO figure out correct datatype
	Lat    string
	Depth  string
	QcFlag string
}
