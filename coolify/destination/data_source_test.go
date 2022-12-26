package destination_test

import (
	"testing"

	"terraform-provider-coolify/shared/tests"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDestination_DataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: TestAccProviders,
		CheckDestroy: tests.TestAccCheckDestroy(
			TestAccProvider,
			GetDestination,
		),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSpannerInstanceBasic(),
				Check: resource.ComposeTestCheckFunc(
					tests.TestAccCheckExists(
						"data.coolify_destination.network",
						TestAccProvider,
						GetDestination,
					),
					tests.CheckAttribute("data.coolify_destination.network", "name", "my-network"),
				),
			},
		},
	})
}

func testAccDataSourceSpannerInstanceBasic() string {
	return (`
resource "coolify_destination" "test_item" {
	name           = "my-network"
}

data "coolify_destination" "network" {
	network = coolify_destination.test_item.network
}
`)
}
