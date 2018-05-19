package main

import (
	"log"

	"github.com/a-hilaly/go-ftpsd/core/config"
	"github.com/a-hilaly/go-ftpsd/core/data"
	"github.com/a-hilaly/go-ftpsd/core/data/engine"
	"github.com/a-hilaly/go-ftpsd/server"
)

// Load and set configuration
func configureServer() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	engine.SetMagicWordFromConfig(conf.Database)
	err_engine := engine.Init()
	if err_engine != nil {
		log.Fatal(err_engine)
	}
	err_autom := data.AutoMigrate()
	if err_autom != nil {
		log.Fatal(err_autom)
	}
	server.Init(conf.Mode, conf.Port, conf.Auth.Token)
}

func main() {
	configureServer()
	server.Run()
}
