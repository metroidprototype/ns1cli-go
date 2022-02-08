// https://ns1.com/api#putcreate-a-dns-record
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
	command.Register("record create", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui, a), nil })
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

const synopsis = "Create a record on a zone"
const help = `
Usage: ns1 record create [options] ZONE

Example: ns1 record create '{"zone":"example.com", "domain":"arecord.example.com", "type":"A", "answers":[{"answer":["1.2.3.4"]}]}'

Full payload details can be found here: https://ns1.com/api#putcreate-a-dns-record
`
