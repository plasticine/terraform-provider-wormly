package wormly

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceWormlySensor() *schema.Resource {
	return &schema.Resource{
		Create: resourceWormlySensorCreate,
		Read:   resourceWormlySensorRead,
		Update: resourceWormlySensorUpdate,
		Delete: resourceWormlySensorDelete,

		Schema: map[string]*schema.Schema{},
	}
}

func resourceWormlySensorCreate(d *schema.ResourceData, meta interface{}) error {
	// TODO
	return nil
}

func resourceWormlySensorRead(d *schema.ResourceData, meta interface{}) error {
	// TODO
	return nil
}

func resourceWormlySensorUpdate(d *schema.ResourceData, meta interface{}) error {
	// TODO
	return nil
}

func resourceWormlySensorDelete(d *schema.ResourceData, meta interface{}) error {
	// TODO
	return nil
}
