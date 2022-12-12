package client

import (
	"bytes"
	"encoding/json"
)

type CreateDestinationDTO struct {
	Name string `json:"name"`
	Network string `json:"network"`
}

type CreateDestinationResponse struct {
	Id string `json:"id"`
}

func (c *Client) NewDestination(destination *CreateDestinationDTO) (*string, error) {
	body, err := c.httpRequest("api/v1/destinations/new", "POST", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	response := &CreateDestinationResponse{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	return &response.Id, nil
}