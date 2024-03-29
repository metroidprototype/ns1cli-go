// This package implements the record delete API
// https://ns1.com/api#deletedelete-a-record
package delete

import (
	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/metroidprototype/ns1cli-go/command/record/helper"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	helper.Cmd
}

var cmd_name = "record delete"

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
		},
	}
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}

const synopsis = "Delete a record"
const help = `
Usage:
  ns1 record delete [OPTIONS]

Options:
  -zone ZONE - zone name; required
  -record RECORD - record name not including the zone; required
  -type TYPE - record type; required

Example:
  ns1 record delete -zone foo.com -record bar -type CNAME
`
