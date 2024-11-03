package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	hostname   string
	apiToken   string
	httpClient *http.Client

	team   *TeamInstance
	server *ServerInstance
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string, apiToken string) *Client {
	client := &Client{
		hostname:   hostname,
		apiToken:   apiToken,
		httpClient: &http.Client{},
	}

	client.team = &TeamInstance{client: client}
	client.server = &ServerInstance{client: client}

	return client
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
	return c.hostname + "/api/v1/" + path
}

func (c *Client) HeathCheck() (*string, error) {
	body, err := c.httpRequest("healthcheck", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	defer body.Close()

	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var healCheckResponse string

	if string(bodyBytes) == "OK" {
		healCheckResponse = "success"
	} else {
		healCheckResponse = "failure"
	}

	return &healCheckResponse, nil
}

func decodeResponse[T any](body io.ReadCloser, target *T) (*T, error) {
	err := json.NewDecoder(body).Decode(target)
	if err != nil {
		return nil, err
	}

	return target, nil
}

func encodeRequest[T any](target *T) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(target)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
