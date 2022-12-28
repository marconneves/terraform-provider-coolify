package database_test

import (
	"testing"

	"terraform-provider-coolify/shared/tests"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatabase_DataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: TestAccProviders,
		CheckDestroy: tests.TestAccCheckDestroy(
			TestAccProvider,
			GetDatabase,
		),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSpannerInstanceBasic(),
				Check: resource.ComposeTestCheckFunc(
					tests.TestAccCheckExists(
						"data.coolify_database.postgress",
						TestAccProvider,
						GetDatabase,
					),
					tests.CheckAttribute("data.coolify_database.postgress", "name", "my-db"),
				),
			},
		},
	})
}

func testAccDataSourceSpannerInstanceBasic() string {
	return (`
resource "coolify_database" "test_item" {
	name           = "my-db"
	engine         = "postgresql:13.8.0"
  
	settings {
	  destination_id = "clb9wrx87001fmo9dvvog6xet"
	  is_public	  	 = true
	  default_database = "postgres"
	  user = "myuser"
	  password = "mypassword"
	  root_password = "rootpassword"
	}
}

data "coolify_database" "postgress" {
	database_id = coolify_database.test_item.id
}
`)
}
