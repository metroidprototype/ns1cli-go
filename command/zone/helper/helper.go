// helper functions for zone commands
package helper

import (
	"fmt"
	"strings"

	"github.com/ryanuber/columnize"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

func FormatZone(z *dns.Zone) string {
	result := []string{"Domain | TTL | Type | Answers"}
	for _, rec := range z.Records {
		result = append(result, fmt.Sprintf("%s | %d | %s | %s",
			rec.Domain, rec.TTL, rec.Type, strings.Join(rec.ShortAns, ", "),
		))
	}
	return columnize.SimpleFormat(result)
}
