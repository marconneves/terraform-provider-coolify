package database_postgresql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// DeletePostgres deletes a PostgreSQL database resource.
func (r *PostgresResource) DeletePostgres(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DatabasePostgresModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Database.Delete(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete PostgreSQL database, got error: %s", err))
		return
	}
}
