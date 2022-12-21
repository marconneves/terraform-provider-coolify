package application

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Status struct {
	Domain string
}

type Secret struct {
	Name string
	Value string
	IsBuildEnv bool
}

type Application struct {
	Name string
	Domain string
	IsBot bool

	Template struct {
		BuildPack string
		Image string
		BuildImage string

		Settings struct {
			IsCoolifyBuildPack bool
			InstallCommand string
			BuildCommand string
			StartCommand string
		}
	}

	Secrets []Secret

	Repository struct {
		ProjectId string
		Repository string
		branch string
		AutoDeploy bool
	}
	
	Settings struct {
		DestinationId string
		SourceId string
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
		Name: "SECRET",
		Value: "VALUE",
		IsBuildEnv: true,
	}

	secrets := []Secret{secretOne}
	
	app := &Application{
		Name: "Project Test",
		Domain: "https://terraform.s.b4.run",
		IsBot: false,
		Secrets: secrets,
		
	}
	app.Template.BuildPack = "node"
	app.Template.Image = "node:14"
	app.Template.BuildImage = "node:14"
	app.Template.Settings.InstallCommand = "npm install"
	app.Template.Settings.BuildCommand = ""
	app.Template.Settings.StartCommand = "npm start"
	app.Template.Settings.IsCoolifyBuildPack = true

	app.Repository.ProjectId = "5f9f9f9f9f9f9f9f9f9f9f9f"
	app.Repository.Repository = "5f9f9f9f9f9f9f9f9f9f9f9f"
	app.Repository.branch = "master"
	app.Repository.AutoDeploy = true




	status := make(map[string]string)
	app.Name = d.Get("name").(string)


	d.Set("status", status)

	return nil
}