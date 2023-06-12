package main

import (
	_ "embed"
	"github.com/anhgelus/setup-contribution/src"
	"github.com/anhgelus/setup-contribution/src/config"
	"github.com/pelletier/go-toml/v2"
	"os"
	"os/user"
	"path/filepath"
)

//go:embed example/config.toml
var c string

func main() {
	u, err := user.Current()
	errPanic(err)
	home := u.HomeDir
	toUse := "default"
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err := src.Install(home, c)
			if err != nil {
				panic(err)
			}
			return
		}
		toUse = os.Args[1]
	}
	var conf config.Config
	path := filepath.Join(home, config.Datas+"config.toml")
	a, err := os.ReadFile(path)
	errPanic(err)
	errPanic(toml.Unmarshal(a, &conf))
	for _, dis := range conf.DisabledFolders {
		if toUse == dis {
			println("This folder is disabled")
			return
		}
	}
	errPanic(src.Copy(conf, home, toUse))
}

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}
