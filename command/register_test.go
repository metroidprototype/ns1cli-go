package command

import (
	"testing"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

type testCommand struct{}

func (c *testCommand) Synopsis() string {
	return "synopsis"
}

func (c *testCommand) Help() string {
	return "help"
}

func (c *testCommand) Run(args []string) int {
	return 0
}

func TestRegistry(t *testing.T) {
	Register("test", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return &testCommand{}, nil })
	cmds, names := RegisterdCommands(&cli.BasicUi{}, &api.Client{})
	if len(names) != 1 {
		t.Fatalf("more registered commands than expected")
	}
	test, _ := cmds["test"]()
	if test.Help() != "help" {
		t.Fatal("help check failed")
	}
	if test.Synopsis() != "synopsis" {
		t.Fatal("synopsis check failed")
	}
}

func TestDupeRegister(t *testing.T) {
	Register("test", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return &testCommand{}, nil })
	err := Register("test", func(ui cli.Ui, a *api.Client) (cli.Command, error) { return &testCommand{}, nil })
	if err == nil {
		t.Fatal("duplicate registration allowed")
	}
}
