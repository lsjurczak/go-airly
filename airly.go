package airly

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// HTTPDoer is a single-method interface for performing HTTP requests.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is a client for working with the Airly API.
type Client struct {
	// HTTP client used to communicate with the API.
	client HTTPDoer

	apiKey   string
	baseURL  *url.URL
	language string

	Installation *InstallationService
	Measurement  *MeasurementService
	Meta         *MetaService
}

// NewClient creates a Client that will use the specified access apiKey
// for its API requests.
func NewClient(client HTTPDoer, apiKey string) (*Client, error) {
	if client == nil {
		client = httpClient
	}

	if apiKey == "" {
		return nil, errors.New("missing api key")
	}

	c := &Client{
		client: client,
		apiKey: apiKey,
		baseURL: &url.URL{
			Host:   "airapi.airly.eu",
			Scheme: "https",
			Path:   "/v2/",
		},
	}

	c.Installation = &InstallationService{client: c}
	c.Measurement = &MeasurementService{client: c}
	c.Meta = &MetaService{client: c}

	return c, nil
}

func (c *Client) Language(lang string) *Client {
	c.language = lang
	return c
}

// Violation represents an error that the requested value is invalid.
type Violation struct {
	Parameter     string `json:"parameter"`
	Message       string `json:"message"`
	RejectedValue int    `json:"rejectedValue"`
}

// Details represent a list of violations when interacting with the Airly API.
type Details struct {
	Violations []Violation `json:"violations"`
}

// Error represents an error returned by the Airly API.
type Error struct {
	ErrorCode string  `json:"errorCode"`
	Message   string  `json:"message"`
	Details   Details `json:"details"`
}

func (e Error) Error() string {
	return e.Message
}

func (c *Client) decodeError(resp *http.Response) error {
	var e Error

	err := json.NewDecoder(resp.Body).Decode(&e)
	if err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	if e.Message == "" {
		e.Message = fmt.Sprintf(
			"airly: unexpected HTTP %d %s (empty error)",
			resp.StatusCode,
			http.StatusText(resp.StatusCode),
		)
	}

	return e
}

func (c *Client) get(path string, params url.Values, result interface{}) error {
	u := c.baseURL.ResolveReference(
		&url.URL{
			Path:     path,
			RawQuery: params.Encode(),
		},
	)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Add("apiKey", c.apiKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if c.language != "" {
		req.Header.Add("Accept-Language", c.language)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("doer.Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.decodeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}
