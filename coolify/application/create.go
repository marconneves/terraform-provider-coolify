package application

import (
	"context"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func applicationCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	status := make(map[string]string)

	app := &Application{}

	app.Name = d.Get("name").(string)
	app.Domain = d.Get("domain").(string)
	app.IsBot = d.Get("is_bot").(bool)
	// TODO: When is bot, not accept a domain

	for _, template := range d.Get("template").([]interface{}) {
		i := template.(map[string]interface{})

		app.Template.BuildPack = i["build_pack"].(string)
		// TODO: validate if images and commands is valid for image

		app.Template.Image = i["image"].(string)
		app.Template.BuildImage = i["build_image"].(string)

		for _, settings := range i["settings"].([]interface{}) {
			j := settings.(map[string]interface{})

			app.Template.Settings.InstallCommand = j["install_command"].(string)
			app.Template.Settings.BuildCommand = j["build_command"].(string)
			app.Template.Settings.StartCommand = j["start_command"].(string)
			app.Template.Settings.IsCoolifyBuildPack = true
			app.Template.Settings.AutoDeploy = j["auto_deploy"].(bool)
		}

		app.Template.Envs = []Env{}

		for _, env := range i["env"].([]interface{}) {
			j := env.(map[string]interface{})

			secretOne := Env{
				Key:        j["key"].(string),
				Value:      j["value"].(string),
				IsBuildEnv: j["is_build_env"].(bool),
			}

			app.Template.Envs = append(app.Template.Envs, secretOne)
		}
	}

	// TODO: validate when is a docker project to not deploy using repository
	for _, repository := range d.Get("repository").([]interface{}) {
		i := repository.(map[string]interface{})

		app.Repository.RepositoryId = i["repository_id"].(int)
		app.Repository.Repository = i["repository"].(string)
		app.Repository.Branch = i["branch"].(string)
		app.Repository.commitHash = i["commit_hash"].(string)
	}

	for _, setting := range d.Get("settings").([]interface{}) {
		i := setting.(map[string]interface{})

		app.Settings.SourceId = i["source_id"].(string)
		app.Settings.DestinationId = i["destination_id"].(string)
	}

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
		ProjectId:  app.Repository.RepositoryId,
		Repository: app.Repository.Repository,
		Branch:     app.Repository.Branch,
		AutoDeploy: app.Template.Settings.AutoDeploy,
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

	for _, env := range app.Template.Envs {
		secret := &client.ApplicationEnvironmentDTO{
			Name:          env.Key,
			Value:         env.Value,
			IsBuildEnv:    env.IsBuildEnv,
			IsNew:         true,
			PreviewSecret: false,
		}
		err = apiClient.AddEnvironmentToApplication(*id, secret)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	settingsToUpdate := &client.UpdateApplicationSettingsDTO{
		IsBot: &app.IsBot,
	}
	err = apiClient.UpdateApplicationSettings(*id, settingsToUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	deploy := &client.DeployApplicationDTO{
		PullMergeRequestId: nil,
		Branch:             app.Repository.Branch,
		ForceRebuild:       true,
	}
	_, err = apiClient.DeployApplication(*id, deploy)
	if err != nil {
		return diag.FromErr(err)
	}

	// TODO: Await deploy finish

	d.Set("status", status)

	return nil
}
