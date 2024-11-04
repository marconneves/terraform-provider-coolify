package tests

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	provider "github.com/marconneves/terraform-provider-coolify/coolify"
)

const (
	ENV_KEY_ADDRESS = "COOLIFY_ADDRESS"
	ENV_KEY_TOKEN   = "COOLIFY_TOKEN"
)

func TestAccPreCheck(t *testing.T) {
	if os.Getenv("TF_ACC") != "1" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	variables := []string{
		ENV_KEY_ADDRESS,
		ENV_KEY_TOKEN,
	}

	for _, variable := range variables {
		value := os.Getenv(variable)
		if value == "" {
			t.Fatalf("`%s` must be set for acceptance tests!", variable)
		}
	}
}

var (
	ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"coolify": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)
