package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/chad-bekmezian-snap/k8s-port-forwarding/compose/service"
	"os"
	"sort"
	"strings"
)

func printAppsSilently(path string) {
	_ = os.Setenv("FORWARD_SILENT", "true")
	nameToService, err := service.Load(path)
	if err != nil {
		fmt.Println(color.Ize(color.Red, err))
		panic(err)
	}
	printValidApplications(nameToService)
}

func printValidApplications(nameToService map[string]service.Service) {
	result := make(sort.StringSlice, 0, len(nameToService))
	for key, _ := range nameToService {
		result = append(result, " - "+key)
	}
	result.Sort()

	fmt.Println(strings.Join(result, "\n"))
}
