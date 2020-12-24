package main

import (
	"github.com/the-sanctuary/waddles/pkg/waddles"
)

func main() {
	wadl := new(waddles.Waddles)

	wadl.LoadPlugins()

	wadl.Run()
}
