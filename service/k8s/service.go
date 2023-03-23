package k8s

import "strings"

type ServiceSlice []Service

func (s ServiceSlice) FindServiceByClosestMatchingName(v string, namespace ...string) (Service, bool) {
	currentMatchIndex := -1

ServiceLoop:
	for i, svc := range s {
		for _, ns := range namespace {
			// Skip this iteration if namespace doesn't match provided
			if ns == svc.Detail.Namespace {
				break
			}
			continue ServiceLoop
		}

		if strings.Contains(svc.Detail.Name, v) {
			if currentMatchIndex > -1 && len(svc.Detail.Name) >= len(s[currentMatchIndex].Detail.Name) {
				continue
			}
			currentMatchIndex = i
		}
	}

	if currentMatchIndex == -1 {
		return Service{}, false
	}

	return s[currentMatchIndex], true
}

type Service struct {
	Kind   string `json:"kind"`
	Detail struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
		Uid             string `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		Ports []struct {
			Port     int    `json:"port"`
			Protocol string `json:"protocol"`
		} `json:"ports"`
	} `json:"spec"`
}
