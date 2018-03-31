package main

import (
    "fmt"

    "github.com/a-hilaly/supfile-api/server"
    "github.com/a-hilaly/supfile-api/core/config"
    "github.com/a-hilaly/supfile-api/core/models"
    "github.com/a-hilaly/supfile-api/core/data/engine"
)

// Load and set configuration
func Configure() {
    conf := config.LoadConfig()
    engine.SetMagicWordFromConfig(conf.Database)
    engine.Init()
    models.Init()
    server.Init(conf.Port)
    fmt.Println(conf)
}

func main() {
    Configure()
    server.Run()
}
