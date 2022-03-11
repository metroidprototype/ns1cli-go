package update

import (
	"encoding/json"

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
	z := &dns.Zone{}
	if err := json.Unmarshal([]byte(c.Flags.Zone), z); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	_, err := c.Ns1.Zones.Update(z)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	c.Ui.Info(helper.FormatZone(&c.Cmd, z))
	return 0
}
