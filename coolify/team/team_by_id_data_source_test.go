package team_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccTeamByIDDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTeamByIDDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.coolify_team_by_id.test", "id", "123"),
					resource.TestCheckResourceAttr("data.coolify_team_by_id.test", "name", "example-team"),
					resource.TestCheckResourceAttr("data.coolify_team_by_id.test", "description", "An example team"),
				),
			},
		},
	})
}

const testAccTeamByIDDataSourceConfig = `
data "coolify_team_by_id" "test" {
  id = 123
}
`
