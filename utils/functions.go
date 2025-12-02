package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
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

func ShowMenu(m Model, title string, choices []string, top string) string {
	s := title + "\n"
	for i, choice := range choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\n" + top + "\n"
	return s
}

func GetGitHubUsername() (string, error) {
	cmd := exec.Command("gh", "api", "user", "--jq", ".login")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}
