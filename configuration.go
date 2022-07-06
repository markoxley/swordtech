package swordtech

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/markoxley/daggertech"
)

// ServerConfig is the server configuration
type ServerConfig struct {
	Port    int    `json:"port"`
	UseSSL  bool   `json:"usessl"`
	SSLCert string `json:"sslcert"`
	SSLKey  string `json:"sslkey"`
}

// DatabaseConfig is the database configuration
type DatabaseConfig struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"username"`
	Password string `json:"password"`
}

// Config is the application configuration
type Config struct {
	Server          ServerConfig   `json:"server"`
	Database        DatabaseConfig `json:"database"`
	Secret          string         `json:"secret"`
	ApplicationName string         `json:"appname"`
	ShowLog         bool           `json:"showlog"`
}

func loadConfiguration() (*Config, error) {
	configData, err := ioutil.ReadFile("appconfig.json")
	if err != nil {
		return nil, errors.New("Unable to load configuration file")
	}

	configuration := &Config{}
	err = json.Unmarshal(configData, configuration)
	if err != nil {
		return nil, errors.New("Error in configuration file")
	}

	showLog = configuration.ShowLog

	dataConfiguration := daggertech.CreateConfig(configuration.Database.Server, configuration.Database.Database, configuration.Database.User, configuration.Database.Password, false)
	if !daggertech.Configure(dataConfiguration) {
		return nil, errors.New("Unable to configure database")
	}

	seclen := len(configuration.Secret)
	secret := make([]byte, 0, seclen*2)
	secInput := []byte(configuration.Secret)
	for idx, b := range secInput {
		secret = append(secret, b)
		secret = append(secret, secInput[seclen-(idx+1)])
	}
	configuration.Secret = string(secret)

	return configuration, nil
}
