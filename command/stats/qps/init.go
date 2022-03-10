// This package implements the record info API
// https://ns1.com/api#getview-record-details
package qps

import (
	"github.com/metroidprototype/ns1cli-go/command"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type cmd struct {
	UI    cli.Ui
	ns1   *api.Client
	flags flag
}

type flag struct {
	Level  string
	Zone   string
	Record string
	Type   string
}

var cmd_name = "stats qps"

func init() {
	err := command.Register(cmd_name, func(ui cli.Ui, a *api.Client) (cli.Command, error) { return new(ui, a), nil })
	if err != nil {
		panic(err)
	}
}

func new(ui cli.Ui, a *api.Client) *cmd {
	c := &cmd{
		UI:  ui,
		ns1: a,
		flags: flag{
			Level: "account",
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

const synopsis = "View QPS related stats"
const help = `
Usage: ns1 stats qps [OPTIONS]
  Options:
    -level LEVEL - account, zone or record-level qps; default: account
    -zone ZONENAME - zone name; required for zone and record qps
    -record FQDN - record fqdn; required for record qps
    -type TYPE - record type; required for record qps
`
