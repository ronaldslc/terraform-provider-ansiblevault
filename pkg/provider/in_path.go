package provider

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronaldslc/terraform-provider-ansiblevault/v2/pkg/vault"
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
				Optional:    true,
			},
			"value": {
				Computed:    true,
				Description: "Vault value found",
				Type:        schema.TypeString,
				Sensitive:   true,
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
