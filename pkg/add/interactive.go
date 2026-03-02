package add

import (
	"github.com/ondrejhonus/bubblegit/utils"
	
	tea "github.com/charmbracelet/bubbletea"
)

func InteractiveAdd(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	output := utils.RunCommand("git", "add", "-p")
				m.Viewport.SetContent(output)
				m.Viewport.GotoTop()
				m.State = "InteractiveAdd"
				m.Cursor = 0
				return m, nil
}