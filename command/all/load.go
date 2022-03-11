// This package initializes all top level commands via their init() functions
package all

import (
	_ "github.com/metroidprototype/ns1cli-go/command/record"
	_ "github.com/metroidprototype/ns1cli-go/command/stats"
	_ "github.com/metroidprototype/ns1cli-go/command/version"
	_ "github.com/metroidprototype/ns1cli-go/command/zone"
)
