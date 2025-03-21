package pkg

import (
	"bubblegit/utils"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func PullRequest(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				m.Cursor++
			case 1:
				m.Cursor++
			case 2:
				if m.BranchName != "" {
					m.BranchName = "main"
				}
				if m.OldBranchName == "" {
					m.OldBranchName = strings.TrimSpace(utils.RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD"))
				}
				output := utils.RunCommand("git", "checkout", m.BranchName)
				output += utils.RunCommand("git", "merge", m.OldBranchName, m.BranchName)
				m.StatusMessage = output + "\n\n merged " + m.OldBranchName + " into " + m.BranchName
				m.State = "status"
				m.Cursor = 0
				m.BranchName = ""
				m.OldBranchName = ""
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
			if len(m.BranchName) > 0 {
				m.BranchName = m.BranchName[:len(m.BranchName)-1]
			}
		default:
			switch m.Cursor {
			case 0:
				m.OldBranchName += keyMsg.String()
			case 1:
				m.BranchName += keyMsg.String()
			}
		}
	}
	return m, nil
}

func ShowPullRequest(m utils.Model) string {
	s := "Pull request\n\n"
	branchChoices := []string{
		fmt.Sprintf("Source branch (blank for current): %s", m.OldBranchName),
		fmt.Sprintf("Target branch (blank for main): %s", m.BranchName),
		"[Merge branches]",
	}
	for i, choice := range branchChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to go back, press [enter] to confirm.\n"
	return s
}
