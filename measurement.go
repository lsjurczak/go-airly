package airly

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Value represents name of the measurement (e.g., PM2.5)
// and measured value (e.g., concentration 60µg/m3)
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

// Standard represents particular air quality standard.
type Standard struct {
	Name      string  `json:"name"`
	Pollutant string  `json:"pollutant"`
	Limit     float64 `json:"limit"`
	Percent   float64 `json:"percent"`
}

// Current represents a current measurement data.
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

// Measurement is a response format that contains measurements from
// particular installation or area.
type Measurement struct {
	Current  Current    `json:"current"`
	History  []History  `json:"history"`
	Forecast []Forecast `json:"forecast"`
}

// GetMeasurement returns measurements for concrete installation,
// given by installationID.
func (c *Client) GetMeasurement(installationID int64) (*Measurement, error) {
	var measurement Measurement

	err := c.get(
		"measurements/installation",
		url.Values{
			"installationId": []string{strconv.Itoa(int(installationID))},
		},
		&measurement,
	)
	if err != nil {
		return nil, err
	}

	return &measurement, nil
}

// GetNearestMeasurement returns measurement for an installation closest
// to a given location.
func (c *Client) GetNearestMeasurement(lat, lng float64) (*Measurement, error) {
	var measurement Measurement

	err := c.get(
		"measurements/nearest",
		url.Values{
			"lat": []string{fmt.Sprintf("%f", lat)},
			"lng": []string{fmt.Sprintf("%f", lng)},
		},
		&measurement,
	)
	if err != nil {
		return nil, err
	}

	return &measurement, nil
}

// GetPointMeasurement returns measurements for any geographical location.
func (c *Client) GetPointMeasurement(lat, lng float64) (*Measurement, error) {
	var measurement Measurement

	err := c.get(
		"measurements/point",
		url.Values{
			"lat": []string{fmt.Sprintf("%f", lat)},
			"lng": []string{fmt.Sprintf("%f", lng)},
		},
		&measurement,
	)
	if err != nil {
		return nil, err
	}

	return &measurement, nil
}
