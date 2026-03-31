package pkg

import (
	"github.com/ondrejhonus/bubblegit/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func ListMenu(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				// Short Commit list
				output := utils.RunCommand("git", "log", "--oneline", "--graph", "--decorate", "--all")
				m.StatusMessage = "All commits in brief graph view:\n"
				m.StatusMessage += output
				m.State = "status"
				m.Cursor = 0
			case 1:
				// list branches
				output := utils.RunCommand("git", "branch", "-a")
				m.StatusMessage = "All branches:\n"
				m.StatusMessage += output
				m.State = "status"
				m.Cursor = 0
			case 2:
				// branch diff
				output := utils.RunCommand("git", "diff", "--color=always")
				m.Viewport.SetContent(output)
				m.Viewport.GotoTop()
				m.State = "diff"
				m.Cursor = 0
			case 3:
				// stash list
				output := utils.RunCommand("git", "stash", "list")
				m.StatusMessage = "All stashes:\n"
				m.StatusMessage += output
				m.State = "status"
				m.Cursor = 0
			case 4:
				// tag list
				output := utils.RunCommand("git", "tag")
				m.StatusMessage = "All tags:\n"
				m.StatusMessage += output
				m.State = "status"
				m.Cursor = 0
			case 5:
				// remote list
				output := utils.RunCommand("git", "remote", "-v")
				m.StatusMessage = "All remotes:\n"
				m.StatusMessage += output
				m.State = "status"
				m.Cursor = 0
			case 6:
				// config list
				output := utils.RunCommand("git", "config", "--list")
				m.StatusMessage = "All configurations:\n"
				m.StatusMessage += output
				m.State = "status"
				m.Cursor = 0
			case 7:
				// tracked files
				output := utils.RunCommand("git", "ls-files")
				m.StatusMessage = "All tracked files:\n"
				m.StatusMessage += output
				m.State = "status"
				m.Cursor = 0
			case 8:
				// untracked
				output := utils.RunCommand("git", "ls-files", "--others", "--exclude-standard")
				m.StatusMessage = "All untracked files:\n"
				m.StatusMessage += output
				m.State = "status"
				m.Cursor = 0
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 8 {
				m.Cursor++
			}
		case "ctrl+c", "q":
			m.State = "menu"
			m.BranchName = ""
			m.Cursor = 0
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if len(keyMsg.String()) == 1 {
				num := int(keyMsg.String()[0] - '1')
				if num >= 0 && num < 9 {
					m.Cursor = num
				}
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

func ShowListMenu(m utils.Model) string {
	prChoices := []string{
		"1 | Commits",
		"2 | Branches",
		"3 | Diff",
		"4 | Stashes",
		"5 | Tags",
		"6 | Remotes",
		"7 | Configs",
		"8 | Tracked Files",
		"9 | Untracked Files",
	}
	btmMsg := "Press [q] or [ctrl+c] to go back to the main menu"
	return utils.ShowMenu(m, "Show list of", prChoices, btmMsg)
}
