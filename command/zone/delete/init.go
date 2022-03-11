// This package implements the zone delete API
// https://ns1.com/api#delete-delete-a-zone
package delete

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
	err := command.Register("zone delete", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui, a), nil })
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

const synopsis = "Delete a zone"
const help = `
Usage:
  ns1 zone delete [OPTIONS]
  
Options:
  -zone ZONE - zone name; required

Example:
  ns1 zone delete -zone foo.com
`
