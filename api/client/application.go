package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreateApplicationResponse struct {
	Id string `json:"id"`
}

func (c *Client) NewApplication() (*string, error) {
	body, err := c.httpRequest("api/v1/applications/new", "POST", bytes.Buffer{})
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

type UpdateApplicationDTO struct {
	Name string  `json:"name"`
	Fqdn *string `json:"fqdn"`
	Port *string `json:"port"`
	Type string  `json:"type"`

	PublishDirectory           *string `json:"publishDirectory"`
	DockerComposeFile          *string `json:"dockerComposeFile"`
	DockerComposeFileLocation  *string `json:"dockerComposeFileLocation"`
	DockerComposeConfiguration string  `json:"dockerComposeConfiguration"`

	GitSourceId string `json:"gitSourceId"`
	ProjectId   int    `json:"projectId"`
	Repository  string `json:"repository"`
	Branch      string `json:"branch"`

	IsCoolifyBuildPack bool   `json:"isCoolifyBuildPack"`
	BuildPack          string `json:"buildPack"`
	BaseImage          string `json:"baseImage"`
	BaseBuildImage     string `json:"baseBuildImage"`
	InstallCommand     string `json:"installCommand"`
	BuildCommand       string `json:"buildCommand"`
	StartCommand       string `json:"startCommand"`
}

func (c *Client) UpdateApplication(id string, application *UpdateApplicationDTO) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(application)
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/applications/%v", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}

type SetSourceDTO struct {
	GitSourceId string `json:"gitSourceId"`
}

func (c *Client) SetSourceOnApplication(id string, sourceId string) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(&SetSourceDTO{
		GitSourceId: sourceId,
	})
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/applications/%v/configuration/source", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}

type SetDestinationDTO struct {
	DestinationId string `json:"destinationId"`
}

func (c *Client) SetDestinationOnApplication(id string, destinationId string) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(&SetDestinationDTO{
		DestinationId: destinationId,
	})
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/applications/%v/configuration/destination", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}

type SetRepositoryDTO struct {
	ProjectId  int    `json:"projectId"`
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
	AutoDeploy bool   `json:"autodeploy"`
}

func (c *Client) SetRepositoryOnApplication(id string, repository *SetRepositoryDTO) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(repository)
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/applications/%v/configuration/repository", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}

type DeployApplicationDTO struct {
	PullMergeRequestId *string `json:"pullmergeRequestId"`
	Branch             string  `json:"branch"`
	ForceRebuild       bool    `json:"forceRebuild"`
}
type DeployApplicationResponse struct {
	BuildId string `json:"buildId"`
}

func (c *Client) DeployApplication(id string, deploy *DeployApplicationDTO) (*string, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(deploy)
	if err != nil {
		return nil, err
	}

	body, err := c.httpRequest(fmt.Sprintf("api/v1/applications/%v/deploy", id), "POST", buf)
	if err != nil {
		return nil, err
	}

	response := &DeployApplicationResponse{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	return &response.BuildId, nil

}
