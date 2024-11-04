package provider_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	provider "github.com/marconneves/terraform-provider-coolify/coolify"
)

var (
	providerFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"coolify": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)

func checkEnvironmentVariables(t *testing.T) {
	if os.Getenv("TF_ACC") != "1" {
		t.Skip("Skipping acceptance tests unless 'TF_ACC' is set")
	}
	requiredVars := []string{
		"COOLIFY_API_URL",
		"COOLIFY_API_KEY",
	}
	for _, envVar := range requiredVars {
		if os.Getenv(envVar) == "" {
			t.Fatalf("Environment variable `%s` must be set for tests!", envVar)
		}
	}
}

func createDynamicProviderConfig(config map[string]interface{}) (tfprotov6.DynamicValue, error) {
	configTypes := map[string]tftypes.Type{
		"url": tftypes.String,
		"key": tftypes.String,
	}
	configObjectType := tftypes.Object{AttributeTypes: configTypes}
	configObjectValue := tftypes.NewValue(configObjectType, map[string]tftypes.Value{
		"url": tftypes.NewValue(tftypes.String, config["url"]),
		"key": tftypes.NewValue(tftypes.String, config["key"]),
	})
	value, err := tfprotov6.NewDynamicValue(configObjectType, configObjectValue)
	if err != nil {
		return tfprotov6.DynamicValue{}, fmt.Errorf("error creating dynamic value: %w", err)
	}
	return value, nil
}

func TestProviderSchemaVersion(t *testing.T) {
	t.Parallel()
	providerServer, err := providerFactories["coolify"]()
	require.NoError(t, err)
	require.NotNil(t, providerServer)

	resp, err := providerServer.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp.Provider)

	assert.Empty(t, resp.Diagnostics)
	assert.EqualValues(t, 0, resp.Provider.Version)
}

func TestProviderConfiguration(t *testing.T) {
	checkEnvironmentVariables(t)
	apiURL := os.Getenv("COOLIFY_API_URL")
	apiKey := os.Getenv("COOLIFY_API_KEY")

	testCases := map[string]struct {
		config        map[string]interface{}
		envVars       map[string]string
		expectSuccess bool
	}{
		"only url in config": {
			config: map[string]interface{}{
				"url": apiURL,
			},
			expectSuccess: false,
		},
		"only key in config": {
			config: map[string]interface{}{
				"key": apiKey,
			},
			expectSuccess: false,
		},
		"url and key in config": {
			config: map[string]interface{}{
				"url": apiURL,
				"key": apiKey,
			},
			expectSuccess: true,
		},
		"invalid url in config": {
			config: map[string]interface{}{
				"url": "invalid://url",
				"key": apiKey,
			},
			expectSuccess: false,
		},
		"invalid key in config": {
			config: map[string]interface{}{
				"url": apiURL,
				"key": "invalid_key",
			},
			expectSuccess: false,
		},
		"url in env": {
			envVars: map[string]string{
				"COOLIFY_API_URL": apiURL,
			},
			expectSuccess: false,
		},
		"url and key in env": {
			envVars: map[string]string{
				"COOLIFY_API_URL": apiURL,
				"COOLIFY_API_KEY": apiKey,
			},
			expectSuccess: true,
		},
		"url in config, key in env": {
			config: map[string]interface{}{
				"url": apiURL,
			},
			envVars: map[string]string{
				"COOLIFY_API_KEY": apiKey,
			},
			expectSuccess: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Setenv("COOLIFY_API_URL", "")
			t.Setenv("COOLIFY_API_KEY", "")
			for key, value := range tc.envVars {
				t.Setenv(key, value)
			}

			providerServer, err := providerFactories["coolify"]()
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
