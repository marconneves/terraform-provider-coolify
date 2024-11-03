package database_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/coolify"
	"terraform-provider-coolify/shared/tests"

	sdk "github.com/marconneves/coolify-sdk-go"
)

var TestAccProviders map[string]*schema.Provider
var TestAccProvider *schema.Provider

func init() {
	TestAccProvider = coolify.Provider()
	TestAccProviders = map[string]*schema.Provider{
		"coolify": TestAccProvider,
	}
}

func GetDatabase(client *sdk.Client, id string) (interface{}, error) {
	return client.GetDatabase(id)
}

func TestAccDatabase_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: TestAccProviders,
		CheckDestroy: tests.TestAccCheckDestroy(
			TestAccProvider,
			GetDatabase,
		),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					tests.TestAccCheckExists(
						"coolify_database.test_item",
						TestAccProvider,
						GetDatabase,
					),
					resource.TestCheckResourceAttr(
						"coolify_database.test_item", "name", "my-db"),
					resource.TestCheckResourceAttr(
						"coolify_database.test_item", "engine", "postgresql:13.8.0"),
					resource.TestCheckResourceAttrSet(
						"coolify_database.test_item", "status.port"),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
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
`)
}
