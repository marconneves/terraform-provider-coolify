package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreateDBResponse struct {
	Id string `json:"id"`
}

func (c *Client) NewDatabase() (*string, error) {
	body, err := c.httpRequest("api/v1/databases/new", "POST", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	response := &CreateApplicationResponse{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	return &response.Id, nil
}

type Database struct {
	PrivatePort int `json:"privatePort"`
	Database    struct {
		Id                  string `json:"id"`
		Name                string `json:"name"`
		PublicPort          *int   `json:"publicPort"`
		DefaultDatabase     string `json:"defaultDatabase"`
		User                string `json:"dbUser"`
		Password            string `json:"dbUserPassword"`
		RootUser            string `json:"rootUser"`
		RootPassword        string `json:"rootUserPassword"`
		Type                string `json:"type"`
		Version             string `json:"version"`
		DestinationDockerId string `json:"destinationDockerId"`
		CreatedAt           string `json:"createdAt"`
		UpdatedAt           string `json:"updatedAt"`
		Settings            struct {
			IsPublic   bool `json:"isPublic"`
			AppendOnly bool `json:"appendOnly"`
		} `json:"settings"`
	} `json:"database"`
	Settings struct {
		IpV4 *string `json:"ipv4"`
		IpV6 *string `json:"ipv6"`
	} `json:"settings"`
}

func (c *Client) GetDatabase(id string) (*Database, error) {
	body, err := c.httpRequest(fmt.Sprintf("api/v1/databases/%v", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	response := &Database{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type SetEngineDatabaseRequestDTO struct {
	Type string `json:"type"`
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
	DestinationId string `json:"destinationId"`
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

type UpdateDatabaseDTO struct {
	Name    string `json:"name"`
	Version string `json:"version"`

	DefaultDatabase  string `json:"defaultDatabase"`
	DbUser           string `json:"dbUser"`
	DbUserPassword   string `json:"dbUserPassword"`
	RootUser         string `json:"rootUser"`
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

type UpdateNameDatabaseDTO struct {
	Name string `json:"name"`
}

func (c *Client) UpdateNameDatabase(id string, name string) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(&UpdateNameDatabaseDTO{Name: name})
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/databases/%v", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) StartDatabase(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("api/v1/databases/%v/start", id), "POST", bytes.Buffer{})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) StopDatabase(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("api/v1/databases/%v/stop", id), "POST", bytes.Buffer{})
	if err != nil {
		return err
	}

	return nil
}

type DeleteDatabaseRequestDTO struct {
	Force bool `json:"force"`
}

func (c *Client) DeleteDatabase(id string) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(&DeleteDatabaseRequestDTO{Force: true})
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/databases/%v", id), "DELETE", buf)
	if err != nil {
		return err
	}

	return nil
}

type UpdateSettingsDatabaseDTO struct {
	IsPublic   bool `json:"isPublic"`
	AppendOnly bool `json:"appendOnly"`
}
type UpdateSettingsDatabase struct {
	PublicPort *int `json:"publicPort"`
}

func (c *Client) UpdateSettings(id string, settings *UpdateSettingsDatabaseDTO) (*UpdateSettingsDatabase, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(settings)
	if err != nil {
		return nil, err
	}

	body, err := c.httpRequest(fmt.Sprintf("api/v1/databases/%v/settings", id), "POST", buf)
	if err != nil {
		return nil, err
	}

	response := &UpdateSettingsDatabase{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
