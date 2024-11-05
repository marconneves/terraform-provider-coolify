package project_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccProjectDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectDataSourceConfig(`id = "ngcskck00wogog4o00o4kgk4"`),
				Check:  testAccProjectDataSourceCheck(),
			},
			{
				Config: testAccProjectDataSourceConfig(`name = "example-project"`),
				Check:  testAccProjectDataSourceCheck(),
			},
		},
	})
}

func testAccProjectDataSourceConfig(attribute string) string {
	return fmt.Sprintf(`
provider "coolify" {
  address = "%s"
  token   = "%s"
}

data "coolify_project" "test" {
  %s
}
`, os.Getenv("COOLIFY_ADDRESS"), os.Getenv("COOLIFY_TOKEN"), attribute)
}

func testAccProjectDataSourceCheck() resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.coolify_project.test", "id", "ngcskck00wogog4o00o4kgk4"),
		resource.TestCheckResourceAttr("data.coolify_project.test", "name", "example-project"),
		resource.TestCheckResourceAttr("data.coolify_project.test", "description", ""),
	)
}
