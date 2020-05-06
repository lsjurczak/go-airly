package airly

// https://developer.airly.eu/docs#endpoints.meta
type MetaService struct {
	client *Client
}

// Level represents a definition of a single index level.
type Level struct {
	MinValue    float64 `json:"minValue"`
	MaxValue    float64 `json:"maxValue"`
	Values      string  `json:"values"`
	Level       string  `json:"level"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
}

// IndexType represents an air quality type.
type IndexType struct {
	Name   string  `json:"name"`
	Levels []Level `json:"levels"`
}

// MeasurementType represents a measurement type.
type MeasurementType struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Unit  string `json:"unit"`
}

// Indexes return a list of all the index types supported in the API along
// with lists of levels defined per each index type.
// https://developer.airly.eu/docs#endpoints.meta.indexes
func (c *MetaService) Indexes() ([]IndexType, error) {
	var indexTypes []IndexType
	err := c.client.get("meta/indexes", nil, &indexTypes)
	if err != nil {
		return nil, err
	}
	return indexTypes, nil
}

// Measurements return a list of all the measurement types supported
// in the API along with their names and units.
// https://developer.airly.eu/docs#endpoints.meta.measurements
func (c *MetaService) Measurements() ([]MeasurementType, error) {
	var measurementTypes []MeasurementType
	err := c.client.get("meta/measurements", nil, &measurementTypes)
	if err != nil {
		return nil, err
	}

	return measurementTypes, nil
}
