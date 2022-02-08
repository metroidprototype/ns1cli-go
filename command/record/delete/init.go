// https://ns1.com/api#deletedelete-a-record
package delete

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
	command.Register("record delete", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui, a), nil })
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

const synopsis = "Delete a record"
const help = `
Usage: ns1 record delete [options] ZONENAME FQDN TYPE
`
