// helper functions for zone commands
package helper

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/ryanuber/columnize"
	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

// Cmd is the zone command object
type Cmd struct {
	Ui    cli.Ui
	Ns1   *api.Client
	Flags Flag
}

// Flag represents all flags used by zone commands
type Flag struct {
	Zone   string
	QPS    bool
	Usage  bool
	Period string
}

func FormatZone(c *Cmd, z *dns.Zone) string {
	header := "Domain | TTL | Type | Answers"
	if c.Flags.QPS {
		header = fmt.Sprintf("%s | QPS(24h)", header)
	}
	if c.Flags.Usage {
		header = fmt.Sprintf("%s | Queries(%s)", header, c.Flags.Period)
	}
	result := []string{header}
	for _, rec := range z.Records {
		record := fmt.Sprintf("%s | %d | %s | %s",
			rec.Domain, rec.TTL, rec.Type, strings.Join(rec.ShortAns, ", "),
		)
		if c.Flags.QPS {
			qps, _, err := c.Ns1.Stats.GetRecordQPS(z.Zone, rec.Domain, rec.Type)
			if err != nil {
				c.Ui.Warn(fmt.Sprintf("failed to get QPS for %s", rec.Domain))
			} else {
				record = fmt.Sprintf("%s | %.2f", record, qps)
			}
		}
		result = append(result, record)
	}
	return columnize.SimpleFormat(result)
}
