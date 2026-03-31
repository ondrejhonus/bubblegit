package pkg

import (
	"github.com/ondrejhonus/bubblegit/utils"

	tea "github.com/charmbracelet/bubbletea"
)


func TypeCommitMessage(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+s", "enter":
			if m.CommitMessage == "" {
				m.StatusMessage = "Commit message cannot be empty!"
				m.State = "status"
				return m, nil
			}
			output := utils.RunCommand("git", "commit", "-m", m.CommitMessage)
			if m.CommitDesc != "" {
				output = utils.RunCommand("git", "commit", "-m", m.CommitMessage, "-m", m.CommitDesc)
			}
			m.StatusMessage = output
			m.State = "status"
			m.CommitMessage = ""
			m.CommitDesc = ""
		case "ctrl+d":
			m.State = "commitDesc"
		case "backspace":
			if len(m.CommitMessage) > 0 {
				m.CommitMessage = m.CommitMessage[:len(m.CommitMessage)-1]
			}
			if len(m.CommitDesc) > 0 {
				m.CommitDesc = m.CommitDesc[:len(m.CommitDesc)-1]
			}
		case "ctrl+c":
			m.State = "menu"
			m.CommitMessage = ""
			m.CommitDesc = ""
		default:
			m.CommitMessage += keyMsg.String()
		}
	}
	return m, nil
}

func TypeCommitDesc(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+s", "enter":
			output := utils.RunCommand("git", "commit", "-m", m.CommitMessage, "-m", m.CommitDesc)
			m.StatusMessage = output
			m.State = "status"
			m.CommitDesc = ""
		case "backspace":
			if len(m.CommitDesc) > 0 {
				m.CommitDesc = m.CommitDesc[:len(m.CommitDesc)-1]
			}
		case "ctrl+c":
			m.State = "menu"
			m.CommitMessage = ""
			m.CommitDesc = ""
		default:
			m.CommitDesc += keyMsg.String()
		}
	}
	return m, nil
}
