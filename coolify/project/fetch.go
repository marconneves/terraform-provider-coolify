package project

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func fetchProjectByID(client coolify_sdk.Sdk, ctx context.Context, id string) (*coolify_sdk.Project, error) {
	tflog.Debug(ctx, "Fetching project by ID", map[string]interface{}{
		"project_id": id,
	})

	project, err := client.Project.Get(id)
	if err != nil {
		return nil, fmt.Errorf("unable to read project data by ID, got error: %w", err)
	}

	return project, nil
}

func fetchProjectByName(client coolify_sdk.Sdk, ctx context.Context, name string) (*coolify_sdk.Project, error) {
	tflog.Debug(ctx, "Fetching project by Name", map[string]interface{}{
		"name": name,
	})

	projects, err := client.Project.List()
	if err != nil {
		return nil, fmt.Errorf("unable to list projects, got error: %w", err)
	}

	for _, p := range *projects {
		if p.Name == name {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("no project found with name: %s", name)
}
