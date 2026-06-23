package application_stop_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccApplicationStopResource(t *testing.T) {
	appUUID := os.Getenv(tests.ENV_KEY_TEST_APP_UUID)
	if appUUID == "" {
		t.Skipf("Skipping test: %s must be set", tests.ENV_KEY_TEST_APP_UUID)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationStopResourceConfig(appUUID, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("coolify_application_stop.test", "application_uuid", appUUID),
					resource.TestCheckResourceAttr("coolify_application_stop.test", "docker_cleanup", "false"),
				),
			},
			{
				ResourceName:      "coolify_application_stop.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccApplicationStopResourceConfig(appUUID, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("coolify_application_stop.test", "application_uuid", appUUID),
					resource.TestCheckResourceAttr("coolify_application_stop.test", "docker_cleanup", "true"),
				),
			},
		},
	})
}

func testAccApplicationStopResourceConfig(appUUID string, dockerCleanup bool) string {
	return fmt.Sprintf(`
resource "coolify_application_stop" "test" {
  application_uuid = %[1]q
  docker_cleanup   = %[2]t
}
`, appUUID, dockerCleanup)
}
