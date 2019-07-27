package splunk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	resty "gopkg.in/resty.v1"
)

type Dashboard struct {
	Name string
	Data string
}

type dashboardResponse struct {
	Entries []struct {
		Content struct {
			Data string `json:"eai:data"`
		} `json:"content"`
	} `json:"entry"`
}

func dashboardCreate(c *resty.Client, d *Dashboard) (r *Dashboard, err error) {
	body := fmt.Sprintf("name=%s&eai:data=%s", url.QueryEscape(d.Name), d.Data)
	resp, err := c.R().SetBody([]byte(body)).Post(fmt.Sprintf("/servicesNS/%s/search/data/ui/views", c.UserInfo.Username))
	if err != nil {
		return r, err
	}

	log.Printf("[DEBUG] response: %+v", resp)

	return dashboardRead(c, d.Name)

	return r, err
}

func dashboardDelete(c *resty.Client, n string) (err error) {
	resp, err := c.R().Delete(fmt.Sprintf("servicesNS/%s/search/data/ui/views/%s", c.UserInfo.Username, n))
	log.Printf("[DEBUG] response: %+v", resp)
	return err
}

func dashboardRead(c *resty.Client, n string) (r *Dashboard, err error) {
	resp, err := c.R().Get(fmt.Sprintf("servicesNS/%s/search/data/ui/views/%s", c.UserInfo.Username, n))
	if err != nil {
		return r, err
	}
	log.Printf("[DEBUG] response: %+v", resp)
	log.Printf("[DEBUG] response body: %+v", string(resp.Body()))

	var dbr dashboardResponse
	json.Unmarshal(resp.Body(), &dbr)

	r = &Dashboard{
		Name: n,
		Data: dbr.Entries[0].Content.Data,
	}

	return r, err
}
