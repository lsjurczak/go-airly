package airly

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestMeasurementService_ByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	opt := NewByIDMeasurementOpts(6600)
	mux.HandleFunc("/measurements/installation", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, opt.opts)
		fmt.Fprint(w, mockMeasurementResponse)
	})

	got, err := client.Measurement.ByID(opt)
	if err != nil {
		t.Errorf("Measurement.ByID: %v", err)
	}
	if !reflect.DeepEqual(got, mockMeasurement) {
		t.Errorf("Measurement.ByID returned %+v, want %+v", got, mockMeasurement)
	}
}

func TestMeasurementService_Nearest(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	opt := NewNearestMeasurementOpts(52.287217, 21.108757)
	mux.HandleFunc("/measurements/nearest", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, opt.opts)
		fmt.Fprint(w, mockMeasurementResponse)
	})

	got, err := client.Measurement.Nearest(opt)
	if err != nil {
		t.Errorf("Measurement.Nearest: %v", err)
	}
	if !reflect.DeepEqual(got, mockMeasurement) {
		t.Errorf("Measurement.Nearest returned %+v, want %+v", got, mockMeasurement)
	}
}

func TestMeasurementService_ForPoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	opt := NewForPointMeasurementOpts(52.287217, 21.108757)
	mux.HandleFunc("/measurements/point", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, opt.opts)
		fmt.Fprint(w, mockMeasurementResponse)
	})

	got, err := client.Measurement.ForPoint(opt)
	if err != nil {
		t.Errorf("Measurement.ForPoint: %v", err)
	}
	if !reflect.DeepEqual(got, mockMeasurement) {
		t.Errorf("Measurement.ForPoint returned %+v, want %+v", got, mockMeasurement)
	}
}

var mockMeasurementResponse = `
	{
		"current":{
			"fromDateTime":"2020-05-07T14:00:00.000Z",
			"tillDateTime":"2020-05-07T15:00:00.000Z",
			"values":[
				{
					"name":"PM1",
					"value":2.73
				}
			],
			"indexes":[
				{
					"name":"AIRLY_CAQI",
					"value":6.7,
					"level":"VERY_LOW",
					"description":"Great air here today!",
					"advice":"Perfect air for exercising! Go for it!",
					"color":"#6BC926"
				}
			],
			"standards":[
				{
					"name":"WHO",
					"pollutant":"PM25",
					"limit":25.0,
					"percent":16.08,
					"averaging":"24h"
				}
			]
		},
		"history":[
			{
				"fromDateTime":"2020-05-06T15:00:00.000Z",
				"tillDateTime":"2020-05-06T16:00:00.000Z",
				"values":[
					{
						"name":"PM1",
						"value":14.59
					}
				],
				"indexes":[
					{
						"name":"AIRLY_CAQI",
						"value":36.54,
						"level":"LOW",
						"description":"Air is quite good.",
						"advice":"Take a deep breath. Today, you can. ;)",
						"color":"#D1CF1E"
					}
				],
				"standards":[
					{
						"name":"WHO",
						"pollutant":"PM25",
						"limit":25.0,
						"percent":87.7,
						"averaging":"24h"
					}
				]
			}
		],
		"forecast":[
			{
				"fromDateTime":"2020-05-07T15:00:00.000Z",
				"tillDateTime":"2020-05-07T16:00:00.000Z",
				"values":[
					{
						"name":"PM25",
						"value":3.87
					}
				],
				"indexes":[
					{
						"name":"AIRLY_CAQI",
						"value":7.99,
						"level":"VERY_LOW",
						"description":"Great air here today!",
						"advice":"Dear me, how wonderful!",
						"color":"#6BC926"
					}
				],
				"standards":[
					{
						"name":"WHO",
						"pollutant":"PM25",
						"limit":25.0,
						"percent":15.49,
						"averaging":"24h"
					}
				]
			}
		]
	}
`

var mockMeasurement = Measurement{
	Current: Current{
		FromDateTime: time.Date(2020, 5, 7, 14, 0, 0, 0, time.UTC),
		TillDateTime: time.Date(2020, 5, 7, 15, 0, 0, 0, time.UTC),
		Values: []Value{
			{
				Name:  "PM1",
				Value: 2.73,
			},
		},
		Indexes: []Index{
			{
				Name:        "AIRLY_CAQI",
				Value:       6.7,
				Level:       "VERY_LOW",
				Description: "Great air here today!",
				Advice:      "Perfect air for exercising! Go for it!",
				Color:       "#6BC926",
			},
		},
		Standards: []Standard{
			{
				Name:      "WHO",
				Pollutant: "PM25",
				Limit:     25.0,
				Percent:   16.08,
				Averaging: "24h",
			},
		},
	},
	History: []History{
		{
			FromDateTime: time.Date(2020, 5, 6, 15, 0, 0, 0, time.UTC),
			TillDateTime: time.Date(2020, 5, 6, 16, 0, 0, 0, time.UTC),
			Values: []Value{
				{
					Name:  "PM1",
					Value: 14.59,
				},
			},
			Indexes: []Index{
				{
					Name:        "AIRLY_CAQI",
					Value:       36.54,
					Level:       "LOW",
					Description: "Air is quite good.",
					Advice:      "Take a deep breath. Today, you can. ;)",
					Color:       "#D1CF1E",
				},
			},
			Standards: []Standard{
				{
					Name:      "WHO",
					Pollutant: "PM25",
					Limit:     25.0,
					Percent:   87.7,
					Averaging: "24h",
				},
			},
		},
	},
	Forecast: []Forecast{
		{
			FromDateTime: time.Date(2020, 5, 7, 15, 0, 0, 0, time.UTC),
			TillDateTime: time.Date(2020, 5, 7, 16, 0, 0, 0, time.UTC),
			Values: []Value{
				{
					Name:  "PM25",
					Value: 3.87,
				},
			},
			Indexes: []Index{
				{
					Name:        "AIRLY_CAQI",
					Value:       7.99,
					Level:       "VERY_LOW",
					Description: "Great air here today!",
					Advice:      "Dear me, how wonderful!",
					Color:       "#6BC926",
				},
			},
			Standards: []Standard{
				{
					Name:      "WHO",
					Pollutant: "PM25",
					Limit:     25.0,
					Percent:   15.49,
					Averaging: "24h",
				},
			},
		},
	},
}
