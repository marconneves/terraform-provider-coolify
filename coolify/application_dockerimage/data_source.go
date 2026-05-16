package application_dockerimage

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

var _ datasource.DataSource = &ApplicationDockerImageDataSource{}

func NewApplicationDockerImageDataSource() datasource.DataSource {
	return &ApplicationDockerImageDataSource{}
}

type ApplicationDockerImageDataSource struct {
	client *coolify_sdk.Sdk
}

func (d *ApplicationDockerImageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_dockerimage"
}

func (d *ApplicationDockerImageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Read a Coolify docker image application.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Application UUID.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Application name.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Application description.",
				Computed:            true,
			},
			"docker_registry_image_name": schema.StringAttribute{
				MarkdownDescription: "Docker registry image name.",
				Computed:            true,
			},
			"docker_registry_image_tag": schema.StringAttribute{
				MarkdownDescription: "Docker registry image tag.",
				Computed:            true,
			},
			"ports_exposes": schema.StringAttribute{
				MarkdownDescription: "Exposed ports.",
				Computed:            true,
			},
			"ports_mappings": schema.StringAttribute{
				MarkdownDescription: "Port mappings.",
				Computed:            true,
			},
			"fqdn": schema.StringAttribute{
				MarkdownDescription: "Application domain(s).",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Application status.",
				Computed:            true,
			},
			"environment_variables": schema.MapAttribute{
				MarkdownDescription: "All non-sensitive environment variables currently set on the application.",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func (d *ApplicationDockerImageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &d.client)
}

// ApplicationDockerImageDataSourceModel mirrors the data source schema. Kept
// separate from the resource model so the data source only exposes computed
// fields.
type ApplicationDockerImageDataSourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	Description             types.String `tfsdk:"description"`
	DockerRegistryImageName types.String `tfsdk:"docker_registry_image_name"`
	DockerRegistryImageTag  types.String `tfsdk:"docker_registry_image_tag"`
	PortsExposes            types.String `tfsdk:"ports_exposes"`
	PortsMappings           types.String `tfsdk:"ports_mappings"`
	FQDN                    types.String `tfsdk:"fqdn"`
	Status                  types.String `tfsdk:"status"`
	EnvironmentVariables    types.Map    `tfsdk:"environment_variables"`
}

func (d *ApplicationDockerImageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ApplicationDockerImageDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	app, err := d.client.Application.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read docker image application, got error: %s", err))
		return
	}

	data.Id = types.StringValue(app.UUID)
	data.Name = types.StringValue(app.Name)
	data.Description = configure.ValueStringValue(app.Description, data.Description)
	data.DockerRegistryImageName = configure.ValueStringValue(app.DockerRegistryImageName, data.DockerRegistryImageName)
	data.DockerRegistryImageTag = configure.ValueStringValue(app.DockerRegistryImageTag, data.DockerRegistryImageTag)
	data.PortsExposes = types.StringValue(app.PortsExposes)
	data.PortsMappings = configure.ValueStringValue(app.PortsMappings, data.PortsMappings)
	data.FQDN = configure.ValueStringValue(app.FQDN, types.StringNull())
	data.Status = types.StringValue(app.Status)

	envs, err := d.client.Application.ListEnvs(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list envs, got error: %s", err))
		return
	}
	values := map[string]string{}
	for _, e := range *envs {
		values[e.Key] = e.Value
	}
	m, diags := types.MapValueFrom(ctx, types.StringType, values)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.EnvironmentVariables = m

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
