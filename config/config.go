package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Bind_Port int32
	}
	Log struct {
		Path string
	}
	Request struct {
		Cache_Http bool
		Cookie_Jar []string
	}
	Response struct {
		Validators []string
	}
}

func InitConfig() *Config {

	configPath := "/go/src/forxy/"
	if os.Getenv("FORXY_CONFIG_PATH") != "" {
		configPath = os.Getenv("FORXY_CONFIG_PATH")
	}
	yfile, err := ioutil.ReadFile(configPath + "forxy.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var config Config

	err2 := yaml.Unmarshal(yfile, &config)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(config.Response.Validators)
	return &config
}

var Configuration = InitConfig()
