package airly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

var (
	// DefaultBaseURL is an URL for interacting with the Airly API.
	DefaultBaseURL = &url.URL{
		Host:   "airapi.airly.eu",
		Scheme: "https",
		Path:   "/v2/",
	}
)

// Client is a client for working with the Airly API.
type Client struct {
	http  *http.Client
	token string
}

// NewClient creates a Client that will use the specified access token
// for its API requests.
func NewClient(token string) Client {
	return Client{
		http:  http.DefaultClient,
		token: token,
	}
}

// Violation represents an error which requested value is invalid.
type Violation struct {
	Parameter     string `json:"parameter"`
	Message       string `json:"message"`
	RejectedValue int    `json:"rejectedValue"`
}

// Details represents list of violations when interacting with the Airly API.
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
		return errors.Wrap(err, "decode response")
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
	if params == nil {
		params = url.Values{}
	}

	u := DefaultBaseURL.ResolveReference(
		&url.URL{
			Path:     path,
			RawQuery: params.Encode(),
		},
	)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return errors.Wrap(err, "http.NewRequest")
	}

	req.Header.Add("apikey", c.token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http.Do")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.decodeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return errors.Wrap(err, "decode response")
	}

	return nil
}
