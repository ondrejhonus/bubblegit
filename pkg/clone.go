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
			case 3:
				if m.RepoName == "" {
					m.StatusMessage = "Link can't be empty!"
					m.State = "status"
				} else {
					output := ""
					if m.CloneDepth == "" && m.Source == "" {
						output = utils.RunCommand("git", "clone", m.RepoName)
					} else if utils.IsPositiveInteger(m.CloneDepth) && m.Source == "" {
						output = utils.RunCommand("git", "clone", "--depth ", m.CloneDepth, m.RepoName)
					} else if utils.IsPositiveInteger(m.CloneDepth) && m.Source != "" {
						output = utils.RunCommand("git", "clone", "--depth ", m.CloneDepth, m.RepoName, m.Source)
					} else if m.Source != "" && m.CloneDepth == "" {
						output = utils.RunCommand("git", "clone", m.RepoName, m.Source)
					} else {
						m.StatusMessage = "Invalid depth. Please enter a positive integer."
						m.State = "status"
					}
					m.StatusMessage += (output + "\nSuccesfully cloned repository " + m.RepoName + " into " + m.Source)
					m.State = "status"
					m.State = "menu"
					m.Cursor = 0
					m.Source = ""
					m.CloneDepth = ""
					m.RepoName = ""
					m.CreateBranch = false
				}
			default:
				m.Cursor++
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 3 {
				m.Cursor++
			}
		case "ctrl+c":
			m.State = "menu"
			m.RepoName = ""
		case "backspace":
			if m.Cursor == 0 && len(m.RepoName) > 0 {
				m.RepoName = m.RepoName[:len(m.RepoName)-1]
			} else if m.Cursor == 1 && len(m.Source) > 0 {
				m.Source = m.Source[:len(m.Source)-1]
			} else if m.Cursor == 2 && len(m.CloneDepth) > 0 {
				m.CloneDepth = m.CloneDepth[:len(m.CloneDepth)-1]
			}
		default:
			if m.Cursor == 0 {
				m.RepoName += keyMsg.String()
			}
			if m.Cursor == 2 {
				if utils.IsPositiveInteger(keyMsg.String()) || (len(m.CloneDepth) > 0 && keyMsg.String() == "0") {
					m.CloneDepth += keyMsg.String()
				}
			}
		}
	}
	return m, nil
}

func ShowCloneRepo(m utils.Model) string {
	s := "Enter git repo link"
	choices := []string{
		fmt.Sprintf("Git clone URL: %s", m.RepoName),
		fmt.Sprintf("Directory (\"./\"): %s", m.Source),
		fmt.Sprintf("Depth -> int (\"\"): %s", m.CloneDepth),
		"[Clone repo]",
	}
	btmMsg := "\nPress [ctrl+c] to go back, press [enter] to confirm or toggle true/false.\n"
	return utils.ShowMenu(m, s, choices, btmMsg)
}
