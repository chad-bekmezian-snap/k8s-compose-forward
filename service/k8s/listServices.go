package k8s

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ListServices kubectl get service --field-selector metadata.namespace!=default -A -o=json
func ListServices() (ServiceSlice, error) {
	cmd := exec.Command("kubectl", "get", "service", "-A", "-o=json", "--field-selector=metadata.namespace!=default")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(out)
		return nil, err
	}

	var result struct {
		Services ServiceSlice `json:"items"`
	}
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, err
	}

	return result.Services, nil
}

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
