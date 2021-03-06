package wormly

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Global MutexKV
var mutexKV = mutexkv.NewMutexKV()

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"API_KEY",
					"WORMLY_API_KEY",
				}, nil),
			},
		},

		ConfigureFunc: providerConfigure,
		DataSourcesMap: map[string]*schema.Resource{
			"wormly_hosts": dataSourceWormlyHosts(),
		},
		ResourcesMap: map[string]*schema.Resource{
			// "wormly_host":   resourceWormlyHost(),
			// "wormly_sensor": resourceWormlySensor(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	api_key := d.Get("api_key").(string)
	config := Config{
		api_key: api_key,
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
