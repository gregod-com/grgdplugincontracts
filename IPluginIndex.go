package grgdplugincontracts

import (
	"time"
)

// IPluginIndex interface ...
type IPluginIndex interface {
	// Update persist all changes made to entries in pluginlist
	Update(allMetadata map[string]PluginMetadataImpl) error
	// GetConfigPath return location of config file for persistence
	GetConfigPath() string
	// GetPluginList return a map of all plugins
	GetPluginList() map[string]PluginMetadataImpl
	// GetPluginList return a map of all active plugins
	GetPluginListActive() map[string]PluginMetadataImpl
	// GetPluginList return a map of all inactive plugins
	GetPluginListInactive() map[string]PluginMetadataImpl
	// GetLastChecked ...
	GetLastChecked() time.Time
	// AddPlugin add a new plugin to the list
	// Must check if plugin is already loaded
	// Must check if an older (or newer) version of the plugin already exists
	AddPlugin(newplug PluginMetadataImpl) string
}
