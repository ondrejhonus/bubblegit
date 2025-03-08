package pkg

import (
	"bubblegit/utils"

	tea "github.com/charmbracelet/bubbletea"
)

///////////////////////////////////
/////////// COMMIT ////////////////
///////////////////////////////////

// Get keypresses and update the commit message
func TypeCommitMessage(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+s", "enter":
			// Commit when message is entered
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
			// Handle backspace for commit message
			if len(m.CommitMessage) > 0 {
				m.CommitMessage = m.CommitMessage[:len(m.CommitMessage)-1]
			}
			// Handle backspace for commit description
			if len(m.CommitDesc) > 0 {
				m.CommitDesc = m.CommitDesc[:len(m.CommitDesc)-1]
			}
		case "ctrl+c":
			// Handle exit to menu and clear both fields
			m.State = "menu"
			m.CommitMessage = ""
			m.CommitDesc = ""
		default:
			m.CommitMessage += keyMsg.String()
		}
	}
	return m, nil
}

// Get keypresses and update the commit description
func TypeCommitDesc(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+s", "enter":
			// Commit when description is entered
			output := utils.RunCommand("git", "commit", "-m", m.CommitMessage, "-m", m.CommitDesc)
			m.StatusMessage = output
			m.State = "status"
			m.CommitDesc = ""
		case "backspace":
			// Handle backspace for commit description
			if len(m.CommitDesc) > 0 {
				m.CommitDesc = m.CommitDesc[:len(m.CommitDesc)-1]
			}
		case "ctrl+c":
			// Handle exit to menu and clear both fields
			m.State = "menu"
			m.CommitMessage = ""
			m.CommitDesc = ""
		default:
			// Append input to commit description
			m.CommitDesc += keyMsg.String()
		}
	}
	return m, nil
}
