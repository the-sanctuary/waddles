package plugin

import (
	"fmt"
	"os"
	"plugin"
	"testing"

	"github.com/the-sanctuary/waddles/pkg/util"
	"github.com/the-sanctuary/waddles/pkg/waddles"
)

func Test_Test1(t *testing.T) {
	os.Chdir("../../")

	wadl := new(waddles.Waddles)

	wadl.LoadPlugins()
}
func Test_LoadPlugin(t *testing.T) {
	util.InitializeLogging()
	util.SetupLogging()

	wd, _ := os.Getwd()

	rawPlugin, err := plugin.Open(wd + "/plugins/example.so")

	if util.DebugError(err) {
		// t.Fail()
	}

	ps := PluginSystem{}
	wPlugin := ps.LoadPlugin(rawPlugin)
	fmt.Printf("%+v\n", wPlugin)
}
