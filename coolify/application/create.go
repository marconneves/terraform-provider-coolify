package application

import (
	"context"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Status struct {
	Domain string
}

type Secret struct {
	Name       string
	Value      string
	IsBuildEnv bool
}

type Application struct {
	Name   string
	Domain string
	IsBot  bool

	Template struct {
		BuildPack  string
		Image      string
		BuildImage string

		Settings struct {
			IsCoolifyBuildPack bool
			InstallCommand     string
			BuildCommand       string
			StartCommand       string
		}
	}

	Secrets []Secret

	Repository struct {
		ProjectId  int
		Repository string
		Branch     string
		AutoDeploy bool
	}

	Settings struct {
		DestinationId string
		SourceId      string
	}

	Status Status
}

func applicationCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// 1. New Application
	// 2. Configuration -> Set Source
	// 3. Configuration -> Set Repository
	// 4. Configuration -> Set Destination
	// 5. Configuration -> Set BuildPack
	// 6. Update Application - Setup NodeJS
	// if is bot
	// 7. Update Application - Setup Bot
	// if not is bot
	// 7. Update Application - Setup Web domain
	// 8. Request Deploy

	secretOne := Secret{
		Name:       "SECRET",
		Value:      "VALUE",
		IsBuildEnv: true,
	}
	status := make(map[string]string)

	secrets := []Secret{secretOne}

	app := &Application{
		IsBot:   false,
		Secrets: secrets,
	}
	app.Name = d.Get("name").(string)
	app.Domain = "https://terraform.s.b4.run"
	app.Template.BuildPack = "node"
	app.Template.Image = "node:14"
	app.Template.BuildImage = "node:14"
	app.Template.Settings.InstallCommand = "npm install"
	app.Template.Settings.BuildCommand = ""
	app.Template.Settings.StartCommand = "npm start"
	app.Template.Settings.IsCoolifyBuildPack = true

	app.Repository.ProjectId = 579493141
	app.Repository.Repository = "cool-sample/sample-nodejs"
	app.Repository.Branch = "main"
	app.Repository.AutoDeploy = true

	app.Settings.DestinationId = "clb9wrx87001fmo9dvvog6xet"
	app.Settings.SourceId = "clb9y09gs000f9dmod69f7dce"

	apiClient := m.(*client.Client)

	id, err := apiClient.NewApplication()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*id)

	err = apiClient.SetSourceOnApplication(*id, app.Settings.SourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	repository := &client.SetRepositoryDTO{
		ProjectId:  app.Repository.ProjectId,
		Repository: app.Repository.Repository,
		Branch:     app.Repository.Branch,
		AutoDeploy: app.Repository.AutoDeploy,
	}
	err = apiClient.SetRepositoryOnApplication(*id, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	err = apiClient.SetDestinationOnApplication(*id, app.Settings.DestinationId)
	if err != nil {
		return diag.FromErr(err)
	}

	appToUpdate := &client.UpdateApplicationDTO{
		Name: app.Name,
		Fqdn: &app.Domain,
		Port: nil,
		Type: "base",

		PublishDirectory:           nil,
		DockerComposeFile:          nil,
		DockerComposeFileLocation:  nil,
		DockerComposeConfiguration: "{}",

		IsCoolifyBuildPack: true,
		BuildPack:          app.Template.BuildPack,
		BaseImage:          app.Template.Image,
		BaseBuildImage:     app.Template.BuildImage,
		InstallCommand:     app.Template.Settings.InstallCommand,
		BuildCommand:       app.Template.Settings.BuildCommand,
		StartCommand:       app.Template.Settings.StartCommand,
	}

	err = apiClient.UpdateApplication(*id, appToUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	// DeployApplication
	deploy := &client.DeployApplicationDTO{
		PullMergeRequestId: nil,
		Branch:             app.Repository.Branch,
		ForceRebuild:       true,
	}
	deployId, err := apiClient.DeployApplication(*id, deploy)

	status["deployId"] = *deployId

	// TODO: Await deploy finish

	d.Set("status", status)

	return nil
}
