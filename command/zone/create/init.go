// This package implements the zone create API
// https://ns1.com/api#put-create-a-new-zone
package create

import (
	"github.com/metroidprototype/ns1cli-go/command"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	UI  cli.Ui
	ns1 *api.Client
}

func init() {
	err := command.Register("zone create", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui, a), nil })
	if err != nil {
		panic(err)
	}
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

const synopsis = "Create a zone"
const help = `
Usage: ns1 zone create [options] ZONE

Example: ns1 zone create '{"zone":"example.com", "nx_ttl":60}'

Ful payload deatuls can be found here: https://ns1.com/api#put-create-a-new-zone
`
