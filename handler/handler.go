package handler

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// ConfigFile Structure for config
type ConfigFile struct {
	SlackConfig struct {
		WebhookURL string `yaml:"webhook_url"`
		Username   string `yaml:"username"`
		Channel    string `yaml:"channel"`
	} `yaml:"slack_config"`
	WhitelistPorts []int `yaml:"whitelist_ports"`
	Hosts          []struct {
		Host      string `yaml:"host"`
		Whitelist []int  `yaml:"whitelist"`
	} `yaml:"hosts"`
}

// ReadConfig for reading config file
func ReadConfig(filename string) (*ConfigFile, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Not able to open %q file %v", filename, err)
	}
	config := &ConfigFile{}

	err = yaml.Unmarshal(buf, config)
	if err != nil {
		log.Fatalf("Not able to unmarshal %q file %v", filename, err)
	}
	return config, nil
}
