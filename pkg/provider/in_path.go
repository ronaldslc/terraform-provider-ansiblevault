package provider

import (
	"fmt"
	"time"

	"github.com/MeilleursAgents/terraform-provider-ansiblevault/pkg/vault"
	"github.com/hashicorp/terraform/helper/schema"
)

func inPathResource() *schema.Resource {
	return &schema.Resource{
		Read: inPathRead,
		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Description: "Ansible environment searched",
				Required:    true,
			},
			"key": {
				Type:        schema.TypeString,
				Description: "Vault key searched",
				Required:    true,
			},
			"value": {
				Computed:    true,
				Description: "Vault value found",
				Type:        schema.TypeString,
			},
		},
	}
}

func inPathRead(data *schema.ResourceData, m interface{}) error {
	path := data.Get("path").(string)
	key := data.Get("key").(string)

	data.SetId(time.Now().UTC().String())

	value, err := m.(*vault.App).InPath(path, key)
	if err != nil {
		data.SetId("")

		if err == vault.ErrKeyNotFound {
			return fmt.Errorf("%s not found in %s vault", key, path)
		}

		return err
	}

	if err := data.Set("value", value); err != nil {
		data.SetId("")
		return err
	}

	return nil
}
