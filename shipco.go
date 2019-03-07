package shipco

import (
	"net/http"
	"net/url"
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
