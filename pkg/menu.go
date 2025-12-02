package pkg

import (
	"bubblegit/utils"

	tea "github.com/charmbracelet/bubbletea"
)

///////////////////////////////////
/////////// MENU ////////////////
///////////////////////////////////

// Get keypresses and update the cursor
func MenuFunctions(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "enter":
			switch m.Cursor {
			case 0:
				// Add
				m.State = "add"
				m.Cursor = 0
			case 1:
				// Commit
				m.IsTypingMsg = true
				m.State = "commitMessage"
				m.Cursor = 0
			case 2:
				// Push
				output := utils.RunCommand("git", "push")
				m.StatusMessage = output
				m.State = "status"
			case 3:
				// Clone
				m.State = "clone"
				m.Cursor = 0
			case 4:
				// List
				m.State = "list"
				m.Cursor = 0
			case 5:
				// Branch
				m.State = "branches"
				m.Cursor = 0
			case 6:
				// Pull Request
				m.State = "pullRequest"
				m.Cursor = 0
			case 7:
				output := utils.RunCommand("git", "init")
				m.StatusMessage = output
				m.State = "status"
			case 8:
				m.State = "createRepo"
				m.Cursor = 0
			}
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if len(keyMsg.String()) == 1 {
				num := int(keyMsg.String()[0] - '1')
				if num >= 0 && num < 9 {
					m.Cursor = num
				}
			}
		}
	}
	return m, nil
}

// Print the menu on the screen
func ShowMenu(m utils.Model) string {
	s := "What would you like to do?"
	btmMsg := "\nPress [ctrl+c] or [q] to go back.\n"
	return utils.ShowMenu(m, s, m.Choices, btmMsg)
}
