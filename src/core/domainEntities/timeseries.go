package domainEntities

type TimeSeries struct {
	Tstype string
	Tseries []SeriesEntry
}

type SeriesEntry struct {
	Header SeriesHeader
	Observations []SeriesObservation
}

type SeriesHeader struct {
	Id HeaderId
	Extra HeaderExtra
}

type HeaderId struct {
	GliderId string
	Parameter string
}

type HeaderExtra struct {
	Source string
	Name string
}

type SeriesObservation struct {
	Time string // TODO: switch to date datatype??
	Body ObservationBody
}

type ObservationBody struct {
	Pos ObservationPosition
	Value string
	QcFlag string
}

type ObservationPosition struct {
	Lon string // TODO figure out correct datatype
	Lat string
	Depth string
	QcFlag string
}

