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

func ShowMenu(m Model, title string, choices []string, top string) string {
	s := title + "\n\n"
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
