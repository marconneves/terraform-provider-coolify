package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreateDestinationDTO struct {
	Name               string `json:"name"`
	Network            string `json:"network"`
	RemoteEngine       bool   `json:"remoteEngine"`
	Engine             string `json:"engine"`
	IsCoolifyProxyUsed bool   `json:"isCoolifyProsyUsed"`
}

type CreateDestinationResponse struct {
	Id string `json:"id"`
}

func (c *Client) NewDestination(destination *CreateDestinationDTO) (*string, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(destination)
	if err != nil {
		return nil, err
	}

	body, err := c.httpRequest("api/v1/destinations/new", "POST", buf)
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

type Destination struct {
	Id                 string `json:"id"`
	Network            string `json:"network"`
	Name               string `json:"name"`
	Engine             string `json:"engine"`
	RemoteEngine       bool   `json:"remoteEngine"`
	IsCoolifyProxyUsed bool   `json:"isCoolifyProsyUsed"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
}

type DestinationUnique struct {
	Destination Destination `json:"destination"`
}

func (c *Client) GetDestination(id string) (*Destination, error) {
	body, err := c.httpRequest(fmt.Sprintf("api/v1/destinations/%v", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	response := &DestinationUnique{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	return &response.Destination, nil
}

type DestinationMany struct {
	Destinations []Destination `json:"destinations"`
}

func (c *Client) GetDestinations() (*[]Destination, error) {
	body, err := c.httpRequest("api/v1/destinations", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	response := &DestinationMany{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	return &response.Destinations, nil
}

type CheckIfNetworkNameExistRequestDTO struct {
	Network string `json:"network"`
}

func (c *Client) CheckIfNetworkNameExist(networkName string) bool {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(&CheckIfNetworkNameExistRequestDTO{Network: networkName})
	if err != nil {
		return true
	}

	_, err = c.httpRequest("api/v1/destinations/check", "POST", buf)

	return err != nil
}

func (c *Client) StopDestination(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("api/v1/destinations/%v/stop", id), "POST", bytes.Buffer{})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteDestination(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("api/v1/destinations/%v", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}

	return nil
}

type UpdateDestinationDTO struct {
	Name string `json:"name"`
}

func (c *Client) UpdateNameDestination(id string, name string) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(&UpdateDestinationDTO{
		Name: name,
	})
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/destinations/%v", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}
