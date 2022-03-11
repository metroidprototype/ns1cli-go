// helper functions for zone commands
package helper

import (
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/ryanuber/columnize"
	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

type Cmd struct {
	Ui    cli.Ui
	Ns1   *api.Client
	Flags Flag
}

type Flag struct {
	Level  string
	Zone   string
	Record string
	Type   string
}

func FormatRecord(rec *dns.Record) string {
	result := []string{
		"Domain | TTL | Type | Answers",
		fmt.Sprintf("%s | %d | %s | %s",
			rec.Domain, rec.TTL, rec.Type, rec.Answers,
		),
	}
	return columnize.SimpleFormat(result)
}