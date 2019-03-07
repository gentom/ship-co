package shipco

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	/*
		rel, err := url.Parse(pathStr)
		if err != nil {
			return nil, err
		}
		u := c.URL.ResolveReference(rel)
	*/

	c.URL.Path = path.Join(c.URL.Path, pathStr)
	fmt.Println(c.URL.String())

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
