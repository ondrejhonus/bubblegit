package pkg

import (
	"github.com/ondrejhonus/bubblegit/utils"

	tea "github.com/charmbracelet/bubbletea"
)

// main add menu
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
		case "1", "2", "3", "4":
			if len(keyMsg.String()) == 1 {
				num := int(keyMsg.String()[0] - '1')
				if num >= 0 && num < 4 {
					m.Cursor = num
				}
			}
		}
	}
	return m, nil
}

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

func ShowAddMenu(m utils.Model) string {
	s := "What would you like to add?"
	choices := []string{"1 | All files", "2 | Add file", "3 | Un-add file", "4 | Reset added"}
	choices = append(choices)
	return utils.ShowMenu(m, s, choices, m.ExitMessage)
}
