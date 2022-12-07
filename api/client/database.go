package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type CreateDBResponse struct {
	id string
}

func (c *Client) NewDatabase() (*string, error) {
	body, err := c.httpRequest("api/v1/databases/new", "POST", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	out, err := json.Marshal(body)

	f, err := os.Create("body.txt")
	f.WriteString(string(out))
	f.Close()

	response := &CreateDBResponse{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	return &response.id, nil
}

type SetEngineDatabaseRequestDTO struct {
    Type    string `json:"type"`
}

func (c *Client) SetEngineDatabase(id string, engine string) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(&SetEngineDatabaseRequestDTO{Type: engine})
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/databases/%v/configuration/type", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}

type SetDestinationDatabaseRequestDTO struct {
    DestinationId    string `json:"destinationId"`
}

func (c *Client) SetDestinationDatabase(id string, destination string) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(&SetDestinationDatabaseRequestDTO{DestinationId: destination})
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/databases/%v/configuration/destination", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}

// Update database

type UpdateDatabaseDTO struct {
	Name        string `json:"name"`
	Version 	string `json:"version"`
	
	DefaultDatabase string `json:"defaultDatabase"`
	DbUser string `json:"dbUser"`
	DbUserPassword string `json:"dbUserPassword"`
	RootUserPassword string `json:"rootUserPassword"`
}

func (c *Client) UpdateDatabase(id string, database *UpdateDatabaseDTO) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(database)
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/databases/%v", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}