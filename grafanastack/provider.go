package grafanastack

import (
	"context"

	smapi "github.com/grafana/synthetic-monitoring-api-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"url": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("GRAFANA_URL", ""),
				},
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("GRAFANA_STACK_NAME", ""),
				},
				"accesskey": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("GRAFANA_ACCESS_KEY", ""),
				},
			},

			ResourcesMap: map[string]*schema.Resource{
				// Grafana

			},
		}

		p.ConfigureContextFunc = configure(version, p)
		return p
	}
}

type client struct {
	smapi *smapi.Client
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics
		c := &client{}
		address := d.Get("url").(string)
		token := d.Get("accesskey").(string)
		c.smapi = smapi.NewClient(address, token, nil)
		return c, diags
	}
}
