package airly

import (
	"time"
)

// MeasurementService is used to measurement operations.
// https://developer.airly.eu/docs#endpoints.measurements
type MeasurementService struct {
	client *Client
}

// https://developer.airly.eu/docs#concepts.indexes
type indexType string

const (
	// AirlyCAQI is an Airly quality index.
	// https://developer.airly.eu/docs#concepts.indexes.airlycaqi
	AirlyCAQI indexType = "AIRLY_CAQI"
	// CAQI is a European air quality index.
	CAQI indexType = "CAQI"
	// PIJP is a Polish air quality index.
	PIJP indexType = "PIJP"
)

// Value represents the name of the measurement (e.g., PM2.5)
// and measured value (e.g., concentration 60µg/m³).
type Value struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// Index represents an index value calculated for the measurements.
type Index struct {
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	Level       string  `json:"level"`
	Description string  `json:"description"`
	Advice      string  `json:"advice"`
	Color       string  `json:"color"`
}

// Standard represents a particular air quality standard.
type Standard struct {
	Name      string  `json:"name"`
	Pollutant string  `json:"pollutant"`
	Limit     float64 `json:"limit"`
	Percent   float64 `json:"percent"`
	Averaging string  `json:"averaging"`
}

// Current represents current measurement data.
type Current struct {
	FromDateTime time.Time  `json:"fromDateTime"`
	TillDateTime time.Time  `json:"tillDateTime"`
	Values       []Value    `json:"values"`
	Indexes      []Index    `json:"indexes"`
	Standards    []Standard `json:"standards"`
}

// History represents a historical measurement data.
type History struct {
	FromDateTime time.Time  `json:"fromDateTime"`
	TillDateTime time.Time  `json:"tillDateTime"`
	Values       []Value    `json:"values"`
	Indexes      []Index    `json:"indexes"`
	Standards    []Standard `json:"standards"`
}

// Forecast represents a measurement forecast.
type Forecast struct {
	FromDateTime time.Time  `json:"fromDateTime"`
	TillDateTime time.Time  `json:"tillDateTime"`
	Values       []Value    `json:"values"`
	Indexes      []Index    `json:"indexes"`
	Standards    []Standard `json:"standards"`
}

// Measurement is a response format that contains measurements
// from a particular installation or area.
type Measurement struct {
	Current  Current    `json:"current"`
	History  []History  `json:"history"`
	Forecast []Forecast `json:"forecast"`
}

type byIDMeasurementOpts struct {
	*urlQuery
}

// NewByIDMeasurementOpts is an opts builder for the installation id measurement query.
func NewByIDMeasurementOpts(id int64) *byIDMeasurementOpts {
	return &byIDMeasurementOpts{
		NewURLQuery().SetInstallationID(id),
	}
}

func (q *byIDMeasurementOpts) IndexType(index indexType) *byIDMeasurementOpts {
	q.SetIndexType(index)
	return q
}

// ByID returns measurements for concrete installation given by installationID.
func (c *MeasurementService) ByID(opts *byIDMeasurementOpts) (Measurement, error) {
	var measurement Measurement
	err := c.client.get("measurements/installation", opts.opts, &measurement)
	if err != nil {
		return Measurement{}, err
	}
	return measurement, nil
}

type nearestMeasurementOpts struct {
	*urlQuery
}

// NewNearestMeasurementOpts is an opts builder for the nearest measurement query.
func NewNearestMeasurementOpts(lat, lng float64) *nearestMeasurementOpts {
	return &nearestMeasurementOpts{
		NewURLQuery().SetLocation(lat, lng),
	}
}

func (q *nearestMeasurementOpts) MaxDistance(km float64) *nearestMeasurementOpts {
	q.SetMaxDistance(km)
	return q
}

func (q *nearestMeasurementOpts) IndexType(index indexType) *nearestMeasurementOpts {
	q.SetIndexType(index)
	return q
}

// Nearest returns measurement for an installation closest to a given location.
func (c *MeasurementService) Nearest(opts *nearestMeasurementOpts) (Measurement, error) {
	var measurement Measurement
	err := c.client.get("measurements/nearest", opts.opts, &measurement)
	if err != nil {
		return Measurement{}, err
	}
	return measurement, nil
}

type forPointMeasurementOpts struct {
	*urlQuery
}

// NewForPointMeasurementOpts is an opts builder for the point measurement query.
func NewForPointMeasurementOpts(lat, lng float64) *forPointMeasurementOpts {
	return &forPointMeasurementOpts{
		NewURLQuery().SetLocation(lat, lng),
	}
}

func (q *forPointMeasurementOpts) IndexType(index indexType) *forPointMeasurementOpts {
	q.SetIndexType(index)
	return q
}

// ForPoint returns measurements for any geographical location.
func (c *MeasurementService) ForPoint(opts *forPointMeasurementOpts) (Measurement, error) {
	var measurement Measurement
	err := c.client.get("measurements/point", opts.opts, &measurement)
	if err != nil {
		return Measurement{}, err
	}
	return measurement, nil
}
