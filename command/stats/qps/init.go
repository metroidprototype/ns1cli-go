// This package implements the qps APIs
// https://ns1.com/api#qps
package qps

import (
	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/metroidprototype/ns1cli-go/command/stats/helper"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	helper.Cmd
}

var cmd_name = "stats qps"

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
				Level: "account",
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

const synopsis = "View QPS related stats over the last 24h"
const help = `
Usage:
  ns1 stats qps [OPTIONS]

Options:
  -level LEVEL - account, zone or record-level qps; default: account
  -zone ZONE - zone name; required for zone and record qps
  -record FQDN - record fqdn; required for record qps
  -type TYPE - record type; required for record qps

Example:
  ns1 stats qps -level zone -zone foo.com
`
