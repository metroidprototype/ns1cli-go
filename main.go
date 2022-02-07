package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/metroidprototype/ns1cli-go/command"
	_ "github.com/metroidprototype/ns1cli-go/command/all"
	"github.com/metroidprototype/ns1cli-go/version"

	"github.com/mitchellh/cli"
	api "gopkg.in/ns1/ns1-go.v2/rest"
)

const NAME = "ns1"

func main() {
	os.Exit(runCli())
}

func runCli() int {
	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}
	c, exitCode := createNS1Cli(ui)
	if exitCode != 0 {
		return exitCode
	}

	cmds, names := command.RegisterdCommands(ui, c)

	cli := &cli.CLI{
		Args:         os.Args[1:],
		Commands:     cmds,
		Autocomplete: true,
		Name:         NAME,
		Version:      version.Version,
		HelpFunc:     cli.FilteredHelpFunc(names, cli.BasicHelpFunc(NAME)),
		HelpWriter:   os.Stdout,
		ErrorWriter:  os.Stderr,
	}

	if cli.IsVersion() {
		cli.Args = []string{"version"}
	}

	exitCode, err := cli.Run()
	if err != nil {
		ui.Error(fmt.Sprintf("Error executing CLI: %v", err))
		return 1
	}

	return exitCode
}

func createNS1Cli(ui cli.Ui) (*api.Client, int) {
	k := os.Getenv("NS1_APIKEY")
	if k == "" {
		ui.Error("NS1_APIKEY environment variable is not set, giving up")
		return nil, 1
	}

	httpClient := &http.Client{Timeout: time.Second * 10}

	return api.NewClient(httpClient, api.SetAPIKey(k)), 0
}
