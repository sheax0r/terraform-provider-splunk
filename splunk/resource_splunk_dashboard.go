package splunk

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSplunkDashboard() *schema.Resource {
	return &schema.Resource{
		Create: resourceSplunkDashboardCreate,
		Read:   resourceSplunkDashboardRead,
		Update: resourceSplunkDashboardCreate,
		Delete: resourceSplunkDashboardDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				ForceNew: true,
				Type:     schema.TypeString,
				Required: true,
			},
			"data": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dashboardFromResourceData(d *schema.ResourceData) (r *Dashboard) {
	r = &Dashboard{
		Name: d.Get("name").(string),
		Data: d.Get("data").(string),
	}
	return r
}

func resourceSplunkDashboardCreate(d *schema.ResourceData, meta interface{}) error {
	c, _ := meta.(*Config).RestClient()
	db := dashboardFromResourceData(d)

	log.Printf("[DEBUG] Splunk Dashboard create configuration: %#v", db)

	_, err := dashboardCreate(c, db)
	if err != nil {
		return fmt.Errorf("Failed to create splunk dashboard: %s", err)
	}
	d.SetId(db.Name)

	log.Printf("[INFO] Splunk Dashboard ID: %s", d.Id())
	return resourceSplunkDashboardRead(d, meta)
}

func resourceSplunkDashboardRead(d *schema.ResourceData, meta interface{}) error {
	c, _ := meta.(*Config).RestClient()
	db, err := dashboardRead(c, d.Id())
	if err != nil {
		return err
	}

	d.Set("data", db.Data)

	return nil
}

func resourceSplunkDashboardUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSplunkDashboardDelete(d *schema.ResourceData, meta interface{}) error {
	c, _ := meta.(*Config).RestClient()
	err := dashboardDelete(c, d.Id())
	return err
}
