package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	Key string `yaml:"key"`
}

func getConfig(filename string) (*config, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return &config{}, err
	}
	c := config{}
	err = yaml.Unmarshal(yamlFile, &c)
	return &c, err
}
