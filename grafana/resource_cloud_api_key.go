package grafana

import (
	"context"
	"fmt"
	"strings"

	gapi "github.com/grafana/grafana-api-golang-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var cloudAPIKeyRoles = []string{"Viewer", "Editor", "Admin", "MetricsPublisher", "PluginPublisher"}

func ResourceCloudAPIKey() *schema.Resource {
	return &schema.Resource{
		Description: `Manages a single API key on the Grafana Cloud portal (on the organization level)
* [API documentation](https://grafana.com/docs/grafana-cloud/reference/cloud-api/#api-keys)
`,
		CreateContext: resourceCloudAPIKeyCreate,
		ReadContext:   resourceCloudAPIKeyRead,
		DeleteContext: resourceCloudAPIKeyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"cloud_org_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the organization to create the API key in. This is the same slug as the organization name in the URL.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the API key.",
			},
			"role": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  fmt.Sprintf("Role of the API key. Should be one of %s. See https://grafana.com/docs/grafana-cloud/api/#create-api-key for details.", cloudAPIKeyRoles),
				ValidateFunc: validation.StringInSlice(cloudAPIKeyRoles, false),
			},
			"key": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The generated API key.",
			},
		},
	}
}

func resourceCloudAPIKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client).gcloudapi

	req := &gapi.CreateCloudAPIKeyInput{
		Name: d.Get("name").(string),
		Role: d.Get("role").(string),
	}
	org := d.Get("cloud_org_slug").(string)

	resp, err := c.CreateCloudAPIKey(org, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("key", resp.Token)
	d.SetId(org + "-" + resp.Name)

	return resourceCloudAPIKeyRead(ctx, d, meta)
}

func resourceCloudAPIKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client).gcloudapi

	splitID := strings.SplitN(d.Id(), "-", 2)
	org, name := splitID[0], splitID[1]

	resp, err := c.ListCloudAPIKeys(org)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, apiKey := range resp.Items {
		if apiKey.Name == name {
			d.Set("name", apiKey.Name)
			d.Set("role", apiKey.Role)
			break
		}
	}
	d.Set("cloud_org_slug", org)

	return nil
}

func resourceCloudAPIKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client).gcloudapi

	if err := c.DeleteCloudAPIKey(d.Get("cloud_org_slug").(string), d.Get("name").(string)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
