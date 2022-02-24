package list

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"

	"github.com/metroidprototype/ns1cli-go/command"
	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/require"
	"gopkg.in/ns1/ns1-go.v2/mockns1"
	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

func TestCreate(t *testing.T) {
	// Setup the mock service
	mock, doer, err := mockns1.New(t)
	require.Nil(t, err)

	defer mock.Shutdown()

	// Create your NS1 client and configure it for the mock service
	ns1 := api.NewClient(doer, api.SetAPIKey("apikey"))
	ns1.Endpoint, _ = url.Parse("https://" + mock.Address + "/v1/")

	require.Nil(t, mock.AddTestCase(http.MethodGet, "zones", http.StatusOK, nil, nil, "",
		[]*dns.Zone{{Zone: "test.com"}}))

	var buf bytes.Buffer
	ui := &cli.BasicUi{Writer: &buf, ErrorWriter: &buf}
	cmds, _ := command.RegisterdCommands(ui, ns1)
	rec, _ := cmds["zone list"]()
	if rec.Help() != help {
		t.Fatal("zone list help check failed")
	}
	if rec.Synopsis() != synopsis {
		t.Fatal("zone list synopsis check failed")
	}
	res := rec.Run([]string{})
	require.Equal(t, 0, res)
}
