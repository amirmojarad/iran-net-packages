package config

import (
	"github.com/BurntSushi/toml"
	"internet_service/internal/utils"
	"os"
	"sync"
)

type Config struct {
	Providers struct {
		Irancell   string `json:"Irancell"`
		Rightell   string `json:"Rightell"`
		HamrahAval string `json:"HamrahAval"`
	}
	Databases struct {
		Postgres struct {
			Host     string `json:"Host"`
			User     string `json:"User"`
			Password string `json:"Password"`
			Name     string `json:"Name"`
			Port     int    `json:"Port"`
			TimeZone string `json:"TimeZone"`
		}

		Test struct {
			Host     string `json:"Host"`
			User     string `json:"User"`
			Password string `json:"Password"`
			Name     string `json:"Name"`
			Port     int    `json:"Port"`
			TimeZone string `json:"TimeZone"`
		}
	}
}

var lock = &sync.Mutex{}

var Instance *Config

func loadTomlFile(tomlFilePath string) string {
	rootPath := utils.GetPathAsString()
	tomlContent, err := os.ReadFile(rootPath + tomlFilePath)
	if err != nil {
		panic(err)
	}
	return string(tomlContent)
}

// getConf returns Conf instant, loads env file and conf file combine them together
// convert it to Conf struct to use.
func getTomlConf() (*Config, error) {
	// first load env file
	confContent := loadTomlFile("/conf.toml")
	conf := &Config{}
	if _, err := toml.Decode(confContent, &conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func GetConfToml() *Config {
	if Instance == nil {
		lock.Lock()
		defer lock.Unlock()
		createdConf, err := getTomlConf()
		if err != nil {
			return nil
		}
		Instance = createdConf
		return Instance
	} else {
		return Instance
	}
}
