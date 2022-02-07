package create

import (
	"encoding/json"

	"github.com/metroidprototype/ns1cli-go/command/zone/helper"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

func (c *cmd) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("zone create ony accepts a single argument")
		c.UI.Info(c.Help())
		return 1
	}
	zone := args[0]
	var z dns.Zone
	if err := json.Unmarshal([]byte(zone), &z); err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	_, err := c.ns1.Zones.Create(&z)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Info(helper.FormatZone(&z))
	return 0
}
