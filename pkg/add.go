package pkg

import (
	"bubblegit/utils"
	"fmt"

	// Replace with the actual module path to your main package
	tea "github.com/charmbracelet/bubbletea"
)

///////////////////////////////////
/////////// ADD ///////////////////
///////////////////////////////////

// Handle keypresses for the add menu
func Add(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				utils.RunCommand("git", "add", ".")
				m.State = "menu"
			case 1:
				m.State = "addFile"
			}
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < 1 {
				m.Cursor++
			}
		case "ctrl+c", "q":
			m.State = "menu"
		}
	}
	return m, nil
}

// Get keypresses and update the file name to add
func AddFile(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			utils.RunCommand("git", "add", m.CommitMessage)
			m.CommitMessage = ""
		case "ctrl+c":
			m.State = "menu"
			m.CommitMessage = ""
		case "backspace":
			if len(m.CommitMessage) > 0 {
				m.CommitMessage = m.CommitMessage[:len(m.CommitMessage)-1]
			}
		default:
			m.CommitMessage += keyMsg.String()
		}
	}
	return m, nil
}

// Print the add menu on the screen
func ShowAddMenu(m utils.Model) string {
	s := "What would you like to add?\n\n"
	addChoices := []string{"All files", "Specific file"}

	for i, choice := range addChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] or [q] to go back.\n"
	return s
}
