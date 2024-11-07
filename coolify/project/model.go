package project

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

type ProjectModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

type ProjectDataSourceModel struct {
	ProjectModel
	Environments *[]EnvironmentModel `tfsdk:"environments"`
}

type EnvironmentModel struct {
	Id          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	ProjectId   types.Int64  `tfsdk:"project_id"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

func mapCommonProjectFields(projectData *ProjectModel, project *coolify_sdk.Project) {
	projectData.Id = types.StringValue(project.UUID)
	projectData.Name = types.StringValue(project.Name)
	if project.Description != nil {
		projectData.Description = types.StringValue(*project.Description)
	} else {
		projectData.Description = types.StringNull()
	}
}

func mapProjectDataSourceModel(projectData *ProjectDataSourceModel, project *coolify_sdk.Project) {
	mapCommonProjectFields(&projectData.ProjectModel, project)

	if project.Environments != nil {
		environments := make([]EnvironmentModel, len(project.Environments))
		for i, env := range project.Environments {
			var description types.String
			if env.Description != nil {
				description = types.StringValue(*env.Description)
			} else {
				description = types.StringNull()
			}

			environments[i] = EnvironmentModel{
				Id:          types.Int64Value(env.ID),
				Name:        types.StringValue(env.Name),
				Description: description,
				ProjectId:   types.Int64Value(env.ProjectId),
				CreatedAt:   types.StringValue(env.CreatedAt.String()),
				UpdatedAt:   types.StringValue(env.UpdatedAt.String()),
			}
		}
		projectData.Environments = &environments
	}
}

func mapProjectResourceModel(projectData *ProjectModel, project *coolify_sdk.Project) {
	mapCommonProjectFields(projectData, project)
}
