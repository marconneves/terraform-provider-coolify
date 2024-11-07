package server

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	"github.com/marconneves/terraform-provider-coolify/coolify/private_key"
)

func readServer(ctx context.Context, client coolify_sdk.Sdk, uuid types.String, name types.String) (*coolify_sdk.Server, diag.Diagnostics) {
	var diags diag.Diagnostics

	var server *coolify_sdk.Server
	var err error

	if !uuid.IsNull() {
		server, err = fetchServerByID(client, ctx, uuid.ValueString())
	} else if !name.IsNull() {
		server, err = fetchServerByName(client, ctx, name.ValueString())
	} else {
		diags.AddError("Configuration Error", "Either 'id' or 'name' must be specified.")
		return nil, diags
	}

	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read server data: %s %s", err, uuid.ValueString()))
		return nil, diags
	}

	if server == nil {
		diags.AddError("Not Found", "No server found with the given ID or name")
		return nil, diags
	}

	return server, diags
}

func (s *ServerDataSource) ReadServerDataSource(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var server ServerDataSourceModel

	diags := req.Config.Get(ctx, &server)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	serverSaved, diags := readServer(ctx, *s.client, server.ID, server.Name)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mapServerDataSourceModel(&server, serverSaved)

	tflog.Trace(ctx, "Successfully read team data", map[string]interface{}{
		"server_uuid": serverSaved.UUID,
		"name":        serverSaved.Name,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &server)...)

}

func (r *ServerResource) ReadServerResource(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var server ServerModel

	diags := req.State.Get(ctx, &server)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	serverSaved, diags := readServer(ctx, *r.client, server.ID, server.Name)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mapServerResourceModel(&server, serverSaved)

	privateKey, err := private_key.FetchPrivateKeyByID(*r.client, ctx, serverSaved.PrivateKeyID)
	if err != nil {
		diags.AddError("Not Found", "No server found with the given ID or name")
		return
	}

	server.PrivateKeyUUID = types.StringValue(privateKey.UUID)

	tflog.Trace(ctx, "Successfully read server data", map[string]interface{}{
		"server_uuid": serverSaved.UUID,
		"name":        serverSaved.Name,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &server)...)
}
