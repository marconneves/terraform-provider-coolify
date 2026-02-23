package database_postgresql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// ReadPostgres reads a PostgreSQL database resource.
func (r *PostgresResource) ReadPostgres(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DatabasePostgresModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	db, err := r.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read PostgreSQL database, got error: %s", err))
		return
	}

	mapPostgresResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
