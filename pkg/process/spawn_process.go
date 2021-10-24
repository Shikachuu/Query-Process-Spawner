package process

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func createSubCommand(cmd []string, msg string) *exec.Cmd {
	subCmd := exec.Command(cmd[0], cmd[1:]...)
	subCmd.Stdin = strings.NewReader(msg)
	subCmd.Stdout = os.Stdout
	subCmd.Stderr = os.Stderr
	subCmd.Env = os.Environ()
	return subCmd
}

func RunProcesses(qmc chan string, cmd []string, mw int) {
	select {
	case msg := <-qmc:
		if len(qmc) < mw {
			go func() {
				subCmd := createSubCommand(cmd, msg)
				err := subCmd.Start()
				if err != nil {
					log.Fatalf("failed to start command: %s", err)
				}
				err = subCmd.Wait()
				if err != nil {
					log.Fatalf("an error occurred while running command: %s", err)
				}
			}()
		} else {
			RunProcesses(qmc, cmd, mw)
		}
	}
}
