package client

import (
	"bytes"
	"encoding/json"
)

// {"name":"Local Docker","engine":"/var/run/docker.sock","remoteEngine":false,"network":"clblhbffr00003b5m8jejf511","isCoolifyProxyUsed":true}
type CreateDestinationDTO struct {
	Name string `json:"name"`
	Network string `json:"network"`
	RemoteEngine bool `json:"remoteEngine"`
	Engine string `json:"engine"`
	IsCoolifyProxyUsed bool `json:"isCoolifyProsyUsed"`
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
	if err != nil {
		return true
	}

	return false
}