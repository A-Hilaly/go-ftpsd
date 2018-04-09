package main

import (
    "log"

    "github.com/a-hilaly/supfile-api/server"
    "github.com/a-hilaly/supfile-api/core/data"
    "github.com/a-hilaly/supfile-api/core/config"
    "github.com/a-hilaly/supfile-api/core/data/engine"
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
