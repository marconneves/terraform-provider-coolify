package private_key

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func fetchPrivateKeyByID(client coolify_sdk.Sdk, ctx context.Context, id string) (*coolify_sdk.PrivateKey, error) {
	tflog.Debug(ctx, "Fetching private key by ID", map[string]interface{}{
		"private_key_id": id,
	})

	privateKey, err := client.PrivateKey.Get(id)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key data by ID, got error: %w", err)
	}

	return privateKey, nil
}
