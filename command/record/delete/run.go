package delete

import "fmt"

func (c *cmd) Run(args []string) int {
	if len(args) != 3 {
		c.UI.Error("record delete requires three arguments")
		c.UI.Info(c.Help())
		return 1
	}
	zone := args[0]
	rec := args[1]
	rType := args[2]
	_, err := c.ns1.Records.Delete(zone, rec, rType)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Info(fmt.Sprintf("Record %s/%s/%s deleted.", zone, rec, rType))
	return 0
}
