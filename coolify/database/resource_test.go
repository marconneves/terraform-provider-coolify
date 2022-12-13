package database_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tf "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"terraform-provider-coolify/api/client"
	"terraform-provider-coolify/coolify"
)


var TestAccProviders map[string]*schema.Provider
var TestAccProvider *schema.Provider

func init() {
	TestAccProvider = coolify.Provider()
	TestAccProviders = map[string]*schema.Provider{
		"coolify": TestAccProvider,
	}
}


func TestAccDatabase_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() {},
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleItemExists("coolify_database.test_item"),
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

func testAccCheckItemDestroy(s *tf.State) error {
	apiClient := TestAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "example_item" {
			continue
		}

		_, err := apiClient.GetDatabase(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Alert still exists")
		}
		notFoundErr := "not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}

	return nil
}

func testAccCheckExampleItemExists(resource string) resource.TestCheckFunc {
	return func(state *tf.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		id := rs.Primary.ID
		apiClient := TestAccProvider.Meta().(*client.Client)
		_, err := apiClient.GetDatabase(id)
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", resource, err)
		}
		return nil
	}
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