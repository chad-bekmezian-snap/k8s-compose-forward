package forward

import (
	"fmt"
	"github.com/TwiN/go-color"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

var currentProcesses sync.Map

func ToService(s Service) {
	printStatus(color.Green, "starting.", s)

	cmdArgs := []string{
		fmt.Sprintf("-n=%s", s.Namespace()),
		"port-forward", s.Name(),
		fmt.Sprintf("%d:%d", s.FromPort(), s.ToPort()),
	}

	if alreadyRunning := cmdRunning("kubectl", cmdArgs...); alreadyRunning {
		printStatus(color.Yellow, "already in progress.", s)
		return
	}

	cmd := exec.Command("kubectl", cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		printStatus(color.Red, "failed", s)
		fmt.Println(err)
		return
	}

	printStatus(color.Green, "started", s)
	if err := cmd.Wait(); err != nil {
		printStatus(color.Red, "failed", s)
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

	isDuplicateProcess := func() bool {
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

	if runtime.GOOS == "windows" {
		isDuplicateProcess = func() bool {
			psCmd := exec.Command("ps", "aux")
			selectString := exec.Command("Select-String", fmt.Sprintf(`%s %s`, strings.TrimSpace(name), strings.Join(args, " ")))

			selectStringPipe, _ := psCmd.StdoutPipe()
			defer selectStringPipe.Close()
			selectString.Stdin = selectStringPipe

			_ = psCmd.Start()
			o, _ := selectString.Output()

			return len(o) > 0
		}
	}

	return isDuplicateProcess()
}

func printStatus(c, status string, s Service) {
	fmt.Println(color.Ize(c, fmt.Sprintf("Port forwarding '%s' in namespace '%s' with ports '%d:%d' %s.", s.Name(), s.Namespace(), s.FromPort(), s.ToPort(), status)))
}
