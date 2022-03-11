package info

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/metroidprototype/ns1cli-go/command/zone/helper"
	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/require"
	"gopkg.in/ns1/ns1-go.v2/mockns1"
	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

func TestInfo(t *testing.T) {
	// Setup the mock service
	mock, doer, err := mockns1.New(t)
	require.Nil(t, err)

	defer mock.Shutdown()

	// Create your NS1 client and configure it for the mock service
	ns1 := api.NewClient(doer, api.SetAPIKey("apikey"))
	ns1.Endpoint, _ = url.Parse("https://" + mock.Address + "/v1/")

	zone := &dns.Zone{
		Zone: "test.com",
		Records: []*dns.ZoneRecord{{
			Domain:   "a.test.com",
			TTL:      300,
			Type:     "A",
			ShortAns: []string{"1.2.3.4"},
		}},
	}

	require.Nil(
		t,
		mock.AddTestCase(
			http.MethodGet,
			"zones/test.com",
			http.StatusOK,
			nil,
			nil,
			"",
			zone,
		),
	)

	var buf bytes.Buffer
	ui := &cli.BasicUi{Writer: &buf, ErrorWriter: &buf}
	cmds, _ := command.RegisterdCommands(ui, ns1)
	cmd, _ := cmds["zone info"]()
	if cmd.Help() != help {
		t.Fatal("zone list help check failed")
	}
	if cmd.Synopsis() != synopsis {
		t.Fatal("zone list synopsis check failed")
	}
	res := cmd.Run([]string{})
	require.Equal(t, 1, res, buf.String())
	buf.Reset()
	res = cmd.Run([]string{"foo.bar"})
	require.Equal(t, 1, res, buf.String())
	buf.Reset()
	res = cmd.Run([]string{"test.com"})
	require.Equal(t, 0, res, buf.String())
	require.Equal(t, helper.FormatZone(ui, ns1, zone, false, false), strings.TrimSuffix(buf.String(), "\n"))
}
