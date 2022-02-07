package info

import (
	"github.com/metroidprototype/ns1cli-go/command/zone/helper"
)

func (c *cmd) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("zone info only accepts a single argument")
		c.UI.Info(c.Help())
		return 1
	}
	zone := args[0]
	z, _, err := c.ns1.Zones.Get(zone)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Info(helper.FormatZone(z))
	return 0
}
