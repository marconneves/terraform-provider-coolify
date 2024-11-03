package application

import (
	"context"
	"strings"

	sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func applicationReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*sdk.Client)
	destinationId := d.Id()

	item, err := apiClient.GetApplication(destinationId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return diag.Errorf("error finding Item with ID %v", destinationId)
		}
	}

	d.SetId(item.Application.Id)
	d.Set("name", item.Application.Name)
	d.Set("domain", item.Application.Fqdn)
	d.Set("is_bot", item.Application.Settings.IsBot)

	template := make(map[string]interface{})
	template["build_pack"] = item.Application.BuildPack
	template["image"] = item.Application.BaseImage
	template["build_image"] = item.Application.BaseBuildImage

	templateSettings := make(map[string]interface{})
	templateSettings["install_command"] = item.Application.InstallCommand
	templateSettings["build_command"] = item.Application.BuildCommand
	templateSettings["start_command"] = item.Application.StartCommand
	//TODO: IsCoolifyBuildPack is defined true by default
	templateSettings["auto_deploy"] = item.Application.Settings.AutoDeploy
	template["settings"] = []interface{}{templateSettings}

	templateEnv := []interface{}{}
	for _, secret := range item.Application.Secrets {
		// TODO: Validate secret of preview
		if secret.IsPRMRSecret == false {
			env := make(map[string]interface{})

			env["key"] = secret.Name
			env["value"] = secret.Value
			env["is_build_env"] = secret.IsBuildSecret

			templateEnv = append(templateEnv, env)
		}
	}
	template["env"] = templateEnv

	d.Set("template", []interface{}{template})

	status := make(map[string]string)
	d.Set("status", status)

	return nil
}
