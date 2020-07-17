package grgdplugincontracts

import (
	"github.com/urfave/cli/v2"
)

// ICMDPlugin ...
type ICMDPlugin interface {
	// GetCommands returns all urfave cli commands implemented by this plugin
	GetCommands(i interface{}) []*cli.Command
}
