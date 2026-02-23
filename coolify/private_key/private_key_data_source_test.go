package private_key_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/marconneves/terraform-provider-coolify/shared/tests"
)

func TestAccPrivateKeyDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { tests.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: tests.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPrivateKeyDataSourceConfig(`uuid = "gsssowsgoswo4g8sswsc0kc4"`),
				Check:  testAccPrivateKeyDataSourceCheck(),
			},
		},
	})
}

func testAccPrivateKeyDataSourceConfig(attribute string) string {
	return fmt.Sprintf(`
data "coolify_private_key" "test" {
  %s
}
`, attribute)
}

func testAccPrivateKeyDataSourceCheck() resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.coolify_private_key.test", "uuid", "gsssowsgoswo4g8sswsc0kc4"),
		resource.TestCheckResourceAttr("data.coolify_private_key.test", "name", "example-test"),
		resource.TestCheckResourceAttr("data.coolify_private_key.test", "description", "A test private key"),
		resource.TestCheckResourceAttr("data.coolify_private_key.test", "is_git_related", "true"),
	)
}
