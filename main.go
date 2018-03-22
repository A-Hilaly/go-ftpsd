package main

import (
    "fmt"
    //"io"

    "github.com/a-hilaly/supfile-api/config"
    "github.com/a-hilaly/supfile-api/server"
    "github.com/a-hilaly/supfile-api/core/engine"
    "github.com/a-hilaly/supfile-api/core/models"
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
