package team_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccTeamDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTeamDataSourceConfig("id = 3"),
				Check:  testAccTeamDataSourceCheck(),
			},
			{
				Config: testAccTeamDataSourceConfig(`name = "example-team"`),
				Check:  testAccTeamDataSourceCheck(),
			},
		},
	})
}

func testAccTeamDataSourceConfig(attribute string) string {
	return fmt.Sprintf(`
data "coolify_team" "test" {
  %s
}
`, attribute)
}

func testAccTeamDataSourceCheck() resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.coolify_team.test", "id", "3"),
		resource.TestCheckResourceAttr("data.coolify_team.test", "name", "example-team"),
		resource.TestCheckResourceAttr("data.coolify_team.test", "description", "An example team"),
	)
}
