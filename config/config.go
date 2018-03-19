package config


import (
    "encoding/json"
    "io/ioutil"
    "fmt"
)


type Config struct {
	Name  string `json:"name"`
	Logging LoggingConfig `json:"logging"`
	Database DatabaseConfig `json:"database"`
}

type LoggingConfig struct {
    Type     string `json:"type"`
    Logdir   string `json:"logdir"`
    Logtable string `json:"logtable"`
}

type DatabaseConfig struct {
    Type  string `json:"type"`
    Creds struct {
        Database string `json:"database"`
        Host     string `json:"host"`
        Port     int    `json:"port"`
        User     string `json:"user"`
        Password string `json:"password"`
    } `json:"creds"`
}


var ConfigFileName string = "config.json"


func LoadConfig() Config {
    var config Config
    configFile, err := ioutil.ReadFile(ConfigFileName)
    if err != nil {
        fmt.Println(err)
        return Config{}
    }
    err = json.Unmarshal(configFile, &config)
    if err != nil {
        fmt.Println(err)
        return Config{}
    }
    return config
}

func GenerateConfig(filename string) error {
    config := Config{}
    jsonBytes, err := json.Marshal(config)
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(filename, jsonBytes, 0644)
    return err
}
