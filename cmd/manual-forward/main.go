package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/chad-bekmezian-snap/k8s-port-forwarding/forward"
	"github.com/chad-bekmezian-snap/k8s-port-forwarding/manual/app"
	"github.com/chad-bekmezian-snap/k8s-port-forwarding/manual/service"
	"strings"
	"sync"
)

var appNameToApp = map[string]app.App{
	"ovrc-experience":   app.OvrCExperience,
	"customer-process":  app.CustomerProcess,
	"mobile-experience": app.MobileExperience,
	"license-process":   app.MobileExperience,
}

func main() {
	if len(appFlag) == 0 {
		fmt.Println("No applications specified. Exiting.")
	}

	var wg sync.WaitGroup
	for _, appName := range appFlag {
		services, ok := appNameToApp[appName]
		if !ok {
			fmt.Println(color.Ize(color.Red, "ERROR: Undefined application with name: "+appName))
			fmt.Println("Valid values are:")
			printValidApplications()
			return
		}

		for _, s := range services {
			if omitFlag.Contains(s.Service.Name()) {
				fmt.Println(color.Ize(color.Blue, "Skipping "+s.Service.Name()))
				continue
			}
			wg.Add(1)
			go func(s service.Configurable) {
				portForwardToService(s)
				wg.Done()
			}(s)
		}
	}
	wg.Wait()
}

func portForwardToService(s service.Configurable) {
	var port int
	switch {
	case s.DefaultPort > 0:
		port = s.DefaultPort
	default:
		port = s.Service.DefaultPort
	}

	// Default to port 80 if no port is specified for forwarding
	if s.Service.ForwardToPort == 0 {
		s.Service.ForwardToPort = 80
	}
	s.Service.DefaultPort = port
	forward.ToService(s.Service)
}

func printValidApplications() {
	result := make([]string, 0, len(appNameToApp))
	for key, _ := range appNameToApp {
		result = append(result, " - "+key)
	}

	fmt.Println(strings.Join(result, "\n"))
}
