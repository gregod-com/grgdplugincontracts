package grgdplugincontracts

// IPluginMetadata ...
type IPluginMetadata interface {
	// GetName This is a Comment ...
	GetName() string
	GetVersion() string
	GetSize() uint64
	GetURL() string
	GetCategory() string
	GetActive() bool
	SetActive(bool)
}
