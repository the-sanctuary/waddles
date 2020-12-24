package main

import (
	"fmt"
	"os"
	"plugin"

	wplugin "github.com/the-sanctuary/waddles/pkg/plugin"
	"github.com/the-sanctuary/waddles/pkg/util"
)

func main() {
	util.InitializeLogging()
	util.SetupLogging()

	rawPlugin, err := plugin.Open("./example.so")

	if util.DebugError(err) {
		os.Exit(1)
	}

	ps := wplugin.PluginSystem{}
	wPlugin := ps.LoadPlugin(rawPlugin)
	fmt.Printf("%+v\n", *wPlugin)
}
