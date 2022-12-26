package destination_test

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

func GetDestination(client *client.Client, id string) (interface{}, error) {
	return client.GetDestination(id)
}

func TestAccDestination_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: TestAccProviders,
		CheckDestroy: tests.TestAccCheckDestroy(
			TestAccProvider,
			GetDestination,
		),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					tests.TestAccCheckExists(
						"coolify_destination.test_item",
						TestAccProvider,
						GetDestination,
					),
					resource.TestCheckResourceAttr(
						"coolify_destination.test_item", "name", "my-network"),
					resource.TestCheckResourceAttr(
						"coolify_destination.test_item", "engine", "/var/run/docker.sock"),
					resource.TestCheckResourceAttrSet(
						"coolify_destination.test_item", "id"),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
resource "coolify_destination" "test_item" {
	name           = "my-network"
}
`)
}
