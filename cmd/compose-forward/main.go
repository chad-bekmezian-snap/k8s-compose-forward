package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/chad-bekmezian-snap/cs-port-forwarding/compose/service"
	"github.com/chad-bekmezian-snap/cs-port-forwarding/forward"
	"sort"
	"strings"
	"sync"
)

func main() {
	if len(appFlag) == 0 {
		fmt.Println("No applications specified. Exiting.")
		return
	}

	nameToService, err := service.Load(yamlPathFlag)
	if err != nil {
		fmt.Println(color.Ize(color.Red, err))
		panic(err)
	}

	var wg sync.WaitGroup
	for _, appName := range appFlag {
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
