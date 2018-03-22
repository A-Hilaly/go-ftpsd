package main

import (
    "fmt"
    "io"

    "github.com/a-hilaly/supfile-api/models"
    "github.com/a-hilaly/supfile-api/config"
    "github.com/a-hilaly/supfile-api/engine"
)

var Mode = "DEFAULT"
var Port = ":8080"

func RunWithConfig(conf config.Config, port string, templates string, static string, log_writer io.Writer) {
    e := engines.NewEngine(templates, static, log_writer)
    e.Run(port)
}

func RunDefault() {
    e := engines.Default()
    e.Run(Port)
}

func Run() {
    conf := config.LoadConfig()
    engine.SetMagicWordFromConfig(conf.Database)
    engine.Init()
    models.AutoMigrateModels()

    fmt.Println("Running on port : 8080")
    RunDefault()
}

func main() {
    //models.NewPost("Hella", "llaa", "keyt", "jev", "et")
    Run()
}
