package zone

import (
	"github.com/metroidprototype/ns1cli-go/command"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	UI cli.Ui
}

func init() {
	err := command.Register("zone", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui), nil })
	if err != nil {
		panic(err)
	}
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
	c.UI.Error("Must specify a subcommand")
	return cli.RunResultHelp
}

const synopsis = "Interact with zones"
const help = `
Usage: ns1 zone SUBCOMMAND
`
