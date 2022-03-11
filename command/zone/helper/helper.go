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
	Zone  string
	QPS   bool
	Usage bool
}

type Zone struct {
	Name    string
	Records []Record
}

type Record struct {
	Domain   string
	TTL      int
	Type     string
	ShortAns []string
	QPS      float32
	Usage    int
}

func FormatZone(ui cli.Ui, ns1 *api.Client, z *dns.Zone, qps bool, usage bool) string {
	header := "Domain | TTL | Type | Answers"
	if qps {
		header = fmt.Sprintf("%s | QPS (24h)", header)
	}
	if usage {
		header = fmt.Sprintf("%s | Usage", header)
	}
	result := []string{header}
	for _, rec := range z.Records {
		record := fmt.Sprintf("%s | %d | %s | %s",
			rec.Domain, rec.TTL, rec.Type, strings.Join(rec.ShortAns, ", "),
		)
		if qps {
			qps, _, err := ns1.Stats.GetRecordQPS(z.Zone, rec.Domain, rec.Type)
			if err != nil {
				ui.Warn(fmt.Sprintf("failed to get QPS for %s", rec.Domain))
			} else {
				record = fmt.Sprintf("%s | %.2f", record, qps)
			}
		}
		result = append(result, record)
	}
	return columnize.SimpleFormat(result)
}
