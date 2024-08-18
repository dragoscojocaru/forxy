package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Server struct {
		BindPort int
	}
	Log struct {
		Path string
	}
	Request struct {
		CacheTcp  bool
		CookieJar []string
	}
	Response struct {
		Validators []string
	}
}

func InitConfig() *Config {

	yfile, err := ioutil.ReadFile("forxy.yaml")

	if err != nil {
		log.Fatal(err)
	}

	config := new(Config)

	err2 := yaml.Unmarshal(yfile, &config)

	if err2 != nil {
		log.Fatal(err2)
	}

	return config
}

var Configuration = InitConfig()
