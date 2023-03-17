package main

import (
	"fmt"
	color "github.com/TwiN/go-color"
	"os"
	"os/exec"
	"strings"
	"sync"
	"test/fdcas/cmd/k9s-automation/app"
	"test/fdcas/cmd/k9s-automation/service"
)

var currentProcesses = sync.Map{}
var appNameToApp = map[string]app.App{
	"ovrc-experience":   app.OvrCExperience,
	"customer-process":  app.CustomerProcess,
	"mobile-experience": app.MobileExperience,
	"license-process":   app.MobileExperience,
}

func main() {
	if len(appFlag) == 0 {
		fmt.Println("No applications specified. Exiting.")
	}

	var wg sync.WaitGroup
	for _, appName := range appFlag {
		services, ok := appNameToApp[appName]
		if !ok {
			fmt.Println(color.Ize(color.Red, "ERROR: Undefined application with name: "+appName))
			fmt.Println("Valid values are:")
			printValidApplications()
			return
		}

		for _, s := range services {
			if omitFlag.Contains(s.Service.Name) {
				fmt.Println(color.Ize(color.Blue, "Skipping "+s.Service.Name))
				continue
			}
			wg.Add(1)
			go func(s service.Configurable) {
				portForwardToService(s)
				wg.Done()
			}(s)
		}
	}
	wg.Wait()
}

func portForwardToService(s service.Configurable) {
	var port uint
	switch {
	case s.DefaultPort > 0:
		port = s.DefaultPort
	default:
		port = s.Service.DefaultPort
	}

	// Default to port 80 if no port is specified for forwarding
	if s.Service.ForwardToPort == 0 {
		s.Service.ForwardToPort = 80
	}

	fmt.Println(fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' starting", s.Service.Name, s.Service.Namespace, port, s.Service.ForwardToPort))

	cmdArgs := []string{
		fmt.Sprintf("-n=%s", s.Service.Namespace),
		"port-forward", s.Service.Name,
		fmt.Sprintf("%d:%d", port, s.Service.ForwardToPort),
	}

	if alreadyRunning := cmdRunning("kubectl", cmdArgs...); alreadyRunning {
		fmt.Println(color.Ize(color.Yellow, fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' is already in progress.", s.Service.Name, s.Service.Namespace, port, s.Service.ForwardToPort)))
		return
	}

	// e.g. kubectl -n=experience-services port-forward service/cese-ovrc-experience-k8s 8085:80
	cmd := exec.Command("kubectl", cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Println(color.Ize(color.Red, fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' failed to start.", s.Service.Name, s.Service.Namespace, port, s.Service.ForwardToPort)))
		return
	}

	fmt.Println(color.Ize(color.Green, fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' started", s.Service.Name, s.Service.Namespace, port, s.Service.ForwardToPort)))

	if err := cmd.Wait(); err != nil {
		fmt.Println(color.Ize(color.Red, fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' failed.", s.Service.Name, s.Service.Namespace, port, s.Service.ForwardToPort)))
		fmt.Println(err)
		return
	}
}

func cmdRunning(name string, args ...string) bool {
	for i, arg := range args {
		args[i] = strings.TrimSpace(arg)
	}

	if _, existed := currentProcesses.LoadOrStore(strings.Join(args, ""), nil); existed {
		return true
	}

	psCmd := exec.Command("ps", "aux")
	grep := exec.Command("grep", fmt.Sprintf(`%s %s`, strings.TrimSpace(name), strings.Join(args, " ")))
	removeGrep := exec.Command("grep", "-v", "grep")

	grepPipe, _ := psCmd.StdoutPipe()
	defer grepPipe.Close()
	removeGrepPipe, _ := grep.StdoutPipe()
	defer removeGrepPipe.Close()
	grep.Stdin = grepPipe
	removeGrep.Stdin = removeGrepPipe

	_ = psCmd.Start()
	_ = grep.Start()
	o, _ := removeGrep.Output()

	return len(o) > 0
}

func printValidApplications() {
	result := make([]string, 0, len(appNameToApp))
	for key, _ := range appNameToApp {
		result = append(result, " - "+key)
	}

	fmt.Println(strings.Join(result, "\n"))
}
