package pkg

import (
	"bubblegit/utils"
	"fmt"

	// Replace with the actual module path to your main package
	tea "github.com/charmbracelet/bubbletea"
)

// Get keypresses and update the strings
func CloneRepo(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				if m.RepoName == "" {
					m.StatusMessage = "Link cannot be blank"
					m.State = "status"
				} else {
					m.Cursor++
				}
			case 1:
				output := utils.RunCommand("git", "clone", m.BranchName)
				m.StatusMessage = output
				m.State = "status"
				m.State = "menu"
				m.Cursor = 0
				m.BranchName = ""
				m.CreateBranch = false
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 2 {
				m.Cursor++
			}
		case "ctrl+c":
			m.State = "menu"
			m.BranchName = ""
		case "backspace":
			if m.Cursor == 0 && len(m.BranchName) > 0 {
				m.BranchName = m.BranchName[:len(m.BranchName)-1]
			}
		default:
			if m.Cursor == 0 {
				m.BranchName += keyMsg.String()
			}
		}
	}
	return m, nil
}

// Print the add menu on the screen
func ShowCloneRepo(m utils.Model) string {
	s := "Enter git repo link\n\n"
	branchChoices := []string{
		fmt.Sprintf("Git clone URL: %s", m.BranchName),
		"[Clone repo]",
	}

	for i, choice := range branchChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to go back, press [enter] to confirm or toggle true/false.\n"
	return s
}
