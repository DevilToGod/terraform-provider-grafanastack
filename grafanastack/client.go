package grafanastack

import (
	"net/http"
	"net/url"
	"path"
)

// It should be initialized using the NewClient function in this package.
// Client holds all of the information required to connect to a server
type Client struct {
	client      *http.Client
	accessToken string
	baseURL     string
}

/*func NewClient(baseURL string, accessToken string) *Client {
	client := &http.Client{}

	u, err := url.Parse(baseURL + "/instances")
	if err != nil {
		return nil
	}

	u.Path = path.Clean(u.Path)

	return &Client{
		client:      client,
		accessToken: accessToken,
		baseURL:     u.String(),
	}
}*/

func NewClient(baseURL, accessToken string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	u, err := url.Parse(baseURL + "/api/v1")
	if err != nil {
		return nil
	}

	u.Path = path.Clean(u.Path)

	return &Client{
		client:      client,
		accessToken: accessToken,
		baseURL:     u.String(),
	}
}
