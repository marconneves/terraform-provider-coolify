package application_restart_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccApplicationRestartResource(t *testing.T) {
	appUUID := os.Getenv(tests.ENV_KEY_TEST_APP_UUID)
	if appUUID == "" {
		t.Skipf("Skipping test: %s must be set", tests.ENV_KEY_TEST_APP_UUID)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationRestartResourceConfig(appUUID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("coolify_application_restart.test", "application_uuid", appUUID),
					resource.TestCheckResourceAttrSet("coolify_application_restart.test", "message"),
					resource.TestCheckResourceAttrSet("coolify_application_restart.test", "deployment_uuid"),
				),
			},
			{
				ResourceName:      "coolify_application_restart.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccApplicationRestartResourceConfig(appUUID string) string {
	return fmt.Sprintf(`
resource "coolify_application_restart" "test" {
  application_uuid = %[1]q
}
`, appUUID)
}
