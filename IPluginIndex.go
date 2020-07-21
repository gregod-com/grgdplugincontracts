package grgdplugincontracts

import (
	"time"
)

// IPluginIndex interface ...
type IPluginIndex interface {
	// Update persist all changes made to entries in pluginlist
	Update(allMetadata map[string]IPluginMetadata) error
	// GetConfigPath return location of config file for persistence
	GetConfigPath() string
	// GetPluginList return a map of all plugins
	GetPluginList() []IPluginMetadata
	// GetPluginList return a map of all active plugins
	GetPluginListActive() []IPluginMetadata
	// GetPluginList return a map of all inactive plugins
	GetPluginListInactive() []IPluginMetadata
	// GetLastChecked ...
	GetLastChecked() time.Time
	// AddPlugin add a new plugin to the list
	// Must check if plugin is already loaded
	// Must check if an older (or newer) version of the plugin already exists
	AddPlugin(newplug IPluginMetadata) string
}
