package handler

import (
	"fmt"
	"log"

	"github.com/deveshs23/port-scanner/notification"
)

type DataStruct struct {
	// Hosts []struct {
	Host  string `json:"host"`
	Ports []struct {
		Port       []int  `json:"port"`
		Severity   string `json:"severity"`
		Service    string `json:"service"`
		Tag        string `json:"tag"`
		Compliance bool   `json:"compliance"`
	} `json:"ports"`
	// } `json:"Hosts"`
}

func getData() []DataStruct {
	var getDataStruct []DataStruct
	config, _ := ReadConfig("config.yaml")
	// fmt.Println(config)
	for _, host := range config.Hosts {
		data, err := Scanner(host.Host, host.Whitelist)
		if err != nil {
			log.Fatalf("cehck %v", err)
		}
		// fmt.Println(data)

		sc := notification.SlackClient{
			WebHookUrl: fmt.Sprint(config.SlackConfig.WebhookURL),
			UserName:   fmt.Sprint(config.SlackConfig.Username),
			Channel:    fmt.Sprint(config.SlackConfig.Channel),
		}
		if data.Ports != nil {
			// To send a notification with status (slack attachments)
			sj := notification.SlackJobNotification{
				Text:      fmt.Sprintf("Host: %s\nPorts: %v", data.Host, data.Ports),
				Details:   "Geting information about open ports",
				Color:     "danger",
				IconEmoji: ":hammer_and_wrench",
			}
			getDataStruct = append(getDataStruct, DataStruct{
				Host:  data.Host,
				Ports: data.Ports,
				
			})
			err = sc.SendJobNotification(sj)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return getDataStruct
}
