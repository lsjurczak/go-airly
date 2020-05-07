package airly

import (
	"fmt"
)

// InstallationService is used to installation operations.
// https://developer.airly.eu/docs#endpoints.installations
type InstallationService struct {
	client *Client
}

// SetLocation represents the geographical coordinates of sensor installation.
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
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Logo        string  `json:"logo"`
	Link        *string `json:"link"`
	DisplayName *string `json:"displayName"`
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
	*urlQuery
}

// NewNearestInstallationOpts is an opts builder for the nearest installation query.
func NewNearestInstallationOpts(lat, lng float64) *nearestInstallationOpts {
	return &nearestInstallationOpts{
		NewURLQuery().SetLocation(lat, lng),
	}
}

func (q *nearestInstallationOpts) MaxDistance(km float64) *nearestInstallationOpts {
	q.SetMaxDistance(km)
	return q
}

func (q *nearestInstallationOpts) MaxResults(limit float64) *nearestInstallationOpts {
	q.SetMaxResults(limit)
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
