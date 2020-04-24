package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Configuration struct {
	Kafka    kafka
	Postgres postgres
}

type kafka struct {
	Bootstrap string
	Port      string
}

type postgres struct {
	User     string
	Password string
	Db       string
	Port     string
	Url      string
}

func GetConfig(baseConfig *Configuration) {
	basePath, _ := os.Getwd()
	if _, err := toml.DecodeFile(basePath+"/config.toml", &baseConfig); err != nil {
		fmt.Println(err)
	}
}
