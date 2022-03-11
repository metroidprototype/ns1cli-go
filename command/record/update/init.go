// This package implements the update record command
// https://ns1.com/api#postupdate-a-record
package update

import (
	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/metroidprototype/ns1cli-go/command/record/helper"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	helper.Cmd
}

var cmd_name = "record update"

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

const synopsis = "Update a record"
const help = `
Usage:
  ns1 record update [OPTIONS]

Options:
  -record RECORD - record object as defined by the NS1 API; required

Example:
  ns1 record update -record '{"zone":"example.com", "domain":"arecord.example.com", "type":"A", "answers":[{"answer":["1.2.3.5"]}]}'

Full payload details can be found here: https://ns1.com/api#postupdate-a-record
`
