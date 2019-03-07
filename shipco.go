package shipco

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Client (Ship&Co API Client)
type Client struct {
	Token  string `json:"token"`
	URL    *url.URL
	client *http.Client
}

// NewClient returns Client struct's instance
func NewClient(URL, token string) *Client {
	c := &Client{
		Token:  token,
		client: http.DefaultClient,
	}

	if !strings.HasSuffix(URL, "/") {
		URL += "/"
	}

	var err error
	c.URL, err = url.Parse(URL)
	if err != nil {
		panic(err)
	}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, pathStr string, body interface{}) (*http.Request, error) {
	c.URL.Path = path.Join(c.URL.Path, pathStr)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.URL.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-access-token", c.Token)
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

// Do can be used to perform the request created with NewRequest
func (c *Client) Do(req *http.Request) (interface{}, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
