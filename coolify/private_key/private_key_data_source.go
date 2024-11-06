package private_key

import (
	"context"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &PrivateKeyDataSource{}

func NewPrivateKeyDataSource() datasource.DataSource {
	return &PrivateKeyDataSource{}
}

type PrivateKeyDataSource struct {
	client *coolify_sdk.Sdk
}

func (d *PrivateKeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_key"
}

func (d *PrivateKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Private Key data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "Private Key identifier",
				Computed:            true,
			},
			"uuid": schema.StringAttribute{
				MarkdownDescription: "Private Key unique identifier",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Private Key name",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Private Key description",
				Computed:            true,
			},
			"private_key": schema.StringAttribute{
				MarkdownDescription: "Private Key content",
				Computed:            true,
			},
			"is_git_related": schema.BoolAttribute{
				MarkdownDescription: "Indicates if the key is related to Git",
				Computed:            true,
			},
			"team_id": schema.Int64Attribute{
				MarkdownDescription: "Team identifier associated with the key",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Private Key creation timestamp",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Private Key last update timestamp",
				Computed:            true,
			},
		},
	}
}

func (d *PrivateKeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	resp.Diagnostics.Append(configure.ConfigureClient(ctx, req, &d.client)...)
}

func (d *PrivateKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var privateKey PrivateKeyModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &privateKey)...)
	if resp.Diagnostics.HasError() {
		return
	}

	privateKeySaved, diags := readPrivateKey(ctx, *d.client, privateKey.UUID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mapPrivateKeyModel(&privateKey, privateKeySaved)

	tflog.Trace(ctx, "Successfully read private key data", map[string]interface{}{
		"private_key_id": privateKeySaved.ID,
		"name":           privateKeySaved.Name,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &privateKey)...)
}
