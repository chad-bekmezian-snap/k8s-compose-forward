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

		*f = append(*f, v)
	}
	return nil
}

var (
	omitFlag     multiValueFlag
	appFlag      multiValueFlag
	serviceFlag  multiValueFlag
	yamlPathFlag string
)

func init() {
	flag.Var(&omitFlag, "omit", "Can be used to omit services by the given name(s) from port-forwarding")
	flag.Var(&omitFlag, "o", "Can be used to omit services by the given name(s) from port-forwarding")

	flag.Var(&appFlag, "app", "Used to be specify one or more applications for which to start port-forwarding")
	flag.Var(&appFlag, "a", "Used to be specify one or more applications for which to start port-forwarding")

	flag.StringVar(&yamlPathFlag, "file", "docker-compose.yml", "Used to be specify one or more services for which to start port-forwarding")
	flag.StringVar(&yamlPathFlag, "f", "docker-compose.yml", "Used to be specify one or more services for which to start port-forwarding")

	flag.Parse()

	for i, v := range omitFlag {
		v = strings.TrimSpace(v)
		omitFlag[i] = v
	}

	for i, v := range serviceFlag {
		v = strings.TrimSpace(v)
		omitFlag[i] = v
	}
}
