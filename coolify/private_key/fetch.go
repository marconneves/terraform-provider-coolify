package private_key

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func FetchPrivateKeyByID(client coolify_sdk.Sdk, ctx context.Context, id int) (*coolify_sdk.PrivateKey, error) {
	tflog.Debug(ctx, "Fetching project by Id", map[string]interface{}{
		"id": id,
	})

	projects, err := client.PrivateKey.List()
	if err != nil {
		return nil, fmt.Errorf("unable to list projects, got error: %w", err)
	}

	for _, p := range *projects {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("no project found with id: %v", id)
}

func fetchPrivateKeyByUUID(client coolify_sdk.Sdk, ctx context.Context, id string) (*coolify_sdk.PrivateKey, error) {
	tflog.Debug(ctx, "Fetching private key by UUID", map[string]interface{}{
		"private_key_id": id,
	})

	privateKey, err := client.PrivateKey.Get(id)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key data by UUID, got error: %w", err)
	}

	return privateKey, nil
}
