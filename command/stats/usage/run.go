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

	cmd := &helper.Cmd{
		Ui:    c.Ui,
		Ns1:   c.Ns1,
		Flags: c.Flags,
	}

	usage, err := GetUsage(cmd, path)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	switch c.Flags.Level {
	case "account":
		output = fmt.Sprintf("%d queries in the last %s across the account.",
			usage, c.Flags.Period)
	case "zone":
		output = fmt.Sprintf("%s: %d queries in the last %s.",
			c.Flags.Zone, usage, c.Flags.Period)
	case "record":
		output = fmt.Sprintf("%s/%s: %d queries in the last %s.",
			fmt.Sprintf("%s.%s", c.Flags.Record, c.Flags.Zone),
			c.Flags.Type, usage, c.Flags.Period)
	}
	c.Ui.Info(output)
	return 0
}

// getUsage is a function to pull usage stats from ns1
func GetUsage(c interface{}, path string) (int64, error) {
	cli := c.(*helper.Cmd)
	rel, err := url.Parse(path)
	if err != nil {
		return 0, err
	}

	uri := cli.Ns1.Endpoint.ResolveReference(rel)
	url := fmt.Sprintf("%s?period=%s&networks=%s", uri, cli.Flags.Period, cli.Flags.Networks)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("X-NSONE-Key", cli.Ns1.APIKey)
	req.Header.Add("User-Agent", cli.Ns1.UserAgent)

	usage := make([]helper.Usage, 1)
	_, err = cli.Ns1.Do(req, &usage)
	if err != nil {
		return 0, err
	}
	return usage[0].Queries, nil
}
