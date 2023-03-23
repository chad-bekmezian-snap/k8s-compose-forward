package k8s

import (
	"encoding/json"
	"fmt"
	"os"
)

// ListServices kubectl get service --field-selector metadata.namespace!=default -A -o=json
func ListServices() (ServiceSlice, error) {
	cmd := executeCMD("kubectl", "get", "service", "-A", "-o=json", "--field-selector=metadata.namespace!=default")
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
