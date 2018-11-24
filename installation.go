package airly

import (
	"fmt"
	"net/url"
)

// Location represents geographical coordinates where the sensor was installed.
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Address represents detailed address where the sensor was installed.
type Address struct {
	Country         string `json:"country"`
	City            string `json:"city"`
	Street          string `json:"street"`
	Number          string `json:"number"`
	DisplayAddress1 string `json:"displayAddress1"`
	DisplayAddress2 string `json:"displayAddress2"`
}

// Sponsor represents sponsor who bought the sensor.
type Sponsor struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Link        string `json:"link"`
}

// Installation is an entity that binds together information about
// sensor installation.
type Installation struct {
	ID        int      `json:"id"`
	Location  Location `json:"location"`
	Address   Address  `json:"address"`
	Elevation float64  `json:"elevation"`
	Airly     bool     `json:"airly"`
	Sponsor   Sponsor  `json:"sponsor"`
}

// GetInstallation returns single installation metadata,
// given by installationID.
func (c *Client) GetInstallation(installationID int64) (*Installation, error) {
	var installation Installation

	err := c.get(
		fmt.Sprintf("installations/%d", installationID),
		nil,
		&installation,
	)
	if err != nil {
		return nil, err
	}

	return &installation, nil
}

// GetNearestInstallation returns list of installations which are closest
// to a given point, sorted by distance to that point.
func (c *Client) GetNearestInstallation(lat, lng float64) (*[]Installation, error) {
	var installations []Installation

	err := c.get(
		"installations/nearest",
		url.Values{
			"lat": []string{fmt.Sprintf("%f", lat)},
			"lng": []string{fmt.Sprintf("%f", lng)},
		},
		&installations,
	)
	if err != nil {
		return nil, err
	}

	return &installations, nil
}
