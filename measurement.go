package airly

import (
	"fmt"
	"net/url"
	"time"
)

// Measurement is a response format that contains measurements from particular installation or area.
type Measurement struct {
	Current struct {
		FromDateTime time.Time `json:"fromDateTime"`
		TillDateTime time.Time `json:"tillDateTime"`
		Values       []struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
		} `json:"values"`
		Indexes []struct {
			Name        string  `json:"name"`
			Value       float64 `json:"value"`
			Level       string  `json:"level"`
			Description string  `json:"description"`
			Advice      string  `json:"advice"`
			Color       string  `json:"color"`
		} `json:"indexes"`
		Standards []struct {
			Name      string  `json:"name"`
			Pollutant string  `json:"pollutant"`
			Limit     float64 `json:"limit"`
			Percent   float64 `json:"percent"`
		} `json:"standards"`
	} `json:"current"`
	History []struct {
		FromDateTime time.Time `json:"fromDateTime"`
		TillDateTime time.Time `json:"tillDateTime"`
		Values       []struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
		} `json:"values"`
		Indexes []struct {
			Name        string  `json:"name"`
			Value       float64 `json:"value"`
			Level       string  `json:"level"`
			Description string  `json:"description"`
			Advice      string  `json:"advice"`
			Color       string  `json:"color"`
		} `json:"indexes"`
		Standards []struct {
			Name      string  `json:"name"`
			Pollutant string  `json:"pollutant"`
			Limit     float64 `json:"limit"`
			Percent   float64 `json:"percent"`
		} `json:"standards"`
	} `json:"history"`
	Forecast []struct {
		FromDateTime time.Time `json:"fromDateTime"`
		TillDateTime time.Time `json:"tillDateTime"`
		Values       []struct {
			Name  string  `json:"name"`
			Value float64 `json:"value"`
		} `json:"values"`
		Indexes []struct {
			Name        string  `json:"name"`
			Value       float64 `json:"value"`
			Level       string  `json:"level"`
			Description string  `json:"description"`
			Advice      string  `json:"advice"`
			Color       string  `json:"color"`
		} `json:"indexes"`
		Standards []struct {
			Name      string  `json:"name"`
			Pollutant string  `json:"pollutant"`
			Limit     float64 `json:"limit"`
			Percent   float64 `json:"percent"`
		} `json:"standards"`
	} `json:"forecast"`
}

// GetMeasurement returns measurements for concrete installation given by installationID.
func (c *Client) GetMeasurement(installationID int64) (*Measurement, error) {
	urlValues := url.Values{
		"installationId": []string{fmt.Sprintf("%d", installationID)},
	}
	var measurement Measurement

	err := c.get("measurements/installation", urlValues, &measurement)
	if err != nil {
		return nil, err
	}

	return &measurement, nil
}

// GetNearestMeasurement returns measurement for an installation closest to a given location.
func (c *Client) GetNearestMeasurement(lat, lng float64) (*Measurement, error) {
	urlValues := url.Values{
		"lat": []string{fmt.Sprintf("%f", lat)},
		"lng": []string{fmt.Sprintf("%f", lng)},
	}
	var measurement Measurement

	err := c.get("measurements/nearest", urlValues, &measurement)
	if err != nil {
		return nil, err
	}

	return &measurement, nil
}

// GetPointMeasurement returns measurements for any geographical location.
func (c *Client) GetPointMeasurement(lat, lng float64) (*Measurement, error) {
	urlValues := url.Values{
		"lat": []string{fmt.Sprintf("%f", lat)},
		"lng": []string{fmt.Sprintf("%f", lng)},
	}
	var measurement Measurement

	err := c.get("measurements/point", urlValues, &measurement)
	if err != nil {
		return nil, err
	}

	return &measurement, nil
}
