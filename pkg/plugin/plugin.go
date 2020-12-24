package plugin

//WPlugin represents a plugin for Waddles
type WPlugin interface {
	Version() *Version
	Load(ps *PluginSystem)
	Unload(ps *PluginSystem)
}

type wPluginImpl struct {
	version *Version
	load    func()
	unload  func()
}

func (p *wPluginImpl) Version() *Version {
	return p.version
}

func (p *wPluginImpl) Load(ps *PluginSystem) {
	p.load()
}

func (p *wPluginImpl) Unload(ps *PluginSystem) {
	p.unload()
}

//Version holds the plugin version and related information
type Version struct {
	//Name must be the same as the name of your plugin struct that implements plugin.WPlugin
	Name string
	//HumanName is a human readable name
	HumanName string
	//Version holds a Semver (https://semver.org/) string
	Version string
	URL     string
}
