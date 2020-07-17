package grgdplugincontracts

// IGrgdPlugin ...
type IGrgdPlugin interface {
	// Init
	// initialize the plugin and return concrete metatdata implementation
	Init(i interface{}) IPluginMetadata
	// GetMetadata return concrete metatdata implementation
	GetMetaData(i interface{}) IPluginMetadata
}
