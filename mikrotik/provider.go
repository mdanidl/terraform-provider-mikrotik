package mikrotik

import (
	"fmt"

	"github.com/go-routeros/routeros"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP or Hostname of the Mikrotik device",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username for the Mikrotik device",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password for the user account on the Mikrotik device",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"mikrotik_interface": datasourceMikrotikInterface(),
		},
		ResourcesMap:  map[string]*schema.Resource{},
		ConfigureFunc: providerConfigure,
	}
}
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// config := Config{
	// 	Host:     d.Get("host").(string),
	// 	Ca:       d.Get("username").(string),
	// 	Cert:     d.Get("cert_material").(string),
	// 	Key:      d.Get("key_material").(string),
	// 	CertPath: d.Get("cert_path").(string),
	// }

	client, err := routeros.Dial(
		d.Get("host").(string),
		d.Get("username").(string),
		d.Get("password").(string),
	)
	if err != nil {
		return nil, fmt.Errorf("Error initializing Mikrotik client: %s", err)
	}

	return client, nil
}
