package api

import (
	"net/http"
	"strings"
	"time"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func NewClient(baseURL string) *Client {
	if !strings.HasPrefix(baseURL, "https://") && !strings.HasPrefix(baseURL, "http://") {
		if strings.HasPrefix(baseURL, "localhost") {
			baseURL = "http://" + baseURL
		} else {
			baseURL = "https://" + baseURL
		}
	}

	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
