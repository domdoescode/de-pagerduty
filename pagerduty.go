package pagerduty

import (
	"github.com/domudall/doiici/plugins"
)

type plugin struct{}

func (p *plugin) Match(input string) string {
	return "TODO"
}

func (p *plugin) GetName() string {
	return "TODO"
}

func (p *plugin) Help() string {
	return "TODO"
}

func init() {
	plugins.Add(&plugin{})
}
