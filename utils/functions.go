package utils

import (
	"fmt"
	"os/exec"
)

// /////// RUN GIT COMMAND //////////
func RunCommand(name string, args ...string) string {
	cmd := exec.Command(name, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error: %s\n%s", err, output)
	}
	return string(output)
}

// Show status messages after running a command
func ShowStatus(m Model, msg string) Model {
	m.StatusMessage = msg
	m.State = "status"
	return m
}
