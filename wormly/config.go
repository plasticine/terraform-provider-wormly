package wormly

import (
	"fmt"
	"log"
	"runtime"

	"github.com/hashicorp/terraform/terraform"
)

type Config struct {
	api_key string

	client    *WormlyClient
	userAgent string
}

func (c *Config) loadAndValidate() error {
	log.Printf("[INFO] Configuring Wormly Provider")

	client := NewWormlyClient()

	versionString := terraform.VersionString()
	userAgent := fmt.Sprintf("(%s %s) Terraform/%s", runtime.GOOS, runtime.GOARCH, versionString)

	c.client = client
	c.userAgent = userAgent

	return nil
}
