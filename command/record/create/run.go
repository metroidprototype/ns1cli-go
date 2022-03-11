package create

import (
	"encoding/json"

	"github.com/metroidprototype/ns1cli-go/command/record/helper"
	flags "github.com/uber-go/flagoverride"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
	"gopkg.in/ns1/ns1-go.v2/rest/model/filter"
)

func (c *cmd) Run(args []string) int {
	flags.ParseArgs(&c.Flags, args)
	if len(c.Flags.Record) == 0 {
		c.Ui.Error("record option required")
		c.Ui.Info(c.Help())
		return 1
	}
	r := &dns.Record{}
	if err := json.Unmarshal([]byte(c.Flags.Record), r); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	if r.Filters == nil {
		r.Filters = make([]*filter.Filter, 0)
	}
	_, err := c.Ns1.Records.Create(r)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	c.Ui.Info(helper.FormatRecord(&c.Cmd, r))
	return 0
}
