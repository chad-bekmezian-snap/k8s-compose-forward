package main

import (
	"fmt"
	"github.com/chad-bekmezian-snap/k8s-compose-forward/service/compose"
	"os"
	"sort"
	"strings"
)

func printAppsSilently(path string) error {
	_ = os.Setenv("FORWARD_SILENT", "true")
	nameToService, err := compose.Load(path)
	if err != nil {
		return err
	}
	printValidApplications(nameToService)
	return nil
}

func printValidApplications(nameToService map[string]compose.Service) {
	result := make(sort.StringSlice, 0, len(nameToService))
	for key, _ := range nameToService {
		result = append(result, " - "+key)
	}
	result.Sort()

	fmt.Println(strings.Join(result, "\n"))
}
