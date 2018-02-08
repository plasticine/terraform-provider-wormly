package wormly

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceWormlyHost() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceWormlyHostRead,
		Schema: map[string]*schema.Schema{},
	}
}

func dataSourceWormlyHostRead(d *schema.ResourceData, meta interface{}) error {

}
