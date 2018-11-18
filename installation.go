package airly

import (
	"fmt"
	"net/url"
)

// Installation is an entity that binds together a sensor and its location where it's installed.
type Installation struct {
	ID       int `json:"id"`
	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	Address struct {
		Country         string `json:"country"`
		City            string `json:"city"`
		Street          string `json:"street"`
		Number          string `json:"number"`
		DisplayAddress1 string `json:"displayAddress1"`
		DisplayAddress2 string `json:"displayAddress2"`
	} `json:"address"`
	Elevation float64 `json:"elevation"`
	Airly     bool    `json:"airly"`
	Sponsor   struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Logo        string `json:"logo"`
		Link        string `json:"link"`
	} `json:"sponsor"`
}

// GetInstallation returns single installation metadata, given by installationID.
func (c *Client) GetInstallation(installationID int64) (*Installation, error) {
	var installation Installation

	err := c.get(fmt.Sprintf("installations/%d", installationID), nil, &installation)
	if err != nil {
		return nil, err
	}

	return &installation, nil
}

// GetNearestInstallation returns list of installations which are closest to a given point, sorted by distance to that point.
func (c *Client) GetNearestInstallation(lat, lng float64) (*[]Installation, error) {
	urlValues := url.Values{
		"lat": []string{fmt.Sprintf("%f", lat)},
		"lng": []string{fmt.Sprintf("%f", lng)},
	}
	var installations []Installation

	err := c.get("installations/nearest", urlValues, &installations)
	if err != nil {
		return nil, err
	}

	return &installations, nil
}
