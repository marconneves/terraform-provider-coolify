package destination_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDestination_DataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() {},
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSpannerInstanceBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleItemExists("data.coolify_destination.network"),
					checkAttribute("data.coolify_destination.network", "name", "my-network"),
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
