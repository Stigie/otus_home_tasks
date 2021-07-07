package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	cmdExec := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	envs := []string{}
	for key, value := range env {
		if _, ok := os.LookupEnv(key); ok {
			os.Unsetenv(key)
		}
		if value != "" {
			envs = append(envs, key+"="+value)
		}
	}
	cmdExec.Env = append(os.Environ(), envs...)
	cmdExec.Stdin = os.Stdin
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	err := cmdExec.Run()
	if err != nil {
		return 1
	}
	return cmdExec.ProcessState.ExitCode()
}
