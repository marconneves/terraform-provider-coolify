package application_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/api/client"
	"terraform-provider-coolify/coolify"
	"terraform-provider-coolify/shared/tests"
)

var TestAccProviders map[string]*schema.Provider
var TestAccProvider *schema.Provider

func init() {
	TestAccProvider = coolify.Provider()
	TestAccProviders = map[string]*schema.Provider{
		"coolify": TestAccProvider,
	}
}

func GetApplication(client *client.Client, id string) (interface{}, error) {
	return client.GetApplication(id)
}

func TestAccApplication_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: TestAccProviders,
		CheckDestroy: tests.TestAccCheckDestroy(
			TestAccProvider,
			GetApplication,
		),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					tests.TestAccCheckExists(
						"coolify_application.test_item",
						TestAccProvider,
						GetApplication,
					),
					resource.TestCheckResourceAttr(
						"coolify_application.test_item", "name", "first-app"),
					resource.TestCheckResourceAttrSet(
						"coolify_application.test_item", "id"),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
resource "coolify_application" "test_item" {
	name           = "first-app"
	is_bot		   = true

	template {
		build_pack = "node"
		image = "node:14"
		build_image = "node:14"
		
		settings {
			install_command = "npm install"
			start_command = "npm start"
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
		
		env {
			key = "BASE_URL"
			value = ""
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
