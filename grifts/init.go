package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/masterxavierfox/coke/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
