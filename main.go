package main

// https://github.com/spaceapegames/terraform-provider-example/blob/master/provider/resource_item.go
// https://github.com/hashicorp/terraform-provider-hashicups
// https://medium.com/spaceapetech/creating-a-terraform-provider-part-1-ed12884e06d7
// https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc
// Testing
// Implement Test: https://ieftimov.com/posts/testing-in-go-go-test/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-coolify/coolify"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return coolify.Provider()
		},
	})
}
