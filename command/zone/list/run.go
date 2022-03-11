package list

import (
	"fmt"

	"github.com/ryanuber/columnize"
	flags "github.com/uber-go/flagoverride"
)

func (c *cmd) Run(args []string) int {
	flags.ParseArgs(&c.Flags, args)
	zl, _, err := c.Ns1.Zones.List()
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	header := "Zone"
	if c.Flags.QPS {
		header = fmt.Sprintf("%s | QPS (24h)", header)
	}
	result := []string{header}
	for _, z := range zl {
		line := z.Zone
		if c.Flags.QPS {
			qps, _, err := c.Ns1.Stats.GetZoneQPS(z.Zone)
			if err != nil {
				c.Ui.Warn(fmt.Sprintf("failed to get QPS for %s", z.Zone))
			}
			line = fmt.Sprintf("%s | %.2f", line, qps)
		}
		result = append(result, line)
	}
	c.Ui.Info(columnize.SimpleFormat(result))
	return 0
}
