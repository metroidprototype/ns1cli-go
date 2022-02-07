package list

import "github.com/ryanuber/columnize"

func (c *cmd) Run(args []string) int {
	zl, _, err := c.ns1.Zones.List()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	result := []string{"Zone"}
	for _, z := range zl {
		result = append(result, z.Zone)
	}
	c.UI.Info(columnize.SimpleFormat(result))
	return 0
}
