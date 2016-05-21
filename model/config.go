package model

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Webserver struct {
		Enable  bool   `yaml:"enable"`
		Api     bool   `yaml:"api"`
		Port    string `yaml:"port"`
		Address string `yaml:"address"`
		Webroot string `yaml:"webroot"`
	} `yaml:"webserver"`
	Database struct {
		Path string `yaml:"path"`
	} `yaml:"database"`
	DNSServer struct {
		Enable     bool   `yaml:"enable"`
		Domain     string `yaml:"domain"`
		IPV6Prefix string `yaml:"ipv6prefix"`
		IPV4Prefix string `yaml:"ipv4prefix"`
	} `yaml:"dnsserver"`
}

func ReadConfigFile(path string) *Config {
	config := &Config{}
	file, _ := ioutil.ReadFile(path)
	err := yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
