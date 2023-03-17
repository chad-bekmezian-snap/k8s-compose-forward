package service

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// kubectl get service --field-selector metadata.namespace!=default -A -o=json
func k8sServices() (k8sSvcs, error) {
	cmd := exec.Command("kubectl", "get", "service", "-A", "-o=json", "--field-selector=metadata.namespace!=default")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(out)
		return nil, err
	}

	var result k8sGetServices
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, err
	}

	return result.Services, nil
}

type k8sSvc struct {
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

type k8sSvcs []k8sSvc

func (s k8sSvcs) FindServiceByClosestMatchingName(v string, namespace ...string) (k8sSvc, bool) {
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
		return k8sSvc{}, false
	}

	return s[currentMatchIndex], true
}

type k8sGetServices struct {
	Services k8sSvcs `json:"items"`
}
