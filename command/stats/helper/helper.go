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
	Level    string
	Zone     string
	Record   string
	Type     string
	Period   string
	Networks string
}

type Usage struct {
	Queries int64  `json:"queries,omitempty"`
	Zone    string `json:"zone,omitempty"`
	Domain  string `json:"domain,omitempty"`
	Type    string `json:"rectype,omitempty"`
	Period  string `json:"period,omitempty"`
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
