package team_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccTeamMembersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTeamMembersDataSourceConfig(3),
				Check:  testAccTeamMembersDataSourceCheck(),
			},
		},
	})
}

func testAccTeamMembersDataSourceConfig(teamID int) string {
	return fmt.Sprintf(`
provider "coolify" {
  address = "%s"
  token   = "%s"
}

data "coolify_team_members" "test" {
  team_id = %d
}
`, os.Getenv("COOLIFY_ADDRESS"), os.Getenv("COOLIFY_TOKEN"), teamID)
}

func testAccTeamMembersDataSourceCheck() resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.coolify_team_members.test", "members.#", "1"),
		resource.TestCheckResourceAttr("data.coolify_team_members.test", "members.0.id", "0"),
		resource.TestCheckResourceAttr("data.coolify_team_members.test", "members.0.name", "Marcon Neves"),
	)
}
