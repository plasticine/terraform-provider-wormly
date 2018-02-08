import (
  "fmt"

  "github.com/hashicorp/terraform/helper/schema"
)

func resourceWormlyHost() *schema.Resource {
  return &schema.Resource{
    Create: resourceWormlyHostCreate,
    Read:   resourceWormlyHostRead,
    Update: resourceWormlyHostUpdate,
    Delete: resourceWormlyHostDelete,

    Schema: map[string]*schema.Schema{}
  }
}

func resourceWormlyHostCreate(d *schema.ResourceData, meta interface{}) error {
  // TODO
  return nil
}

func resourceWormlyHostRead(d *schema.ResourceData, meta interface{}) error {
  // TODO
  return nil
}

func resourceWormlyHostUpdate(d *schema.ResourceData, meta interface{}) error {
  // TODO
  return nil
}

func resourceWormlyHostDelete(d *schema.ResourceData, meta interface{}) error {
  // TODO
  return nil
}
