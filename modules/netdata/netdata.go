package netdata

import "github.com/netdata/go.d.plugin/agent/module"

type Netdata struct {
	module.Base
	Config `yaml:",inline"`

	charts *module.Charts
	client piholeAPIClient
}
