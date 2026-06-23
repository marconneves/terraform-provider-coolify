package application_start_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccApplicationStartResource(t *testing.T) {
	appUUID := os.Getenv(tests.ENV_KEY_TEST_APP_UUID)
	if appUUID == "" {
		t.Skipf("Skipping test: %s must be set", tests.ENV_KEY_TEST_APP_UUID)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationStartResourceConfig(appUUID, false, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("coolify_application_start.test", "application_uuid", appUUID),
					resource.TestCheckResourceAttr("coolify_application_start.test", "force", "false"),
					resource.TestCheckResourceAttr("coolify_application_start.test", "instant_deploy", "false"),
					resource.TestCheckResourceAttrSet("coolify_application_start.test", "message"),
					resource.TestCheckResourceAttrSet("coolify_application_start.test", "deployment_uuid"),
				),
			},
			{
				ResourceName:      "coolify_application_start.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccApplicationStartResourceConfig(appUUID, true, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("coolify_application_start.test", "application_uuid", appUUID),
					resource.TestCheckResourceAttr("coolify_application_start.test", "force", "true"),
					resource.TestCheckResourceAttr("coolify_application_start.test", "instant_deploy", "true"),
				),
			},
		},
	})
}

func testAccApplicationStartResourceConfig(appUUID string, force, instantDeploy bool) string {
	return fmt.Sprintf(`
resource "coolify_application_start" "test" {
  application_uuid = %[1]q
  force            = %[2]t
  instant_deploy   = %[3]t
}
`, appUUID, force, instantDeploy)
}
