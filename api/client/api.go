package client

import (
	"bytes"
	"encoding/json"
	"io"
)

func (c *Client) HeathCheck() (*string, error) {
	body, err := c.httpRequest("api/v1/healthcheck", "GET", bytes.Buffer{})
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

type Api struct {
	client *Client
}

func (c *Client) NewApi() *Api {
	return &Api{client: c}
}

type EnableApiResponse struct {
	Message string `json:"message"`
}

func (a *Api) Enable() (*string, error) {
	body, err := a.client.httpRequest("api/v1/enable", "GET")
	if err != nil {
		return nil, err
	}

	response := &EnableApiResponse{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	var apiEnabled string
	if response.Message == "API Enabled." {
		apiEnabled = "success"
	} else {
		apiEnabled = "failure"
	}

	return &apiEnabled, nil
}

func (a *Api) Disable() (*string, error) {
	body, err := a.client.httpRequest("api/v1/disable", "GET")
	if err != nil {
		return nil, err
	}

	response := &EnableApiResponse{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	var apiDisabled string
	if response.Message == "API disabled." {
		apiDisabled = "success"
	} else {
		apiDisabled = "failure"
	}

	return &apiDisabled, nil
}
