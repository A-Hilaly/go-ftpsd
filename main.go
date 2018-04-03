package main

import (
    "fmt"

    //"github.com/a-hilaly/supfile-api/core"
    "github.com/a-hilaly/supfile-api/server"
    //"github.com/a-hilaly/supfile-api/core/data"
    "github.com/a-hilaly/supfile-api/core/config"
    "github.com/a-hilaly/supfile-api/core/data/engine"
)

// Load and set configuration
func configure() {
    conf := config.LoadConfig()
    engine.SetMagicWordFromConfig(conf.Database)
    engine.Init()
    //data.Init()
    server.Init(conf.Port, "xyz")
    fmt.Println(conf)
}

func main() {
    //fmt.Println(core.AllowFtpProtocol)
    configure()
    server.Run()
}
