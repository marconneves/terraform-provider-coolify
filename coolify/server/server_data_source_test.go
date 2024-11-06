package server_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccServerDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServerDataSourceConfig(`uuid = "tks4gcsgwowws8gswggkw8gs"`),
				Check:  testAccServerDataSourceCheck(),
			},
			{
				Config: testAccServerDataSourceConfig(`name = "example-server"`),
				Check:  testAccServerDataSourceCheck(),
			},
		},
	})
}

func testAccServerDataSourceConfig(attribute string) string {
	return fmt.Sprintf(`
provider "coolify" {
  address = "%s"
  token   = "%s"
}

data "coolify_server" "test" {
  %s
}
`, os.Getenv("COOLIFY_ADDRESS"), os.Getenv("COOLIFY_TOKEN"), attribute)
}

func testAccServerDataSourceCheck() resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.coolify_server.test", "uuid", "tks4gcsgwowws8gswggkw8gs"),
		resource.TestCheckResourceAttr("data.coolify_server.test", "name", "example-server"),
		resource.TestCheckResourceAttr("data.coolify_server.test", "ip", "localhost"),
		resource.TestCheckResourceAttr("data.coolify_server.test", "high_disk_usage_notification_sent", "false"),
		resource.TestCheckResourceAttr("data.coolify_server.test", "log_drain_notification_sent", "false"),
		resource.TestCheckResourceAttr("data.coolify_server.test", "port", "22"),
	)
}
