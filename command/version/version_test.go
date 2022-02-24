package version

import (
	"bytes"
	"testing"

	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

func TestVersion(t *testing.T) {
	var buf bytes.Buffer
	ui := &cli.BasicUi{Writer: &buf, ErrorWriter: &buf}
	cmds, _ := command.RegisterdCommands(ui, &api.Client{})
	rec, _ := cmds[cmd_name]()
	if rec.Help() != help {
		t.Fatalf("%s help check failed", cmd_name)
	}
	if rec.Synopsis() != synopsis {
		t.Fatalf("%s synopsis check failed", cmd_name)
	}
}
