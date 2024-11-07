package project_environment_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccEnvironmentDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEnvironmentDataSourceConfig(`project_uuid = "ngcskck00wogog4o00o4kgk4"`, `name = "production"`),
				Check:  testAccEnvironmentDataSourceCheck(),
			},
		},
	})
}

func testAccEnvironmentDataSourceConfig(projectUUID, envName string) string {
	return fmt.Sprintf(`
data "coolify_project_environment" "test" {
  %s
  %s
}
`, projectUUID, envName)
}

func testAccEnvironmentDataSourceCheck() resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.coolify_project_environment.test", "name", "production"),
	)
}
