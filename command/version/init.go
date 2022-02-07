package version

import (
	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/metroidprototype/ns1cli-go/version"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	UI cli.Ui
}

func init() {
	command.Register("version", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui), nil })
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
	c.UI.Info(version.Version)
	return 0
}

const synopsis = "Print the client version"
const help = `
Usage: ns1 version
`
