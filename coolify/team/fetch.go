package team

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func fetchTeamByID(client coolify_sdk.Sdk, ctx context.Context, id int64) (*coolify_sdk.Team, error) {
	tflog.Debug(ctx, "Fetching team by ID", map[string]interface{}{
		"team_id": id,
	})

	return client.Team.Get(int(id))
}

func fetchTeamByName(client coolify_sdk.Sdk, ctx context.Context, name string) (*coolify_sdk.Team, error) {
	tflog.Debug(ctx, "Fetching team by Name", map[string]interface{}{
		"name": name,
	})

	teams, err := client.Team.List()
	if err != nil {
		return nil, err
	}

	for _, t := range *teams {
		if t.Name == name {
			return &t, nil
		}
	}

	return nil, nil
}
