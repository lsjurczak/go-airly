package airly

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestInstallationService_ByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/installations/6600", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, mockInstallationResponse)
	})

	got, err := client.Installation.ByID(6600)
	if err != nil {
		t.Errorf("Installation.ByID: %v", err)
	}
	if !reflect.DeepEqual(got, mockInstallation) {
		t.Errorf("Installation.ByID returned %+v, want %+v", got, mockInstallation)
	}
}

func TestInstallationService_Nearest(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	opt := NewNearestInstallationOpts(50.062006, 19.940984)
	mux.HandleFunc("/installations/nearest", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, opt.opts)
		fmt.Fprint(w, mockListInstallationResponse)
	})

	got, err := client.Installation.Nearest(opt)
	if err != nil {
		t.Errorf("Installation.Nearest: %v", err)
	}
	if !reflect.DeepEqual(got, mockListInstallation) {
		t.Errorf("Installation.Nearest returned %+v, want %+v", got, mockListInstallation)
	}
}

var (
	mockInstallationResponse = `
	{
		"id": 9599,
		"location": {
			"latitude": 52.287217,
			"longitude": 21.108757
		},
		"address": {
			"country": "Poland",
			"city": "Ząbki",
			"street": "Piłsudskiego",
			"number": "35",
			"displayAddress1": "Ząbki",
			"displayAddress2": "Piłsudskiego"
		},
		"elevation": 85.02,
		"airly": true,
		"sponsor": {
			"id": 371,
			"name": "Powiat Wołomiński",
			"description": "Airly Sensor's sponsor",
			"logo": "https://cdn.airly.eu/logo/logo.jpg",
			"link": null,
			"displayName": null
		}
	}
`
	mockListInstallationResponse = fmt.Sprintf(`[%s]`, mockInstallationResponse)
)

var (
	mockInstallation = Installation{
		ID: 9599,
		Location: Location{
			Latitude:  52.287217,
			Longitude: 21.108757,
		},
		Address: Address{
			Country:         "Poland",
			City:            "Ząbki",
			Street:          "Piłsudskiego",
			Number:          "35",
			DisplayAddress1: "Ząbki",
			DisplayAddress2: "Piłsudskiego",
		},
		Elevation: 85.02,
		Airly:     true,
		Sponsor: Sponsor{
			ID:          371,
			Name:        "Powiat Wołomiński",
			Description: "Airly Sensor's sponsor",
			Logo:        "https://cdn.airly.eu/logo/logo.jpg",
		},
	}

	mockListInstallation = []Installation{mockInstallation}
)
