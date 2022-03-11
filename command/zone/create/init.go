// This package implements the zone create API
// https://ns1.com/api#put-create-a-new-zone
package create

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
	err := command.Register("zone create", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui, a), nil })
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

const synopsis = "Create a zone"
const help = `
Usage:
  ns1 zone create [OPTIONS]
  
Options:
  -zone ZONE - json formatted zone object as defined by the ns1 API; required

Example:
  ns1 zone create -zone '{"zone":"example.com", "nx_ttl":60}'

Notes:
  Full payload details can be found here: https://ns1.com/api#put-create-a-new-zone
`
