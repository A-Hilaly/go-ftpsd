package config

import (
    "io/ioutil"
    "encoding/json"
)

var (
    ConfigFileName string = "config.json"
    GenFileName    string = "config-gen.json"
)

type Config struct {
	Name     string `json:"name"`
    Port     int `json:"port"`
    Mode     string `json:"mode"`
    Auth     AuthConfig `json:"auth"`
	Logging  LoggingConfig `json:"logging"`
	Database DatabaseConfig `json:"database"`
}

type AuthConfig struct {
	Token string `json:"token"`
	Dev   string `json:"dev"`
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

func LoadConfig() (*Config, error) {
    return LoadConfigFromPath(ConfigFileName)
}

func LoadConfigFromPath(path string) (*Config, error) {
    var config Config
    configFile, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    err = json.Unmarshal(configFile, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
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
