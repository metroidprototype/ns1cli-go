package record

import (
	"bytes"
	"strings"
	"testing"

	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

func TestRecord(t *testing.T) {
	var buf bytes.Buffer
	ui := &cli.BasicUi{Writer: &buf, ErrorWriter: &buf}
	cmds, _ := command.RegisterdCommands(ui, &api.Client{})
	rec, _ := cmds["record"]()
	if rec.Help() != help {
		t.Fatal("record synopsis check failed")
	}
	if rec.Synopsis() != synopsis {
		t.Fatal("record synopsis check failed")
	}
	rec.Run([]string{})
	if !strings.Contains(buf.String(), "Must specify a subcommand") {
		t.Fatalf("Unexpected run result: %s", buf.String())
	}
}
