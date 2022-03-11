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
	Zone   string
	Record string
	Type   string
	QPS    bool
}

func FormatRecord(c *Cmd, rec *dns.Record) string {
	header := "Domain | TTL | Type | Answers"
	line := fmt.Sprintf("%s | %d | %s | %s", rec.Domain, rec.TTL, rec.Type, rec.Answers)
	if c.Flags.QPS {
		qps, _, err := c.Ns1.Stats.GetRecordQPS(rec.Zone, rec.Domain, rec.Type)
		if err != nil {
			c.Ui.Warn(fmt.Sprintf("failed to get QPS for %s", rec.Domain))
		} else {
			header = fmt.Sprintf("%s | QPS (24h)", header)
			line = fmt.Sprintf("%s | %.2f", line, qps)
		}
	}
	result := []string{header, line}
	return columnize.SimpleFormat(result)
}
