package pkg

import (
	"bubblegit/utils"

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
			case 2:
				m.State = "unaddFile"
			case 3:
				utils.RunCommand("git", "add", ".")
				m.State = "menu"
				m.Cursor = 0
			}
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < 3 {
				m.Cursor++
			}
		case "ctrl+c", "q":
			m.State = "menu"
			m.Cursor = 0
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			if len(keyMsg.String()) == 1 {
				num := int(keyMsg.String()[0] - '1')
				if num >= 0 && num < 8 {
					m.Cursor = num
				}
			}
		}
	}
	return m, nil
}

// Get keypresses and update the file name to add
func AddFile(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			utils.RunCommand("git", "add", m.FileName)
			m.FileName = ""
		case "ctrl+c":
			m.State = "add"
			m.FileName = ""
			m.Cursor = 0
		case "backspace":
			if len(m.FileName) > 0 {
				m.FileName = m.FileName[:len(m.FileName)-1]
			}
		default:
			m.FileName += keyMsg.String()
		}
	}
	return m, nil
}

// Print the add menu on the screen
func ShowAddMenu(m utils.Model) string {
	s := "What would you like to add?"
	choices := []string{"1 | All files", "2 | Add file", "3 | Un-add file", "4 | Reset added"}
	btmMsg := "Press [q] or [ctrl+c] to go back to the main menu"
	return utils.ShowMenu(m, s, choices, btmMsg)
}

// Get keypresses and update the file name to unstage
func UnaddFile(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			utils.RunCommand("git", "reset", m.FileName)
			m.FileName = ""
		case "ctrl+c":
			m.State = "add"
			m.FileName = ""
			m.Cursor = 0
		case "backspace":
			if len(m.FileName) > 0 {
				m.FileName = m.FileName[:len(m.FileName)-1]
			}
		default:
			m.FileName += keyMsg.String()
		}
	}
	return m, nil
}
