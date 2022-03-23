// This package implements the usage APIs
// https://ns1.com/api#usage
package usage

import (
	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/metroidprototype/ns1cli-go/command/stats/helper"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	helper.Cmd
}

var cmd_name = "stats usage"

func init() {
	err := command.Register(cmd_name, func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui, a), nil })
	if err != nil {
		panic(err)
	}
}

func new(ui cli.Ui, a *api.Client) *cmd {
	return &cmd{
		Cmd: helper.Cmd{
			Ui:  ui,
			Ns1: a,
			Flags: helper.Flag{
				Level:    "account",
				Period:   "24h",
				Networks: "*",
			},
		},
	}
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}

const synopsis = "View Usage related stats"
const help = `
Usage:
  ns1 stats usage [OPTIONS]

Options:
  -level LEVEL - account, zone or record-level usage; default: account
  -zone ZONE - zone name; required for zone and record usage
  -record FQDN - record fqdn; required for record usage
  -type TYPE - record type; required for record usage
  -period PERIOD - Relative time. Possible values: 1h, 24h, 30d; default: 24h
  -network NETWORK - Identifier of the network for which you want to return usage statistics; default: *

Example:
  ns1 stats usage -level zone -zone foo.com
`
