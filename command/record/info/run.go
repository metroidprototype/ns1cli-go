package info

import "github.com/metroidprototype/ns1cli-go/command/record/helper"

func (c *cmd) Run(args []string) int {
	if len(args) != 3 {
		c.UI.Error("record info requires three arguments")
		c.UI.Info(c.Help())
		return 1
	}
	zone := args[0]
	domain := args[1]
	rType := args[2]
	z, _, err := c.ns1.Records.Get(zone, domain, rType)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Info(helper.FormatRecord(z))
	return 0
}
