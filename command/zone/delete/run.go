package delete

import "fmt"

func (c *cmd) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("zone delete ony accepts a single argument")
		c.UI.Info(c.Help())
		return 1
	}
	zone := args[0]

	_, err := c.ns1.Zones.Delete(zone)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Info(fmt.Sprintf("Zone %s deleted.", zone))
	return 0
}
