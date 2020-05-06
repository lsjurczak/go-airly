package airly

import (
	"fmt"
	"net/url"
	"time"
)

// https://developer.airly.eu/docs#endpoints.measurements
type MeasurementService struct {
	client *Client
}

type indexType string

const (
	// https://developer.airly.eu/docs#concepts.indexes
	AirlyCAQI indexType = "AIRLY_CAQI"
	CAQI                = "CAQI"
	PIJP                = "PIJP"
)

// Value represents the name of the measurement (e.g., PM2.5)
// and measured value (e.g., concentration 60µg/m³)
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
	opts url.Values
}

func NewByIDMeasurementOpts(id int64) *byIDMeasurementOpts {
	q := &byIDMeasurementOpts{opts: map[string][]string{}}
	q.opts.Set("installationId", fmt.Sprint(id))
	return q
}

func (q *byIDMeasurementOpts) IndexType(index indexType) *byIDMeasurementOpts {
	q.opts.Set("indexType", string(index))
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
	opts url.Values
}

func NewNearestMeasurementOpts(lat, lng float64) *nearestMeasurementOpts {
	q := &nearestMeasurementOpts{opts: map[string][]string{}}
	q.opts.Set("lat", fmt.Sprint(lat))
	q.opts.Set("lng", fmt.Sprint(lng))
	return q
}

func (q *nearestMeasurementOpts) MaxDistance(km float64) *nearestMeasurementOpts {
	q.opts.Set("maxDistanceKM", fmt.Sprint(km))
	return q
}

func (q *nearestMeasurementOpts) IndexType(index indexType) *nearestMeasurementOpts {
	q.opts.Set("indexType", string(index))
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
	opts url.Values
}

func NewForPointMeasurementOpts(lat, lng float64) *forPointMeasurementOpts {
	q := &forPointMeasurementOpts{opts: map[string][]string{}}
	q.opts.Set("lat", fmt.Sprint(lat))
	q.opts.Set("lng", fmt.Sprint(lng))
	return q
}

func (q *forPointMeasurementOpts) IndexType(index indexType) *forPointMeasurementOpts {
	q.opts.Set("indexType", string(index))
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
