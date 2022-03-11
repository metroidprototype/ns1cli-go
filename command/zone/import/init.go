// This package implements the zonefile import API
// https://ns1.com/api#put-import-a-zone-from-a-zone-file
package imp

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
	err := command.Register("zone import", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui, a), nil })
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

const synopsis = "Import a zonefile"
const help = `
Usage:
  ns1 zone import [OPTIONS]
  
Options:
  -zone ZONEFILE - zone file; required

Example:
  ns1 zone import -zone ./foo.com

Notes:
  The ZONEFILE name is used as the imported zone name. i.e. a file named "./foo/bar.com" would import to the "bar.com" zone.
`
