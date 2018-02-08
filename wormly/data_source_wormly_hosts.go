package wormly

import (
	"github.com/hashicorp/terraform/helper/schema"
  "log"
  "fmt"
)

func dataSourceWormlyHosts() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceWormlyHostsRead,
		Schema: map[string]*schema.Schema{
      "hosts": &schema.Schema{
        Type:     schema.TypeList,
        Computed: true,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
            "id": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
          },
        },
      },
    },
	}
}

func dataSourceWormlyHostsRead(d *schema.ResourceData, meta interface{}) error {
  log.Printf("[INFO] ---------------------------------------------------------")

  hosts := parseHosts([]string{"Foo", "Bar"})
  log.Printf(fmt.Sprintf("%v", hosts))

  d.Set("hosts", hosts)

  return nil
}

func parseHosts(hosts []string) []map[string]interface{} {
  hostsSchema := make([]map[string]interface{}, 0, len(hosts))
  for _, host := range hosts {
    data := map[string]interface{}{
      "id":    host,
    }
    hostsSchema = append(hostsSchema, data)
  }
  return hostsSchema
}
