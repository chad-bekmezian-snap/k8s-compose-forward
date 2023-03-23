package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/chad-bekmezian-snap/k8s-port-forwarding/service/compose"
	"github.com/chad-bekmezian-snap/k8s-port-forwarding/service/k8s"
	"strings"
	"sync"
)

func main() {
	if listServices {
		printAppsSilently(yamlPathFlag)
		return
	}

	if completions {
		results, _ := compose.GetBashCompletions(yamlPathFlag)
		fmt.Println(strings.Trim(fmt.Sprint(results), "[]"))
		return
	}

	if len(appArgs)+len(serviceFlag) == 0 {
		fmt.Println("No applications or services specified. Exiting.")
		return
	}

	nameToService, err := compose.Load(yamlPathFlag)
	if err != nil {
		fmt.Println(color.Ize(color.Red, err))
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println(fmt.Sprintf("Starting port-forwarding to services: %v", serviceFlag))
		portForwardForServices(serviceFlag, nameToService)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println(fmt.Sprintf("Starting port-forwarding to dependencies of apps: %v", appArgs))
		portForwardForApps(appArgs, nameToService)
		wg.Done()
	}()

	wg.Wait()
}

func portForwardForServices(services multiValueFlag, nameToService map[string]compose.Service) {
	var wg sync.WaitGroup
	for _, serviceName := range services {
		svc, ok := nameToService[serviceName]
		if !ok {
			fmt.Println(color.Ize(color.Red, "ERROR: Undefined service with name: "+serviceName))
			fmt.Println("Valid values are:")
			printValidApplications(nameToService)
			return
		}

		wg.Add(1)
		go func(s compose.Service) {
			k8s.PortForwardToService(s)
			wg.Done()
		}(svc)
	}
	wg.Wait()
}

func portForwardForApps(apps multiValueFlag, nameToService map[string]compose.Service) {
	var wg sync.WaitGroup
	for _, appName := range apps {
		svc, ok := nameToService[appName]
		if !ok {
			fmt.Println(color.Ize(color.Red, "ERROR: Undefined application with name: "+appName))
			fmt.Println("Valid values are:")
			printValidApplications(nameToService)
			return
		}

		if len(svc.Spec.DependsOn) == 0 {
			fmt.Println(color.Ize(color.Yellow, fmt.Sprintf(`Application "%s" has no dependencies. Skipping.`, appName)))
		}

		for _, s := range svc.Spec.DependsOn {
			if omitFlag.Contains(s) {
				fmt.Println(color.Ize(color.Blue, "Omitting "+s))
				continue
			}

			depSvc, ok := nameToService[s]
			if !ok {
				// TODO: logic handling not existent service
				continue
			}

			wg.Add(1)
			go func(s compose.Service) {
				k8s.PortForwardToService(s)
				wg.Done()
			}(depSvc)
		}
	}
	wg.Wait()
}
