package airly

import (
	"fmt"
	"net/url"
)

// InstallationService is used to installation operations.
// https://developer.airly.eu/docs#endpoints.installations
type InstallationService struct {
	client *Client
}

// Location represents the geographical coordinates of sensor installation.
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Address represents the detailed address of sensor installation.
type Address struct {
	Country         string `json:"country"`
	City            string `json:"city"`
	Street          string `json:"street"`
	Number          string `json:"number"`
	DisplayAddress1 string `json:"displayAddress1"`
	DisplayAddress2 string `json:"displayAddress2"`
}

// Sponsor represents the sponsor who bought the sensor.
type Sponsor struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Link        string `json:"link"`
}

// Installation is an entity that binds information about sensor installation.
type Installation struct {
	ID        int64    `json:"id"`
	Location  Location `json:"location"`
	Address   Address  `json:"address"`
	Elevation float64  `json:"elevation"`
	Airly     bool     `json:"airly"`
	Sponsor   Sponsor  `json:"sponsor"`
}

// ByID returns single installation metadata given by installationID.
// https://developer.airly.eu/docs#endpoints.installations.getbyid
func (s *InstallationService) ByID(id int64) (Installation, error) {
	var installation Installation
	u := fmt.Sprintf("installations/%d", id)
	err := s.client.get(u, nil, &installation)
	if err != nil {
		return Installation{}, err
	}
	return installation, nil
}

type nearestInstallationOpts struct {
	opts url.Values
}

// NewNearestInstallationOpts is an opts builder for the nearest installation query.
func NewNearestInstallationOpts(lat, lng float64) *nearestInstallationOpts {
	q := &nearestInstallationOpts{opts: map[string][]string{}}
	q.opts.Set("lat", fmt.Sprint(lat))
	q.opts.Set("lng", fmt.Sprint(lng))
	return q
}

func (q *nearestInstallationOpts) MaxDistance(km float64) *nearestInstallationOpts {
	q.opts.Set("maxDistanceKM", fmt.Sprint(km))
	return q
}

func (q *nearestInstallationOpts) MaxResults(limit float64) *nearestInstallationOpts {
	q.opts.Set("maxResults", fmt.Sprint(limit))
	return q
}

// Nearest returns list of installations closest to a given point,
// sorted by distance to that point.
// https://developer.airly.eu/docs#endpoints.installations.nearest
func (s *InstallationService) Nearest(opts *nearestInstallationOpts) ([]Installation, error) {
	var installations []Installation
	err := s.client.get("installations/nearest", opts.opts, &installations)
	if err != nil {
		return nil, err
	}
	return installations, nil
}
