package plugin

import (
	"plugin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/pkg/util"
)

type PluginSystem struct {
}

func (ps *PluginSystem) Logger(version *Version) zerolog.Logger {
	return log.With().Str("system", "PluginSystem").Str("plugin", version.Name).Logger()
}

func (ps *PluginSystem) LoadPlugin(rawPlugin *plugin.Plugin) *WPlugin {
	// rversion, _ := rawPlugin.Lookup("Version")
	// version := rversion.(*Version)

	// fmt.Println(version.Name)
	// fmt.Println(version.HumanName)
	// fmt.Println(version.Version)
	// fmt.Println(version.URL)

	// fmt.Println(reflect.TypeOf(rawPlugin))

	rawPlug, err := rawPlugin.Lookup("Plugin")

	plugPointer := rawPlug.(*WPlugin)

	plug := *plugPointer

	plug.Load(ps)

	if util.DebugError(err) {
		panic(err)
	}

	return &plug
}
