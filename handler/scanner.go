package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap"
)

// Store the information about the open ports
type PortInfo struct {
	Host  string
	Ports []int
}

// Scanner the scann the given hosts
func Scanner(hosts string, ports []int) (PortInfo, error) {
	var openPortInfo PortInfo
	var openPortList []int
	var listPorts []int
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Minute)
	defer cancel()
	time.Sleep(1 * time.Second)
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(hosts),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			listPorts = append(listPorts, int(port.ID))
		}
		openPorts := filterPorts(listPorts, ports)
		for _, port := range host.Ports {
			for _, portCheck := range openPorts {
				if int(port.ID) == portCheck && port.State.State == "open" {
					openPortList = append(openPortList, portCheck)
					// fmt.Printf("This Port is %d %s\n", portCheck, port.State)
				}
			}
		}
	}
	openPortInfo = PortInfo{
		Host:  hosts,
		Ports: openPortList,
	}
	return openPortInfo, err
}

func filterPorts(lista, listb []int) (filter []int) {

	mapFilter := make(map[int]uint8)
	for _, key := range lista {
		mapFilter[key] |= (1 << 0)
	}
	for _, key := range listb {
		mapFilter[key] |= (1 << 1)
	}

	for key, value := range mapFilter {
		lista := value&(1<<0) != 0
		listb := value&(1<<1) != 0
		if lista && !listb {
			filter = append(filter, key)
		}
	}
	return filter
}
