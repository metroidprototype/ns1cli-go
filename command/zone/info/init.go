package info

import (
	"fmt"
	"strings"

	"github.com/metroidprototype/ns1cli-go/command"

	"github.com/mitchellh/cli"
	"github.com/ryanuber/columnize"
	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

type cmd struct {
	UI  cli.Ui
	ns1 *api.Client
}

func init() {
	command.Register("zone info", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui, a), nil })
}

func new(ui cli.Ui, a *api.Client) *cmd {
	c := &cmd{
		UI:  ui,
		ns1: a,
	}
	return c
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}

func (c *cmd) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("zone info only accepts a single argument")
		c.UI.Info(c.Help())
		return 1
	}
	zone := args[0]
	z, h, err := c.ns1.Zones.Get(zone)
	if err != nil {
		c.UI.Error(err.Error())
		return h.StatusCode
	}
	c.UI.Info(formatZone(z))
	return 0
}

func formatZone(z *dns.Zone) string {
	result := []string{"Domain | TTL | Type | Answers"}
	for _, rec := range z.Records {
		result = append(result, fmt.Sprintf("%s | %d | %s | %s",
			rec.Domain, rec.TTL, rec.Type, strings.Join(rec.ShortAns, ", "),
		))
	}
	return columnize.SimpleFormat(result)
}

const synopsis = "View details on a single zone"
const help = `
Usage: ns1 zone info ZONE
`
