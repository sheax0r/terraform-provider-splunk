package splunk

import (
	"encoding/xml"
	"fmt"

	resty "gopkg.in/resty.v1"
)

type Dashboard struct {
	Name string
	Data string
}

type dashboardResponse struct {
	Entry xmlEntry `xml:"entry"`
}

type xmlEntry struct {
	Content xmlContent `xml:"content"`
}

type xmlContent struct {
	SDict xmlSDict `xml:"s:dict"`
}

type xmlSDict struct {
	Keys []xmlSKey `xml:"s:key"`
}

type xmlSKey struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

func dashboardCreate(c *resty.Client, d *Dashboard) (r *Dashboard, err error) {
	body := fmt.Sprintf("name=%s&eai:data=%s", d.Name, d.Data)
	_, err = c.R().SetBody([]byte(body)).Post("serviceNS/admin/search/data/ui/views")
	if err != nil {
		return r, err
	}

	return dashboardRead(c, d.Name)

	return r, err
}

func dashboardRead(c *resty.Client, n string) (r *Dashboard, err error) {
	resp, err := c.R().Get(fmt.Sprintf("serviceNS/admin/search/data/ui/views/%s", n))
	if err != nil {
		return r, err
	}

	var dbr dashboardResponse
	xml.Unmarshal(resp.Body(), &dbr)
	keys := dbr.Entry.Content.SDict.Keys
	var dbrData string
	for _, n := range keys {
		if n.Name == "eai:data" {
			dbrData = n.Value
		}
	}

	r = &Dashboard{
		Name: n,
		Data: dbrData,
	}

	return r, err
}
