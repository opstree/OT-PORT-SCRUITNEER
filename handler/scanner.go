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
	Severity string
}

// Scanner the scann the given hosts
func Scanner(hosts string, ports []int) (PortInfo, error) {
	var openPortInfo PortInfo
	var openPortList []int
	var listPorts []int
	var severity string
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Minute)
	defer cancel()
	time.Sleep(1 * time.Second)
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(hosts),
		nmap.WithServiceInfo(),
		nmap.WithTimingTemplate(nmap.TimingAggressive),
		nmap.WithContext(ctx),
		// nmap.WithFilterPort(func(p nmap.Port) bool {
		// 	return p.Service.Name == "rtsp"
		// }),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, _, err := scanner.Run()
	// fmt.Println(result)
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
			severity = CriticalityFilter(port.ID)
			fmt.Println(severity)
			fmt.Printf("\tPort %d open with RTSP service\n", port.ID)
		}
		openPorts := filterPorts(listPorts, ports)
		for _, port := range host.Ports {
			for _, portCheck := range openPorts {
				if int(port.ID) == portCheck && port.State.State == "open" {
					openPortList = append(openPortList, portCheck)
				}
			}
		}
	}
	openPortInfo = PortInfo{
		Host:  hosts,
		Ports: openPortList,
		Severity: severity,
	}
	return openPortInfo, err
}

// For check the List o ports that we given in the list
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

func CriticalityFilter(port uint16) string{
	critical := []int{22, 21, 53, 3000, 27017, 27018, 3306, 5432, 6379}
	var severity string
	for por := range critical {
		i := int(port)
		if i == critical[por] {
			severity = "critical"
		}
	}
	return severity
}
