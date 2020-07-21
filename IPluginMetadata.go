package grgdplugincontracts

// IPluginMetadata ...
type IPluginMetadata interface {
	// GetFullname returns a more unique identifier of the plugin implemnetation by using the structure:
	// category-name-version (i.e. commands-Basic-v0.1.2)
	GetFullname() string

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

	// Enable/Disable plugin
	SetActive(bool)

	// Enable/Disable plugin depending on current state
	ToggleActive()

	// Setpath set the current location of the plugin.so file in local filesystem
	SetPath(string)
}

// PluginMetadataImpl ...
type PluginMetadataImpl struct {
	Name     string
	Version  string
	URL      string
	Category string
	Active   bool
	Path     string
}

// GetFullname ...
func (plmeta PluginMetadataImpl) GetFullname() string {
	return plmeta.Category + "-" + plmeta.Name + "-" + plmeta.Version
}

// GetName ...
func (plmeta PluginMetadataImpl) GetName() string {
	return plmeta.Name
}

// GetVersion ...
func (plmeta PluginMetadataImpl) GetVersion() string {
	return plmeta.Version
}

// GetURL ...
func (plmeta PluginMetadataImpl) GetURL() string {
	return plmeta.URL
}

// GetCategory ...
func (plmeta PluginMetadataImpl) GetCategory() string {
	return plmeta.Category
}

// GetActive ...
func (plmeta PluginMetadataImpl) GetActive() bool {
	return plmeta.Active
}

// GetPath ...
func (plmeta PluginMetadataImpl) GetPath() string {
	return plmeta.Path
}

// SetActive ...
func (plmeta *PluginMetadataImpl) SetActive(a bool) {
	plmeta.Active = a
}

// ToggleActive ...
func (plmeta *PluginMetadataImpl) ToggleActive() {
	plmeta.Active = !plmeta.Active
}

// SetPath ...
func (plmeta *PluginMetadataImpl) SetPath(path string) {
	plmeta.Path = path
}
