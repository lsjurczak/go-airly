package airly

// IndexType represents an air quality type.
type IndexType struct {
	Name   string `json:"name"`
	Levels []struct {
		MinValue    float64 `json:"minValue"`
		MaxValue    float64 `json:"maxValue"`
		Values      string  `json:"values"`
		Level       string  `json:"level"`
		Description string  `json:"description"`
		Color       string  `json:"color"`
	} `json:"levels"`
}

// MeasurementType represent a measurement type.
type MeasurementType struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Unit  string `json:"unit"`
}

// GetIndexTypes returns a list of all the index types supported in the API along with lists of levels defined per each index type.
func (c *Client) GetIndexTypes() (*[]IndexType, error) {
	var indexTypes []IndexType

	err := c.get("meta/indexes", nil, &indexTypes)
	if err != nil {
		return nil, err
	}

	return &indexTypes, nil
}

// GetMeasurementTypes returns list of all the measurement types supported in the API along with their names and units.
func (c *Client) GetMeasurementTypes() (*[]MeasurementType, error) {
	var measurementTypes []MeasurementType

	err := c.get("meta/measurements", nil, &measurementTypes)
	if err != nil {
		return nil, err
	}

	return &measurementTypes, nil
}
