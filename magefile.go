//+build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

//Default target is Run
var Default = Run

func DownloadMod() error {
	return sh.RunV("go", "mod", "download")
}

func Build() error {
	mg.Deps(DownloadMod)
	return sh.RunV("go", "build", "./...")
}

func Run() error {
	mg.Deps(DownloadMod)
	return sh.RunV("go", "run", ".")
}

func Test() {
	mg.Deps(DownloadMod)
	sh.RunV("go", "test", "./...", "-cover")
}
