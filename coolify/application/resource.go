package application

import (
	"context"
	"terraform-provider-coolify/shared"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Status struct {
	Domain string
}

type Env struct {
	Key        string
	Value      string
	IsBuildEnv bool
}

type Application struct {
	Name   string
	Domain string
	IsBot  bool

	Template struct {
		BuildPack  string
		Image      string
		BuildImage string

		Settings struct {
			IsCoolifyBuildPack bool
			AutoDeploy         bool
			InstallCommand     string
			BuildCommand       string
			StartCommand       string
		}

		Envs []Env
	}

	Repository struct {
		RepositoryId int
		Repository   string
		Branch       string
		commitHash   string
	}

	Settings struct {
		SourceId      string
		DestinationId string
	}

	Status Status
}

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: applicationCreateItem,
		ReadContext:   applicationReadItem,
		UpdateContext: applicationUpdateItem,
		DeleteContext: applicationDeleteItem,
		Exists:        applicationExistsItem,
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
			"domain": {
				Type:        schema.TypeString,
				Description: "Domain for the application.",
				Optional:    true,
				Default:     "",
			},
			"is_bot": {
				Type:        schema.TypeBool,
				Description: "Is the application a bot.",
				Optional:    true,
				Default:     false,
			},

			"template": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"build_pack": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"image": {
							Type:     schema.TypeString,
							Required: true,
						},
						"build_image": {
							Type:     schema.TypeString,
							Required: true,
						},
						"settings": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"install_command": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"build_command": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"start_command": {
										Type:     schema.TypeString,
										Required: true,
									},
									"is_coolify_build_pack": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"auto_deploy": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
								},
							},
						},
						"env": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
									"is_build_env": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
								},
							},
						},
					},
				},
			},

			"repository": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository_id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"repository": {
							Type:     schema.TypeString,
							Required: true,
						},
						"branch": {
							Type:     schema.TypeString,
							Required: true,
						},
						"commit_hash": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"settings": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"destination_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			"status": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
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
