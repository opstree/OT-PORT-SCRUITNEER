package handler

import (
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// ConfigFile Structure for config
type ConfigFile struct {
	SlackConfig struct {
		WebhookURL string `yaml:"webhook_url"`
		Username   string `yaml:"username"`
		Channel    string `yaml:"channel"`
	} `yaml:"slack_config"`
	ESConfig struct {
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"es_config"`
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

// func EsClient(config string) (elasticsearch.Config){
// 	es_config, _ := ReadConfig(config)
// 	if es_config.ESConfig.UserName != nil {
// 		cfg := elasticsearch.Config{
// 			Addresses: []string{
// 				fmt.Sprintf("%s:%d", es_config.ESConfig.Address, es_config.ESConfig.Port),
// 			},
// 			// Username: fmt.Sprintf("%s", es_config.ESConfig.UserName),
// 			// Password: fmt.Sprintf("%s", es_config.ESConfig.Password),
// 		}
// 	} else {
// 		cfg := elasticsearch.Config{
// 			Addresses: []string{
// 				fmt.Sprintf("%s:%d", es_config.ESConfig.Address, es_config.ESConfig.Port),
// 			},
// 		}
// 	}
// 	es, err := elasticsearch.NewClient(cfg)
// 	if err != nil {
// 		log.Fatalf("Not able to create ES client", err)
// 	}
// }
