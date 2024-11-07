package project

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

var _ resource.Resource = &ProjectResource{}
var _ resource.ResourceWithImportState = &ProjectResource{}

func NewProjectResource() resource.Resource {
	return &ProjectResource{}
}

type ProjectResource struct {
	client *coolify_sdk.Sdk
}

type ProjectResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

func (r *ProjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (r *ProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage Coolify projects",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Project identifier",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Project name",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Project description",
				Optional:            true,
			},
		},
	}
}

func (r *ProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &r.client)
}

func (r *ProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ProjectResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Planned data before creation", map[string]interface{}{
		"name":        data.Name.ValueString(),
		"description": data.Description.ValueString(),
	})

	createDTO := coolify_sdk.CreateProjectDTO{
		Name:        data.Name.ValueStringPointer(),
		Description: data.Description.ValueStringPointer(),
	}

	projectID, err := r.client.Project.Create(&createDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create project, got error: %s", err))
		return
	}

	data.Id = types.StringValue(*projectID)
	tflog.Trace(ctx, "Created a project", map[string]interface{}{"project_id": projectID})

	tflog.Debug(ctx, "Data after project creation", map[string]interface{}{
		"id":          data.Id.ValueString(),
		"name":        data.Name.ValueString(),
		"description": data.Description.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	stateFilePath := filepath.Join(".", "project_state_after_create.json")
	stateData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		resp.Diagnostics.AddError("State Serialization Error", fmt.Sprintf("Unable to serialize state to JSON: %s", err))
		return
	}

	err = os.WriteFile(stateFilePath, stateData, 0644)
	if err != nil {
		resp.Diagnostics.AddError("File Write Error", fmt.Sprintf("Unable to write state to file: %s", err))
		return
	}

	tflog.Debug(ctx, "Project state saved to file after creation", map[string]interface{}{"file_path": stateFilePath})
}

func (r *ProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ProjectResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	project, err := r.client.Project.Get(data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read project, got error: %s", err))
		return
	}

	data.Id = types.StringValue(project.UUID)
	data.Name = types.StringValue(project.Name)
	if project.Description != nil {
		data.Description = types.StringValue(*project.Description)
	}

	tflog.Trace(ctx, "Read project", map[string]interface{}{"project_id": project.ID})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ProjectResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() {
		resp.Diagnostics.AddError("ID Missing", "UUID is required for updating a project.")
		return
	}

	tflog.Debug(ctx, "Project ID for update", map[string]interface{}{"project_id": data.Id.ValueString()})

	updateDTO := coolify_sdk.UpdateProjectDTO{
		Name:        data.Name.ValueStringPointer(),
		Description: data.Description.ValueStringPointer(),
	}

	err := r.client.Project.Update(data.Id.ValueString(), &updateDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update project, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated project", map[string]interface{}{"project_id": data.Id.ValueString()})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ProjectResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Project.Delete(data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete project, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted project", map[string]interface{}{"project_id": data.Id.ValueString()})
}

func (r *ProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
