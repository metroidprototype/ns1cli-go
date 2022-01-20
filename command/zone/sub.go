package zone

// This file registers all subcommands for the zone command path
import (
	_ "github.com/metroidprototype/ns1cli-go/command/zone/create"
	_ "github.com/metroidprototype/ns1cli-go/command/zone/delete"
	_ "github.com/metroidprototype/ns1cli-go/command/zone/info"
	_ "github.com/metroidprototype/ns1cli-go/command/zone/list"
	_ "github.com/metroidprototype/ns1cli-go/command/zone/update"
)
