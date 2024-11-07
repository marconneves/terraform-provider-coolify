package provider

import (
	"context"
	"os"

	"github.com/marconneves/terraform-provider-coolify/coolify/private_key"
	"github.com/marconneves/terraform-provider-coolify/coolify/project"
	"github.com/marconneves/terraform-provider-coolify/coolify/project_environment"
	"github.com/marconneves/terraform-provider-coolify/coolify/server"
	"github.com/marconneves/terraform-provider-coolify/coolify/team"
	"github.com/marconneves/terraform-provider-coolify/coolify/team_members"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
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

const (
	ENV_KEY_ADDRESS = "COOLIFY_ADDRESS"
	ENV_KEY_TOKEN   = "COOLIFY_TOKEN"

	DEFAULT_COOLIFY_ENDPOINT = "https://app.coolify.io"
)

func (p *CoolifyProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "coolify"
	resp.Version = p.version
}

func (p *CoolifyProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	hasEnvToken := os.Getenv(ENV_KEY_TOKEN) != ""
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				MarkdownDescription: "Coolify endpoint. If not set, checks env for `" + ENV_KEY_ADDRESS + "`. Default: `" + DEFAULT_COOLIFY_ENDPOINT + "`.",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				Required:            !hasEnvToken,
				Optional:            hasEnvToken,
				Sensitive:           true,
				MarkdownDescription: "Coolify token. If not set, checks env for `" + ENV_KEY_TOKEN + "`.",
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

	apiEndpoint := getAPIEndpoint(data.Address)
	apiToken := getAPIToken(data.Token, resp)

	if resp.Diagnostics.HasError() {
		return
	}

	client := coolify_sdk.Init(apiEndpoint, apiToken)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func getAPIEndpoint(address types.String) string {
	if !address.IsNull() {
		return address.ValueString()
	}
	if apiEndpointFromEnv, found := os.LookupEnv(ENV_KEY_ADDRESS); found {
		return apiEndpointFromEnv
	}
	return DEFAULT_COOLIFY_ENDPOINT
}

func getAPIToken(token types.String, resp *provider.ConfigureResponse) string {
	if !token.IsNull() {
		return token.ValueString()
	}
	if apiTokenFromEnv, found := os.LookupEnv(ENV_KEY_TOKEN); found {
		return apiTokenFromEnv
	}
	resp.Diagnostics.AddAttributeError(path.Root("token"), "Failed to configure client", "No token provided")
	return ""
}

func (p *CoolifyProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		project.NewProjectResource,
	}
}

func (p *CoolifyProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		team.NewTeamDataSource,
		team_members.NewTeamMembersDataSource,
		project.NewProjectDataSource,
		project_environment.NewEnvironmentDataSource,
		server.NewServerDataSource,
		private_key.NewPrivateKeyDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CoolifyProvider{
			version: version,
		}
	}
}
