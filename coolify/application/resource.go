package application

import (
	"context"
	"terraform-provider-coolify/shared"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationCreateItem,
		ReadContext: applicationReadItem,
		UpdateContext: applicationUpdateItem,
		DeleteContext: applicationDeleteItem,
		Exists: applicationExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Description:  "Name of the application.",
				Required:     true,
				ForceNew:     false,
				ValidateFunc: shared.ValidateName,
			},
		},
	}
}


func applicationReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}


func applicationUpdateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}


func applicationDeleteItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func applicationExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	return true, nil
}