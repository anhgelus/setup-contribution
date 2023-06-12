package main

import (
	"github.com/anhgelus/setup-contribution/src"
	"github.com/anhgelus/setup-contribution/src/config"
	"github.com/pelletier/go-toml/v2"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	u, err := user.Current()
	errPanic(err)
	home := u.HomeDir
	var conf config.Config
	path := filepath.Join(home, config.Datas+"config.toml")
	a, err := os.ReadFile(path)
	errPanic(err)
	errPanic(toml.Unmarshal(a, &conf))
	toUse := "default"
	if len(os.Args) > 1 {
		toUse = os.Args[1]
	}
	errPanic(src.Copy(conf, toUse))
}

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}
