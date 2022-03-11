package delete

import "fmt"

func (c *cmd) Run(args []string) int {
	if len(c.Flags.Zone) == 0 {
		c.Ui.Error("zone option required")
		c.Ui.Info(c.Help())
		return 1
	}

	_, err := c.Ns1.Zones.Delete(c.Flags.Zone)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	c.Ui.Info(fmt.Sprintf("Zone %s deleted.", c.Flags.Zone))
	return 0
}
