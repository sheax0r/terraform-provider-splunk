package splunk

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/denniswebb/go-splunk/splunk"
	resty "gopkg.in/resty.v1"
)

type Config struct {
	URL                string
	Username           string
	Password           string
	InsecureSkipVerify bool
}

// Client() returns a new client for accessing Splunk.
func (c *Config) Client() (*splunk.Client, error) {
	client := splunk.New(c.URL, c.Username, c.Password, c.InsecureSkipVerify)
	log.Printf("[INFO] Splunk Client configured for: %s@%s", c.Username, c.URL)
	return client, nil
}

// RestClient() returns a low-level REST client for accessing Splunk.
func (c *Config) RestClient() (*resty.Client, error) {
	client := resty.New().
		SetBasicAuth(c.Username, c.Password).
		SetHostURL(fmt.Sprintf("%s", c.URL)).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetQueryParam("output_mode", "json").
		SetMode("rest").
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: c.InsecureSkipVerify})
	log.Printf("[INFO] REST Client configured for: %s@%s", c.Username, c.URL)
	return client, nil
}
