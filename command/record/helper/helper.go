// helper functions for zone commands
package helper

import (
	"fmt"

	"github.com/ryanuber/columnize"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

func FormatRecord(rec *dns.Record) string {
	result := []string{
		"Domain | TTL | Type | Answers",
		fmt.Sprintf("%s | %d | %s | %s",
			rec.Domain, rec.TTL, rec.Type, rec.Answers,
		),
	}
	return columnize.SimpleFormat(result)
}
