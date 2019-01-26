package mikrotik

import (
	"log"

	"github.com/go-routeros/routeros"
	"github.com/hashicorp/terraform/helper/schema"
)

func datasourceMikrotikInterface() *schema.Resource {
	return &schema.Resource{
		Read: datasourceMikrotikInterfaceRead,
		Schema: map[string]*schema.Schema{

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "name of the interface",
			},
			// computed attributes
			"mac_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "MAC Address of the interface",
			},
			"id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "id of the interface",
			},
			"disabled": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "whether or not the interface is disabled",
			},
		},
	}
}

func datasourceMikrotikInterfaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*routeros.Client)

	mondatok := []string{
		"/interface/print",
		"?name=" + d.Get("name").(string),
	}
	// mondatok = append(mondatok, "?type=bridge", "=.proplist=.id,name,disabled")

	r, err := client.RunArgs(mondatok)
	if err != nil {
		log.Fatal(err)
	}
	responseData := r.Re[0]
	responseMap := responseData.Map
	d.SetId(responseMap[".id"][1:])
	d.Set("name", responseMap["name"])
	d.Set("id", responseMap[".id"])
	d.Set("mac_address", responseMap["mac-address"])
	d.Set("disabled", responseMap["disabled"])

	return nil
}
