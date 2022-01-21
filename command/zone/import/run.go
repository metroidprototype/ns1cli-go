package imp

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/metroidprototype/ns1cli-go/command/zone/helper"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

func (c *cmd) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("zone import only accepts 1 argument")
		c.UI.Info(c.Help())
		return 1
	}
	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("zonefile", fi.Name())
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	io.Copy(part, file)
	writer.Close()

	path := fmt.Sprintf("import/zonefile/%s", fi.Name())
	req, err := c.newImportRequest("PUT", path, body)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	zone := &dns.Zone{}
	_, err = c.ns1.Do(req, &zone)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Info("Zonefile Imported.")
	if zone != nil {
		c.UI.Info(helper.FormatZone(zone))
	}
	return 0
}

// newImportRequest is a custom request function to properly write the
// file as form data for import calls
func (c *cmd) newImportRequest(method, path string, body *bytes.Buffer) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	uri := c.ns1.Endpoint.ResolveReference(rel)
	req, err := http.NewRequest(method, uri.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-NSONE-Key", c.ns1.APIKey)
	req.Header.Add("User-Agent", c.ns1.UserAgent)
	return req, nil
}
