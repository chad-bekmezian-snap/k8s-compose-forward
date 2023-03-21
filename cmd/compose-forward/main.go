package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/chad-bekmezian-snap/k8s-port-forwarding/compose/service"
	"github.com/chad-bekmezian-snap/k8s-port-forwarding/forward"
	"os"
	"sort"
	"strings"
	"sync"
)

func main() {
	if listServices {
		_ = os.Setenv("FORWARD_SILENT", "true")
		nameToService, err := service.Load(yamlPathFlag)
		if err != nil {
			fmt.Println(color.Ize(color.Red, err))
			panic(err)
		}
		printValidApplications(nameToService)
		return
	}

	if len(appArgs)+len(serviceFlag) == 0 {
		fmt.Println("No applications or services specified. Exiting.")
		return
	}

	nameToService, err := service.Load(yamlPathFlag)
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

func portForwardForServices(services multiValueFlag, nameToService map[string]service.Service) {
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
		go func(s service.Service) {
			forward.ToService(s)
			wg.Done()
		}(svc)
	}
	wg.Wait()
}

func portForwardForApps(apps multiValueFlag, nameToService map[string]service.Service) {
	var wg sync.WaitGroup
	for _, appName := range apps {
		svc, ok := nameToService[appName]
		if !ok {
			fmt.Println(color.Ize(color.Red, "ERROR: Undefined application with name: "+appName))
			fmt.Println("Valid values are:")
			printValidApplications(nameToService)
			return
		}

		for _, s := range svc.DependsOn {
			if omitFlag.Contains(s) {
				fmt.Println(color.Ize(color.Blue, "Skipping "+s))
				continue
			}

			depSvc, ok := nameToService[s]
			if !ok {
				// TODO: logic handling not existent service
				continue
			}

			wg.Add(1)
			go func(s service.Service) {
				forward.ToService(s)
				wg.Done()
			}(depSvc)
		}
	}
	wg.Wait()
}

func printValidApplications(nameToService map[string]service.Service) {
	result := make(sort.StringSlice, 0, len(nameToService))
	for key, _ := range nameToService {
		result = append(result, " - "+key)
	}
	result.Sort()

	fmt.Println(strings.Join(result, "\n"))
}
