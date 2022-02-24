// The command package provides a global registry of commands for the cli.
package command

import (
	"fmt"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

// Factory is a function that returns a new instance of a CLI-sub command.
type Factory func(cli.Ui, *api.Client) (cli.Command, error)

// Register adds a new CLI sub-command to the registry.
func Register(name string, fn Factory) error {
	if registry == nil {
		registry = make(map[string]Factory)
	}

	if registry[name] != nil {
		return fmt.Errorf("command %q is already registered", name)
	}
	registry[name] = fn
	return nil
}

// RegisterdCommands returns a mapping of all registered commands, and a
// list of their names for consumpution by the CLI.
func RegisterdCommands(ui cli.Ui, c *api.Client) (map[string]cli.CommandFactory, []string) {
	m := make(map[string]cli.CommandFactory)
	var n []string
	for name, fn := range registry {
		thisFn := fn
		m[name] = func() (cli.Command, error) {
			return thisFn(ui, c)
		}
		n = append(n, name)
	}
	return m, n
}

// registry has an entry for each available CLI sub-command, indexed by sub
// command name. This should be populated at package init() time via Register().
var registry map[string]Factory
