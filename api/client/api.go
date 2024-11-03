package client

import (
	"encoding/json"
)

type Api struct {
	client *Client
}

func (c *Client) Api() *Api {
	return &Api{client: c}
}

type EnableApiResponse struct {
	Message string `json:"message"`
}

func (a *Api) Enable() (*string, error) {
	body, err := a.client.httpRequest("enable", "GET")
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
	body, err := a.client.httpRequest("disable", "GET")
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
