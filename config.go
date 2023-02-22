package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// config type used to represent the confiuration of the cli tool.
// right now just looks for the api key
type config struct {
	Key string `yaml:"key"`
}

// getConfig parses the config file and returns the values as a struct
func getConfig(filename string) (*config, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return &config{}, err
	}
	c := config{}
	err = yaml.Unmarshal(yamlFile, &c)
	return &c, err
}
