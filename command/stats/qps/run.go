package qps

import (
	"fmt"

	flags "github.com/uber-go/flagoverride"
)

func (c *cmd) Run(args []string) int {
	var (
		qps float32
		res string
		err error
	)

	flags.ParseArgs(&c.Flags, args)

	switch c.Flags.Level {
	case "account":
		qps, _, err = c.Ns1.Stats.GetQPS()
		res = fmt.Sprintf("Account QPS: %.2f", qps)
	case "zone":
		if len(c.Flags.Zone) == 0 {
			c.Ui.Error("zone option required for zone level qps")
			return 1
		}
		qps, _, err = c.Ns1.Stats.GetZoneQPS(c.Flags.Zone)
		res = fmt.Sprintf("%s QPS: %.2f", c.Flags.Zone, qps)
	case "record":
		if len(c.Flags.Zone) == 0 ||
			len(c.Flags.Record) == 0 ||
			len(c.Flags.Type) == 0 {
			c.Ui.Error("zone, record and type options required for record level qps")
			return 1
		}
		qps, _, err = c.Ns1.Stats.GetRecordQPS(c.Flags.Zone, c.Flags.Record, c.Flags.Type)
		res = fmt.Sprintf("%s %s record QPS: %.2f", c.Flags.Zone, c.Flags.Type, qps)
	default:
		c.Ui.Error(fmt.Sprintf("unknown qps level: %s", c.Flags.Level))
		return 1
	}

	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	c.Ui.Info(res)
	return 0
}
