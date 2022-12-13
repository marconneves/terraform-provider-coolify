package destination_test

import (
	"fmt"
	"regexp"

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


// func TestAccItem_Basic(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() {},
// 		Providers:    TestAccProviders,
// 		CheckDestroy: testAccCheckItemDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckItemBasic(),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckExampleItemExists("coolify_destination.test_item"),
// 					resource.TestCheckResourceAttr(
// 						"coolify_destination.test_item", "name", "my-network"),
// 					resource.TestCheckResourceAttr(
// 						"coolify_destination.test_item", "engine", "/var/run/docker.sock"),
// 					resource.TestCheckResourceAttrSet(
// 						"coolify_destination.test_item", "status.id"),
// 				),
// 			},
// 		},
// 	})
// }

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
resource "coolify_destination" "test_item" {
	name           = "my-network"
	network		   = "checkout"
}
`)
}