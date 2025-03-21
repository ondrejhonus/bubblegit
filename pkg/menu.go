package pkg

import (
	"bubblegit/utils"
	"fmt"

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
				// Branch
				m.State = "branches"
				m.Cursor = 0
			case 5:
				output := utils.RunCommand("git", "init")
				m.StatusMessage = output
				m.State = "status"
			case 6:
				m.State = "createRepo"
				m.Cursor = 0
			}
		}
	}
	return m, nil
}

// Print the menu on the screen
func ShowMenu(m utils.Model) string {
	s := "What would you like to do?\n\n"

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] or [q] to go back.\n"
	return s
}
