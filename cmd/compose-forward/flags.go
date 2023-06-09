package main

import (
	"flag"
	"fmt"
	"strings"
)

type multiValueFlag []string

func (f *multiValueFlag) Contains(v string) bool {
	for _, n := range *f {
		if n == v {
			return true
		}
	}

	return false
}

func (f *multiValueFlag) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *multiValueFlag) Set(value string) error {
	vals := strings.Split(value, " ")

	for _, v := range vals {
		if f.Contains(v) {
			continue
		}

		*f = append(*f, strings.TrimSpace(v))
	}
	return nil
}

var (
	omitFlag     multiValueFlag
	appArgs      multiValueFlag
	serviceFlag  multiValueFlag
	yamlPathFlag string
	listServices bool
	completions  bool
)

func init() {
	flag.Var(&omitFlag, "omit", "Can be used to omit services by the given name(s) from port-forwarding")
	flag.Var(&omitFlag, "o", "Can be used to omit services by the given name(s) from port-forwarding")

	flag.Var(&serviceFlag, "service", "Used to be specify one or more services to start port-forwarding to.")
	flag.Var(&serviceFlag, "s", "Used to be specify one or more services to start port-forwarding to.")

	flag.StringVar(&yamlPathFlag, "file", "docker-compose.yml", "Used to be specify one or more services for which to start port-forwarding")
	flag.StringVar(&yamlPathFlag, "f", "docker-compose.yml", "Used to be specify one or more services for which to start port-forwarding")

	flag.BoolVar(&listServices, "list", false, "If provided, nothing will be run and service names will be listed.")
	flag.BoolVar(&listServices, "l", false, "If provided, nothing will be run and service names will be listed.")

	flag.BoolVar(&completions, "completions", false, "If provided, nothing will be run, and services names will be listed that can be used to seeds bash completions.")
	flag.BoolVar(&completions, "c", false, "If provided, nothing will be run, and services names will be listed that can be used to seeds bash completions.")

	flag.Parse()
	appArgs = flag.Args()
}
