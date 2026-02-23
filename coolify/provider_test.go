package provider_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createDynamicProviderConfig(config map[string]interface{}) (tfprotov6.DynamicValue, error) {
	configTypes := map[string]tftypes.Type{
		"address": tftypes.String,
		"token":   tftypes.String,
	}
	configObjectType := tftypes.Object{AttributeTypes: configTypes}
	configObjectValue := tftypes.NewValue(configObjectType, map[string]tftypes.Value{
		"address": tftypes.NewValue(tftypes.String, config["address"]),
		"token":   tftypes.NewValue(tftypes.String, config["token"]),
	})
	value, err := tfprotov6.NewDynamicValue(configObjectType, configObjectValue)
	if err != nil {
		return tfprotov6.DynamicValue{}, fmt.Errorf("error creating dynamic value: %w", err)
	}
	return value, nil
}

func TestProviderSchemaVersion(t *testing.T) {
	t.Parallel()
	providerServer, err := tests.ProviderFactories["coolify"]()
	require.NoError(t, err)
	require.NotNil(t, providerServer)

	resp, err := providerServer.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp.Provider)

	assert.Empty(t, resp.Diagnostics)
	assert.EqualValues(t, 0, resp.Provider.Version)
}

func TestProviderConfiguration(t *testing.T) {
	tests.TestAccPreCheck(t)
	apiURL := os.Getenv("COOLIFY_ADDRESS")
	apiKey := os.Getenv("COOLIFY_TOKEN")

	testCases := map[string]struct {
		config        map[string]interface{}
		envVars       map[string]string
		expectSuccess bool
	}{
		"only url in config": {
			config: map[string]interface{}{
				"address": apiURL,
			},
			expectSuccess: false,
		},
		"only key in config": {
			config: map[string]interface{}{
				"token": apiKey,
			},
			expectSuccess: true,
		},
		"url and key in config": {
			config: map[string]interface{}{
				"address": apiURL,
				"token":   apiKey,
			},
			expectSuccess: true,
		},
		"invalid url in config": {
			config: map[string]interface{}{
				"address": "invalid://url",
				"token":   apiKey,
			},
			expectSuccess: true,
		},
		"invalid key in config": {
			config: map[string]interface{}{
				"address": apiURL,
				"token":   "invalid_key",
			},
			expectSuccess: true,
		},
		"url in env": {
			envVars: map[string]string{
				"COOLIFY_ADDRESS": apiURL,
			},
			expectSuccess: false,
		},
		"url and key in env": {
			envVars: map[string]string{
				"COOLIFY_ADDRESS": apiURL,
				"COOLIFY_TOKEN":   apiKey,
			},
			expectSuccess: true,
		},
		"url in config, key in env": {
			config: map[string]interface{}{
				"address": apiURL,
			},
			envVars: map[string]string{
				"COOLIFY_TOKEN": apiKey,
			},
			expectSuccess: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Setenv("COOLIFY_ADDRESS", "")
			t.Setenv("COOLIFY_TOKEN", "")
			for key, value := range tc.envVars {
				t.Setenv(key, value)
			}

			providerServer, err := tests.ProviderFactories["coolify"]()
			require.NoError(t, err)
			require.NotNil(t, providerServer)

			configValue, err := createDynamicProviderConfig(tc.config)
			require.NoError(t, err)
			require.NotNil(t, configValue)

			resp, err := providerServer.ConfigureProvider(context.Background(), &tfprotov6.ConfigureProviderRequest{
				Config: &configValue,
			})
			require.NoError(t, err)
			require.NotNil(t, resp)

			if tc.expectSuccess {
				assert.Empty(t, resp.Diagnostics)
			} else {
				assert.NotEmpty(t, resp.Diagnostics)
			}
		})
	}
}

const ResourceNamePrefix = "test-resource"

func GenerateRandomResourceName(resourceType string) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	name := make([]rune, 8)
	for i := range name {
		name[i] = letters[rand.Intn(len(letters))]
	}
	return fmt.Sprintf("%s-%s-%s", ResourceNamePrefix, resourceType, string(name))
}
