package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

type Config struct {
	DB DbConfig `yaml:"DB"`
}

func (c *Config) Init(rootPath string) {
	log.Println(rootPath)
	filePath := fmt.Sprintf("%s/conf.yaml", rootPath)
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("failed to connect database")
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
}
