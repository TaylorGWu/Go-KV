package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type GoKVConfig struct {
	LogConfigPath string `yaml:"log_config_path"`
	Mysql
}

type Mysql struct {
	Host     string
	User     string
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

var config GoKVConfig

func ConfigGoKV() {
	yamlFile, err := ioutil.ReadFile("/home/taylor/code/Go/src/Go-KV/config/go_kv.yaml")
	if err != nil {
		panic("read GoKV config file error")
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic("json unmarshal GoKV config error")
	}
	return
}

func Get() *GoKVConfig {
	return &config
}
