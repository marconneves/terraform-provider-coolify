package database_mysql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// ReadMySQL reads a MySQL database resource.
func (r *MySQLResource) ReadMySQL(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DatabaseMySQLModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	db, err := r.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read MySQL database, got error: %s", err))
		return
	}

	mapMySQLResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
