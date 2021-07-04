package domainEntities

type TimeSeries struct {
	TimeSeriesType string        `json:"tstype"`
	TimeSeries     []seriesEntry `json:"tseries"`
}

type seriesEntry struct {
	Header       seriesHeader        `json:"header"`
	Observations []seriesObservation `json:"observations"`
}

type seriesHeader struct {
	Id    headerId    `json:"id"`
	Extra headerExtra `json:"extra"`
}

type headerId struct {
	GliderId  string `json:"gliderID"`
	Parameter string `json:"parameter"`
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