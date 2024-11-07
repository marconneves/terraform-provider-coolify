package team_members_test

import (
	"fmt"
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
data "coolify_team_members" "test" {
  team_id = %d
}
`, teamID)
}

func testAccTeamMembersDataSourceCheck() resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.coolify_team_members.test", "members.#", "1"),
		resource.TestCheckResourceAttr("data.coolify_team_members.test", "members.0.id", "0"),
		resource.TestCheckResourceAttr("data.coolify_team_members.test", "members.0.name", "Marcon Neves"),
	)
}
