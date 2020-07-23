package grgdplugincontracts

// IPluginMetadata ...
type IPluginMetadata interface {
	// GetFullname returns a more unique identifier of the plugin implemnetation by using the structure:
	// category-name-version (i.e. commands-Basic-v0.1.2)
	GetFullname() string

	GetIdentifier() string

	GetName() string

	// GetVersion returns version in semver standard (i.e. v0.1.2 or v2.1.0)
	GetVersion() string

	// GetURL returns the url of the repo that hosts this plugin
	GetURL() string

	// GetCategory returns the type of plugin
	GetCategory() string

	// GetActive returns the current active state for the plugin (used to persist setting if plugin is used)
	GetActive() bool

	// Getpath returns the current location of the plugin.so file in local filesystem
	GetPath() string

	// Get loaded status
	GetLoaded() bool

	// Set info if plugin was loaded into app
	SetLoaded(bool)

	// Enable/Disable plugin
	SetActive(bool)

	// Enable/Disable plugin depending on current state
	ToggleActive() bool

	// Setpath set the current location of the plugin.so file in local filesystem
	SetPath(string)
}
