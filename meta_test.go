package airly

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMetaService_Indexes(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/meta/indexes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, mockIndexesResponse)
	})

	got, err := client.Meta.Indexes()
	if err != nil {
		t.Errorf("Meta.Indexes: %v", err)
	}
	if !reflect.DeepEqual(got, mockIndexes) {
		t.Errorf("Meta.Indexes returned %+v, want %+v", got, mockIndexes)
	}
}

func TestMetaService_Measurements(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/meta/measurements", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, mockMeasurementsResponse)
	})

	got, err := client.Meta.Measurements()
	if err != nil {
		t.Errorf("Meta.Measurements: %v", err)
	}
	if !reflect.DeepEqual(got, mockMeasurements) {
		t.Errorf("Meta.Measurements returned %+v, want %+v", got, mockMeasurements)
	}
}

var (
	mockIndexesResponse = `
	[
    	{
			"name":"AIRLY_CAQI",
			"levels":[
				{
					"minValue": 0,
					"maxValue": 25.0,
					"values": "0-25",
					"level": "VERY_LOW",
					"description": "Very Low",
					"color": "#6BC926"
				}
        	]
    	}
	]
`
	mockMeasurementsResponse = `
	[
    	{
        	"name": "PM1",
        	"label": "PM1",
        	"unit": "µg/m³"
    	}
	]
`
)

var (
	mockIndexes = []IndexType{
		{
			Name: "AIRLY_CAQI",
			Levels: []Level{
				{
					MinValue:    0,
					MaxValue:    25,
					Values:      "0-25",
					Level:       "VERY_LOW",
					Description: "Very Low",
					Color:       "#6BC926",
				},
			},
		},
	}

	mockMeasurements = []MeasurementType{
		{
			Name:  "PM1",
			Label: "PM1",
			Unit:  "µg/m³",
		},
	}
)
