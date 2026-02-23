package database_mariadb

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// ReadMariaDB reads a MariaDB database resource.
func (r *MariaDBResource) ReadMariaDB(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DatabaseMariaDBModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	db, err := r.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read MariaDB database, got error: %s", err))
		return
	}

	mapMariaDBResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
