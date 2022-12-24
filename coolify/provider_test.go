package coolify_test

import (
	"os"
	"testing"

	"terraform-provider-coolify/coolify"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var TestAccProviders map[string]*schema.Provider
var TestAccProvider *schema.Provider

func init() {
	TestAccProvider = coolify.Provider()
	TestAccProviders = map[string]*schema.Provider{
		"coolify": TestAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := coolify.Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = coolify.Provider()
}

func TestAccPreCheck(t *testing.T) {
	if v := os.Getenv("TF_ACC"); v == "" {
		t.Fatal("TF_ACC must be set for acceptance tests")
	}
	if v := os.Getenv("SERVICE_ADDRESS"); v == "" {
		t.Fatal("SERVICE_ADDRESS must be set for acceptance tests")
	}
	if v := os.Getenv("SERVICE_ADDRESS"); v == "" {
		t.Fatal("SERVICE_ADDRESS must be set for acceptance tests")
	}
}
