package usage

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/metroidprototype/ns1cli-go/command/stats/helper"
	flags "github.com/uber-go/flagoverride"
)

func (c *cmd) Run(args []string) int {
	flags.ParseArgs(&c.Flags, args)
	var path, output string

	switch c.Flags.Level {
	case "account":
		path = "stats/usage"
	case "zone":
		if len(c.Flags.Zone) == 0 {
			c.Ui.Error("zone option required for zone level usage")
			return 1
		}
		path = fmt.Sprintf("stats/usage/%s", c.Flags.Zone)
	case "record":
		if len(c.Flags.Zone) == 0 ||
			len(c.Flags.Record) == 0 ||
			len(c.Flags.Type) == 0 {
			c.Ui.Error("zone, record and type options required for record level qps")
			return 1
		}
		rec := fmt.Sprintf("%s.%s", c.Flags.Record, c.Flags.Zone)
		path = fmt.Sprintf("stats/usage/%s/%s/%s", c.Flags.Zone, rec, c.Flags.Type)
	default:
		c.Ui.Error(fmt.Sprintf("unknown usage level: %s", c.Flags.Level))
		return 1
	}
	usage := make([]helper.Usage, 1)
	err := c.GetUsage(path, c.Flags, &usage)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	switch c.Flags.Level {
	case "account":
		output = fmt.Sprintf("%d queries in the last %s across the account.", usage[0].Queries, usage[0].Period)
	case "zone":
		output = fmt.Sprintf("%s: %d queries in the last %s.", usage[0].Zone, usage[0].Queries, usage[0].Period)
	case "record":
		output = fmt.Sprintf("%s/%s: %d queries in the last %s.", usage[0].Domain, usage[0].Type, usage[0].Queries, usage[0].Period)
	}
	c.Ui.Info(output)
	return 0
}

// getUsage is a function to pull usage stats from ns1
func (c *cmd) GetUsage(path string, flag helper.Flag, res *[]helper.Usage) error {
	rel, err := url.Parse(path)
	if err != nil {
		return err
	}

	uri := c.Ns1.Endpoint.ResolveReference(rel)
	url := fmt.Sprintf("%s?period=%s&networks=%s", uri, c.Flags.Period, c.Flags.Networks)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-NSONE-Key", c.Ns1.APIKey)
	req.Header.Add("User-Agent", c.Ns1.UserAgent)

	_, err = c.Ns1.Do(req, &res)
	if err != nil {
		return err
	}
	return nil
}
