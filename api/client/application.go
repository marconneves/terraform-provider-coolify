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

type Application struct {
	Application struct {
		Id                  string  `json:"id"`
		Name                string  `json:"name"`
		Fqdn                string  `json:"fqdn"`
		BuildPack           string  `json:"buildPack"`
		BaseImage           string  `json:"baseImage"`
		BaseBuildImage      string  `json:"baseBuildImage"`
		InstallCommand      string  `json:"installCommand"`
		BuildCommand        string  `json:"buildCommand"`
		StartCommand        string  `json:"startCommand"`
		Repository          string  `json:"repository"`
		RepositoryId        int     `json:"projectId"`
		Branch              string  `json:"branch"`
		CommitHash          string  `json:"commitHash"`
		GitCommitHash       *string `json:"gitCommitHash"`
		GitSourceId         string  `json:"gitSourceId"`
		DestinationDockerId string  `json:"destinationDockerId"`
		Settings            struct {
			AutoDeploy bool `json:"autodeploy"`
			IsBot      bool `json:"isBot"`
		} `json:"settings"`
		Secrets []struct {
			Id            string `json:"id"`
			Name          string `json:"name"`
			Value         string `json:"value"`
			IsBuildSecret bool   `json:"isBuildSecret"`
			IsPRMRSecret  bool   `json:"isPRMRSecret"`
		} `json:"secrets"`
	} `json:"application"`
}

func (c *Client) GetApplication(id string) (*Application, error) {
	body, err := c.httpRequest(fmt.Sprintf("api/v1/applications/%v", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	response := &Application{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// need add body with force
type DeleteApplicationDTO struct {
	Force bool `json:"force"`
}

func (c *Client) DeleteApplication(id string) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(DeleteApplicationDTO{Force: true})
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/applications/%v", id), "DELETE", buf)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) StopApplication(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("api/v1/applications/%v/stop", id), "POST", bytes.Buffer{})
	if err != nil {
		return err
	}

	return nil
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

type ApplicationEnvironmentDTO struct {
	Name          string `json:"name"`
	Value         string `json:"value"`
	IsBuildEnv    bool   `json:"isBuildSecret"`
	IsNew         bool   `json:"isNew"`
	PreviewSecret bool   `json:"previewSecret"`
}

func (c *Client) AddEnvironmentToApplication(id string, environment *ApplicationEnvironmentDTO) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(environment)
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/applications/%v/secrets", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}

type DeleteApplicationEnvironmentDTO struct {
	Name string `json:"name"`
}

func (c *Client) DeleteEnvironmentFromApplication(id string, name string) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(&DeleteApplicationEnvironmentDTO{
		Name: name,
	})
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/applications/%v/secrets", id), "DELETE", buf)
	if err != nil {
		return err
	}

	return nil
}

type UpdateApplicationSettingsDTO struct {
	Debug         *bool   `json:"debug,omitempty"`
	Previews      *bool   `json:"previews,omitempty"`
	DualCerts     *bool   `json:"dualCerts,omitempty"`
	AutoDeploy    *bool   `json:"autodeploy,omitempty"`
	Branch        *string `json:"branch,omitempty"`
	ProjectId     *string `json:"projectId,omitempty"`
	IsBot         *bool   `json:"isBot,omitempty"`
	IsDBBranching *bool   `json:"isDBBranching,omitempty"`
	IsCustomSSL   *bool   `json:"isCustomSSL,omitempty"`
}

func (c *Client) UpdateApplicationSettings(id string, settings *UpdateApplicationSettingsDTO) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(settings)
	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("api/v1/applications/%v/settings", id), "POST", buf)
	if err != nil {
		return err
	}

	return nil
}
