package main

import (
    "log"

    "github.com/a-hilaly/supfile-api/server"
    "github.com/a-hilaly/supfile-api/core/config"
    "github.com/a-hilaly/supfile-api/core/data/engine"
)

// Load and set configuration
func configure() {
    conf, err := config.LoadConfig()
    if err != nil {
        log.Fatal(err)
    }
    engine.SetMagicWordFromConfig(conf.Database)
    engine.Init()
    server.Init(conf.Mode, conf.Port, conf.Auth.Token)
}

func main() {
    configure()
    server.Run()
}
