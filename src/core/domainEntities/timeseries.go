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
	Source string `json:"source"`
	Name   string `json:"name"`
}

type seriesObservation struct {
	Time string          `json:"time"` // TODO: switch to date datatype??
	Body observationBody `json:"body"`
}

type observationBody struct {
	Pos    observationPosition `json:"pos"`
	Value  string              `json:"value"`
	QcFlag string              `json:"qc_flag"`
}

type observationPosition struct {
	Lon    string `json:"lon"` // TODO figure out correct datatype
	Lat    string `json:"lat"`
	Depth  string `json:"depth"`
	QcFlag string `json:"qc_flag"`
}
