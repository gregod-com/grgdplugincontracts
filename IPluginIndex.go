package grgdplugincontracts

import (
	"time"
)

// IPluginIndex interface ...
type IPluginIndex interface {
	// Update persist all changes made to entries in pluginlist
	Update() error
	// GetConfigPath return location of config file for persistence
	GetConfigPath() string
	// GetPluginList return a map of all plugins
	GetPluginList() map[string]IPluginMetadata
	// // GetPluginList return a map of all active plugins
	GetPluginListActive() []IPluginMetadata
	// // GetPluginList return a map of all inactive plugins
	GetPluginListInactive() []IPluginMetadata

	GetPluginListOffline() []IPluginMetadata

	// GetLastChecked ...
	GetLastChecked() time.Time
	// AddPlugin add a new plugin to the list
	// Must check if plugin is already loaded
	// Must check if an older (or newer) version of the plugin already exists
	AddPlugin(newplug IPluginMetadata) string

	ToggleActive(plugID string) bool

	Finalize(activePlugins []IPluginMetadata, availablePlugins []IPluginMetadata) error

	UnmarshalYAML(unmarshal func(interface{}) error) error
}
