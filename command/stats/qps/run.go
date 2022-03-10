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

	flags.ParseArgs(&c.flags, args)

	switch c.flags.Level {
	case "account":
		qps, _, err = c.ns1.Stats.GetQPS()
		res = fmt.Sprintf("Account QPS: %.2f", qps)
	case "zone":
		if len(c.flags.Zone) == 0 {
			c.UI.Error("zone option required for zone level qps")
			return 1
		}
		qps, _, err = c.ns1.Stats.GetZoneQPS(c.flags.Zone)
		res = fmt.Sprintf("%s QPS: %.2f", c.flags.Zone, qps)
	case "record":
		if len(c.flags.Zone) == 0 ||
			len(c.flags.Record) == 0 ||
			len(c.flags.Type) == 0 {
			c.UI.Error("zone, record and type options required for record level qps")
			return 1
		}
		qps, _, err = c.ns1.Stats.GetRecordQPS(c.flags.Zone, c.flags.Record, c.flags.Type)
		res = fmt.Sprintf("%s %s record QPS: %.2f", c.flags.Zone, c.flags.Type, qps)
	default:
		c.UI.Error(fmt.Sprintf("unknown qps level: %s", c.flags.Level))
		return 1
	}

	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Info(res)
	return 0
}
