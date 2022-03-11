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

	flags "github.com/uber-go/flagoverride"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

func (c *cmd) Run(args []string) int {
	flags.ParseArgs(&c.Flags, args)
	if len(c.Flags.Zone) == 0 {
		c.Ui.Error("zone option required")
		c.Ui.Info(c.Help())
		return 1
	}
	file, err := os.Open(c.Flags.Zone)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("zonefile", fi.Name())
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	io.Copy(part, file)
	writer.Close()

	path := fmt.Sprintf("import/zonefile/%s", fi.Name())
	req, err := c.newImportRequest("PUT", path, body)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	zone := &dns.Zone{}
	_, err = c.Ns1.Do(req, &zone)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	c.Ui.Info("Zonefile Imported.")
	if zone != nil {
		c.Ui.Info(helper.FormatZone(&c.Cmd, zone))
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

	uri := c.Ns1.Endpoint.ResolveReference(rel)
	req, err := http.NewRequest(method, uri.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-NSONE-Key", c.Ns1.APIKey)
	req.Header.Add("User-Agent", c.Ns1.UserAgent)
	return req, nil
}
