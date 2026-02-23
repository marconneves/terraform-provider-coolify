package database_redis

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// ReadRedis reads a Redis database resource.
func (r *RedisResource) ReadRedis(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DatabaseRedisModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	db, err := r.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Redis database, got error: %s", err))
		return
	}

	mapRedisResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
