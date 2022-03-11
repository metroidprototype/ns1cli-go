package info

import (
	"github.com/metroidprototype/ns1cli-go/command/zone/helper"

	flags "github.com/uber-go/flagoverride"
)

func (c *cmd) Run(args []string) int {
	flags.ParseArgs(&c.Flags, args)
	if len(c.Flags.Zone) == 0 {
		c.Ui.Error("zone option required")
		c.Ui.Info(c.Help())
		return 1
	}
	z, _, err := c.Ns1.Zones.Get(c.Flags.Zone)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	c.Ui.Info(helper.FormatZone(c.Ui, c.Ns1, z, c.Flags.QPS, c.Flags.Usage))
	return 0
}
