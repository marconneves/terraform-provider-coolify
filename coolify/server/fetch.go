package server

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func fetchServerByID(client coolify_sdk.Sdk, ctx context.Context, id string) (*coolify_sdk.Server, error) {
	tflog.Debug(ctx, "Fetching server by ID", map[string]interface{}{
		"server_id": id,
	})

	server, err := client.Server.Get(id)
	if err != nil {
		return nil, fmt.Errorf("unable to read server data by ID, got error: %w", err)
	}

	return server, nil
}

func fetchServerByName(client coolify_sdk.Sdk, ctx context.Context, name string) (*coolify_sdk.Server, error) {
	tflog.Debug(ctx, "Fetching server by Name", map[string]interface{}{
		"name": name,
	})

	servers, err := client.Server.List()
	if err != nil {
		return nil, fmt.Errorf("unable to list servers, got error: %w", err)
	}

	for _, s := range *servers {
		if s.Name == name {
			return &s, nil
		}
	}

	return nil, fmt.Errorf("no server found with name: %s", name)
}
