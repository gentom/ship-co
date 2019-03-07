package shipco

import (
	"net/http"
	"net/url"
)

// Client (Ship&Co API Client)
type Client struct {
	Token  string `json:"token"`
	URL    *url.URL
	client *http.Client
}
