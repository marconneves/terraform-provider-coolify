package server_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccServerResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServerResourceConfig("example-server-test", "192.168.1.1", "22", "user", "xso0ooc4o0w4cswcwws8gswg"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("coolify_server.test", "name", "example-server-test"),
					resource.TestCheckResourceAttr("coolify_server.test", "ip", "192.168.1.1"),
					resource.TestCheckResourceAttr("coolify_server.test", "port", "22"),
					resource.TestCheckResourceAttr("coolify_server.test", "user", "user"),
					resource.TestCheckResourceAttr("coolify_server.test", "private_key_uuid", "xso0ooc4o0w4cswcwws8gswg"),
				),
			},
			{
				ResourceName:      "coolify_server.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccServerResourceConfig("updated-server", "192.168.1.2", "2222", "newuser", "xso0ooc4o0w4cswcwws8gswg"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("coolify_server.test", "name", "updated-server"),
					resource.TestCheckResourceAttr("coolify_server.test", "ip", "192.168.1.2"),
					resource.TestCheckResourceAttr("coolify_server.test", "port", "2222"),
					resource.TestCheckResourceAttr("coolify_server.test", "user", "newuser"),
					resource.TestCheckResourceAttr("coolify_server.test", "private_key_uuid", "xso0ooc4o0w4cswcwws8gswg"),
				),
			},
		},
	})
}

func testAccServerResourceConfig(name, ip, port, user, privateKeyUUID string) string {
	return fmt.Sprintf(`
    resource "coolify_server" "test" {
      name             = %[1]q
      ip               = %[2]q
      port             = %[3]q
      user             = %[4]q
      private_key_uuid = %[5]q
    }
    `, name, ip, port, user, privateKeyUUID)
}
