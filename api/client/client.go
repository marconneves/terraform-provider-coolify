package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	hostname   string
	apiToken   string
	httpClient *http.Client
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string, apiToken string) *Client {
	return &Client{
		hostname:   hostname,
		apiToken:   apiToken,
		httpClient: &http.Client{},
	}
}

func (client *Client) httpRequest(path, method string, body ...bytes.Buffer) (closer io.ReadCloser, err error) {
	url := client.requestPath(path)
	var bodyBuffer bytes.Buffer

	if len(body) > 0 {
		bodyBuffer = body[0]
	}

	req, err := http.NewRequest(method, url, &bodyBuffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+client.apiToken)
	if bodyBuffer.Len() > 0 {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthenticated")
	}

	if resp.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("invalid token")
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}

	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return c.hostname + "/" + path
}
