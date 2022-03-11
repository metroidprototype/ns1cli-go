package delete

import (
	"fmt"

	flags "github.com/uber-go/flagoverride"
)

func (c *cmd) Run(args []string) int {
	flags.ParseArgs(&c.Flags, args)
	if len(c.Flags.Zone) == 0 ||
		len(c.Flags.Record) == 0 ||
		len(c.Flags.Type) == 0 {
		c.Ui.Error("zone, record and type options required")
		c.Ui.Info(c.Help())
		return 1
	}
	rec := fmt.Sprintf("%s.%s", c.Flags.Record, c.Flags.Zone)
	_, err := c.Ns1.Records.Delete(c.Flags.Zone, rec, c.Flags.Type)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	c.Ui.Info(fmt.Sprintf("Record %s/%s deleted.", rec, c.Flags.Type))
	return 0
}
