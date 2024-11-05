package provider

import (
	"context"

	"github.com/marconneves/terraform-provider-coolify/coolify/project"
	"github.com/marconneves/terraform-provider-coolify/coolify/team"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure CoolifyProvider satisfies various provider interfaces.
var _ provider.Provider = &CoolifyProvider{}

// CoolifyProvider defines the provider implementation.
type CoolifyProvider struct {
	version string
}

// CoolifyProviderModel describes the provider data model.
type CoolifyProviderModel struct {
	Address types.String `tfsdk:"address"`
	Token   types.String `tfsdk:"token"`
}

func (p *CoolifyProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "coolify"
	resp.Version = p.version
}

func (p *CoolifyProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				MarkdownDescription: "The address of the Coolify service.",
				Required:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "The token for authenticating with the Coolify service.",
				Required:            true,
			},
		},
	}
}

func (p *CoolifyProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data CoolifyProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := coolify_sdk.Init(data.Address.ValueString(), data.Token.ValueString())
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CoolifyProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *CoolifyProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		team.NewTeamDataSource,
		team.NewTeamMembersDataSource,
		project.NewProjectDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CoolifyProvider{
			version: version,
		}
	}
}
