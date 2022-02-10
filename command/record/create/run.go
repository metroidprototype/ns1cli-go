package create

import (
	"encoding/json"

	"github.com/metroidprototype/ns1cli-go/command/record/helper"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
	"gopkg.in/ns1/ns1-go.v2/rest/model/filter"
)

func (c *cmd) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error("record create requires a single argument")
		c.UI.Info(c.Help())
		return 1
	}
	rec := args[0]
	var r dns.Record
	if err := json.Unmarshal([]byte(rec), &r); err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	if r.Filters == nil {
		r.Filters = make([]*filter.Filter, 0)
	}
	_, err := c.ns1.Records.Create(&r)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Info(helper.FormatRecord(&r))
	return 0
}
