package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/chad-bekmezian-snap/cs-port-forwarding/compose/service"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var currentProcesses = sync.Map{}

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
				portForwardToService(s)
				wg.Done()
			}(depSvc)
		}
	}
	wg.Wait()
}

func portForwardToService(s service.Service) {
	var fromPort int
	for _, p := range s.Ports {
		ports := strings.Split(p, ":")

		if len(ports) == 1 && strconv.Itoa(s.K8sPort) == ports[0] {
			break
		}

		if len(ports) == 2 {
			if ports[1] != strconv.Itoa(s.K8sPort) {
				continue
			}

			fromPort, _ = strconv.Atoi(ports[0])
			break
		}

	}

	if fromPort == 0 {
		fmt.Println(color.Ize(color.Red, "Failed to find a designated port to use when forwarding."))
		panic("AH")
	}

	fmt.Println(fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' starting", s.K8sName, s.K8sNamespace, fromPort, s.K8sPort))

	cmdArgs := []string{
		fmt.Sprintf("-n=%s", s.K8sNamespace),
		"port-forward", "service/" + s.K8sName,
		fmt.Sprintf("%d:%d", fromPort, s.K8sPort),
	}

	if alreadyRunning := cmdRunning("kubectl", cmdArgs...); alreadyRunning {
		fmt.Println(color.Ize(color.Yellow, fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' is already in progress.", s.K8sName, s.K8sNamespace, fromPort, s.K8sPort)))
		return
	}

	// e.g. kubectl -n=experience-services port-forward service/cese-ovrc-experience-k8s 8085:80
	cmd := exec.Command("kubectl", cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Println(color.Ize(color.Red, fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' failed to start.", s.K8sName, s.K8sNamespace, fromPort, s.K8sPort)))
		return
	}

	fmt.Println(color.Ize(color.Green, fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' started", s.K8sName, s.K8sNamespace, fromPort, s.K8sPort)))

	if err := cmd.Wait(); err != nil {
		fmt.Println(color.Ize(color.Red, fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' failed.", s.K8sName, s.K8sNamespace, fromPort, s.K8sPort)))
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

func printValidApplications(nameToService map[string]service.Service) {
	result := make(sort.StringSlice, 0, len(nameToService))
	for key, _ := range nameToService {
		result = append(result, " - "+key)
	}
	result.Sort()

	fmt.Println(strings.Join(result, "\n"))
}
