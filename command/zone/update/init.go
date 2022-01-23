// https://ns1.com/api#post-edit-a-zone
package update

import (
	"github.com/metroidprototype/ns1cli-go/command"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	UI cli.Ui
}

func init() {
	command.Register("zone update", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui), nil })
}

func new(ui cli.Ui) *cmd {
	c := &cmd{UI: ui}
	return c
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}

func (c *cmd) Run(args []string) int {
	c.UI.Warn("TODO: not implemented")
	return cli.RunResultHelp
}

const synopsis = "Update a zone"
const help = `
Usage: ns1 zone update [options] ZONE
`
