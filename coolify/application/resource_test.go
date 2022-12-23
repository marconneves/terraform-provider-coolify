package application_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tf "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"terraform-provider-coolify/api/client"
	"terraform-provider-coolify/coolify"
)

var TestAccProviders map[string]*schema.Provider
var TestAccProvider *schema.Provider

func init() {
	TestAccProvider = coolify.Provider()
	TestAccProviders = map[string]*schema.Provider{
		"coolify": TestAccProvider,
	}
}

func TestAccApplication_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() {},
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleItemExists("coolify_application.test_item"),
					resource.TestCheckResourceAttr(
						"coolify_application.test_item", "name", "first-app"),
					resource.TestCheckResourceAttrSet(
						"coolify_application.test_item", "id"),
				),
			},
		},
	})
}

func testAccCheckItemDestroy(s *tf.State) error {
	apiClient := TestAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "example_item" {
			continue
		}

		_, err := apiClient.GetApplication(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Alert still exists")
		}
		notFoundErr := "not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}

	return nil
}

func testAccCheckExampleItemExists(resource string) resource.TestCheckFunc {
	return func(state *tf.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		id := rs.Primary.ID
		apiClient := TestAccProvider.Meta().(*client.Client)
		_, err := apiClient.GetApplication(id)
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", resource, err)
		}
		return nil
	}
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
resource "coolify_application" "test_item" {
	name           = "first-app"
	domain		   = "first-app.s.b4.run"

	template {
		build_pack = "node"
		image = "node:14"
		build_image = "node:14"
		
		settings {
			install_command = "npm install"
			start_command = "npm start"
			is_coolify_build_pack = true
			auto_deploy = false
		}

		env {
			key = "BASE_PROJECT"
			value = "production"
		}

		env {
			key = "BASE_URL"
			value = "https://front.s.b4.run"
		}
	}

	repository {
		repository_id = 579493141
		repository = "cool-sample/sample-nodejs"
		branch = "main"
	}
	
	settings {
		destination_id = "clb9wrx87001fmo9dvvog6xet"
		source_id = "clb9y09gs000f9dmod69f7dce"
	}
}
`)
}
