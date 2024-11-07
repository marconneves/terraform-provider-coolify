package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccProjectResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectResourceConfig("example-project-test", "An example project"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("coolify_project.test", "name", "example-project-test"),
					resource.TestCheckResourceAttr("coolify_project.test", "description", "An example project"),
				),
			},
			{
				ResourceName:      "coolify_project.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccProjectResourceConfig("updated-project", "An updated project description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("coolify_project.test", "name", "updated-project"),
					resource.TestCheckResourceAttr("coolify_project.test", "description", "An updated project description"),
				),
			},
		},
	})
}

func testAccProjectResourceConfig(name, description string) string {
	if description == "" {
		description = "Default description"
	}
	return fmt.Sprintf(`
    resource "coolify_project" "test" {
      name        = %[1]q
      description = %[2]q
    }
    `, name, description)
}
