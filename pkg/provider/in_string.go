package provider

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ronaldslc/terraform-provider-ansiblevault/v2/pkg/vault"
)

func inStringResource() *schema.Resource {
	return &schema.Resource{
		Read: inStringRead,
		Schema: map[string]*schema.Schema{
			"encrypted": {
				Type:        schema.TypeString,
				Description: "Ansible-vault string representation",
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

func inStringRead(data *schema.ResourceData, m interface{}) error {
	raw := data.Get("encrypted").(string)
	key := data.Get("key").(string)

	data.SetId(time.Now().UTC().String())

	value, err := m.(*vault.App).InString(raw, key)

	if err != nil {
		data.SetId("")
		return err
	}

	if err := data.Set("value", value); err != nil {
		data.SetId("")
		return err
	}

	return nil
}
