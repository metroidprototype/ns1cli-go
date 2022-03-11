// This package implements the zone info API
// https://ns1.com/api#get-view-zone-details
package info

import (
	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/metroidprototype/ns1cli-go/command/zone/helper"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	helper.Cmd
}

func init() {
	err := command.Register("zone info", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui, a), nil })
	if err != nil {
		panic(err)
	}
}

func new(ui cli.Ui, a *api.Client) *cmd {
	return &cmd{
		Cmd: helper.Cmd{
			Ui:  ui,
			Ns1: a,
		},
	}
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}

const synopsis = "View details on a single zone"
const help = `
Usage: ns1 zone info [OPTIONS]
  Options:
    -zone ZONENAME - zone name; required
    -qps - include QPS metrics on each record over the last 24h; Note: may increase response time
    -usage - include usage statistics on each record; TODO
`
