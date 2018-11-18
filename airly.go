package airly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	defaultBaseURL = &url.URL{Host: "airapi.airly.eu", Scheme: "https", Path: "/v2/"}
)

// Client is a client for working with the Airly API.
type Client struct {
	http    *http.Client
	baseURL *url.URL
	token   string
}

// NewClient creates a Client that will use the specified access token for its API requests.
func NewClient(token string) Client {
	return Client{
		http:    http.DefaultClient,
		baseURL: defaultBaseURL,
		token:   token,
	}
}

// Error represents an error returned by the Airly API.
type Error struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Details   struct {
		Violations []struct {
			Parameter     string `json:"parameter"`
			Message       string `json:"message"`
			RejectedValue int    `json:"rejectedValue"`
		} `json:"violations"`
	} `json:"details"`
}

func (e Error) Error() string {
	return e.Message
}

func (c *Client) decodeError(resp *http.Response) error {
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(responseBody) == 0 {
		return fmt.Errorf("airly: HTTP %d %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(responseBody)

	var e Error
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("airly: couldn't decode error: HTTP %d %s", len(responseBody), responseBody)
	}

	if e.Message == "" {
		e.Message = fmt.Sprintf("airly: unexpected HTTP %d %s (empty error)",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return e
}

func (c *Client) get(path string, params url.Values, result interface{}) error {
	if params == nil {
		params = url.Values{}
	}

	u := c.baseURL.ResolveReference(&url.URL{Path: path, RawQuery: params.Encode()})

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("apikey", c.token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.decodeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
