package tests

import (
	"fmt"
	"regexp"

	sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tf "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func CheckAttribute(resourceOrDataSource string, key string, value string) resource.TestCheckFunc {
	return func(state *tf.State) error {
		rs, ok := state.RootModule().Resources[resourceOrDataSource]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceOrDataSource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		if rs.Primary.Attributes[key] != value {
			return fmt.Errorf(key + "is not set")
		}

		return nil
	}
}

func TestAccCheckDestroy(TestAccProvider *schema.Provider, Verify func(client *sdk.Client, id string) (interface{}, error)) resource.TestCheckFunc {
	return func(s *tf.State) error {
		apiClient := TestAccProvider.Meta().(*sdk.Client)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "example_item" {
				continue
			}

			_, err := Verify(apiClient, rs.Primary.ID)
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
}

func TestAccCheckExists(resourceOrDataSource string, TestAccProvider *schema.Provider, Verify func(client *sdk.Client, id string) (interface{}, error)) resource.TestCheckFunc {
	return func(state *tf.State) error {
		rs, ok := state.RootModule().Resources[resourceOrDataSource]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceOrDataSource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		apiClient := TestAccProvider.Meta().(*sdk.Client)
		_, err := Verify(apiClient, rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", resourceOrDataSource, err)
		}

		return nil
	}
}
