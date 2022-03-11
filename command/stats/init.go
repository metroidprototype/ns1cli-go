package stats

import (
	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/metroidprototype/ns1cli-go/command/stats/helper"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	helper.Cmd
}

func init() {
	err := command.Register("stats", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui), nil })
	if err != nil {
		panic(err)
	}
}

func new(ui cli.Ui) *cmd {
	c := &cmd{
		Cmd: helper.Cmd{
			Ui: ui,
		},
	}
	return c
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}

func (c *cmd) Run(args []string) int {
	c.Ui.Error("Must specify a subcommand")
	return cli.RunResultHelp
}

const synopsis = "Interact with reporting and stats APIs"
const help = `
Usage: ns1 stats SUBCOMMAND
`