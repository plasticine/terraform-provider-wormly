package wormly

import (
	"fmt"
	"log"
)

type Config struct {
	api_key string

	client    *WormlyClient
	userAgent string
}

func (c *Config) loadAndValidate() error {
	client := NewWormlyClient()

	versionString := terraform.VersionString()
	userAgent := fmt.Sprintf("(%s %s) Terraform/%s", runtime.GOOS, runtime.GOARCH, versionString)

	c.client = client
	c.userAgent = userAgent

	log.Printf("[INFO] HELLO WORLD")

	return nil
}
