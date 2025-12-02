package pkg

import (
	"bubblegit/utils"
	"fmt"
	"strings"

	// Replace with the actual module path to your main package
	tea "github.com/charmbracelet/bubbletea"
)

///////////////////////////////////
/////////// BRANCHES //////////////
///////////////////////////////////

// Handle keypresses for the checkout menu
func BranchControl(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				m.State = "checkoutBranch"
				m.Cursor = 0
			case 1:
				m.State = "setUpstream"
				m.Cursor = 0
			case 2:
				m.State = "deleteBranch"
				m.Cursor = 0
			case 3:
				m.State = "renameBranch"
				m.Cursor = 0
			case 4:
				m.State = "mergeBranch"
				m.Cursor = 0
			case 5:
				m.State = "rebaseBranch"
				m.Cursor = 0
			}
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < 5 {
				m.Cursor++
			}
		case "ctrl+c", "q":
			m.State = "menu"

		case "1", "2", "3", "4", "5", "6":
			if len(keyMsg.String()) == 1 {
				num := int(keyMsg.String()[0] - '1')
				if num >= 0 && num < 6 {
					m.Cursor = num
				}
			}
		}
	}
	return m, nil
}

// Print the branches menu on the screen
func ShowBranchesMenu(m utils.Model) string {
	s := "Branches"
	choices := []string{
		"1 | Checkout branch",
		"2 | Set upstream",
		"3 | Delete branch",
		"4 | Rename branch",
		"5 | Merge branch",
		"6 | Rebase branch",
	}
	btmMsg := "Press [q] or [ctrl+c] to go back to the main menu"
	return utils.ShowMenu(m, s, choices, btmMsg)
}

// Get keypresses and update the strings
func CheckoutBranch(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				if m.BranchName == "" {
					m.StatusMessage = "Branch name cannot be empty"
					m.State = "status"
				} else {
					m.Cursor++
				}
			case 1:
				m.CreateBranch = !m.CreateBranch
				m.Cursor++
			case 2:
				if m.CreateBranch {
					utils.RunCommand("git", "stash")
					output := utils.RunCommand("git", "checkout", "-b", m.BranchName)
					m.StatusMessage = output
					m.State = "status"

				} else {
					utils.RunCommand("git", "checkout", m.BranchName)
				}
				m.State = "menu"
				m.Cursor = 0
				m.BranchName = ""
				m.CreateBranch = false
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 2 {
				m.Cursor++
			}
		case "ctrl+c":
			m.State = "branches"
			m.BranchName = ""
			m.Cursor = 0
		case "backspace":
			if m.Cursor == 0 && len(m.BranchName) > 0 {
				m.BranchName = m.BranchName[:len(m.BranchName)-1]
			}
		default:
			if m.Cursor == 0 {
				m.BranchName += keyMsg.String()
			}
		}
	}
	return m, nil
}

// Print the add menu on the screen
func ShowCheckoutBranch(m utils.Model) string {
	s := "Configure branch checkout"
	choices := []string{
		fmt.Sprintf("Branch name: %s", m.BranchName),
		fmt.Sprintf("Create branch: %t", m.CreateBranch),
		"[Checkout/Stash branch]",
	}
	btmMsg := "\nPress [ctrl+c] to go back, press [enter] to confirm.\n"
	return utils.ShowMenu(m, s, choices, btmMsg)
}

func SetUpstream(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			if m.BranchName == "" {
				currentBranch := strings.TrimSpace(utils.RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD"))
				output := utils.RunCommand("git", "push", "--set-upstream", "origin", currentBranch)
				m.StatusMessage = output
				m.State = "status"

				m.Cursor = 0
				m.BranchName = ""
			} else {
				output := utils.RunCommand("git", "push", "--set-upstream", "origin", m.BranchName)
				m.StatusMessage = output
				m.State = "status"
				m.Cursor = 0
				m.BranchName = ""
			}
		case "ctrl+c":
			m.State = "branches"
			m.BranchName = ""
			m.Cursor = 0
		case "backspace":
			if len(m.BranchName) > 0 {
				m.BranchName = m.BranchName[:len(m.BranchName)-1]
			}
		default:
			m.BranchName += keyMsg.String()
		}
	}
	return m, nil
}

func ShowSetUpstream(m utils.Model) string {
	s := "Set upstream"
	choices := []string{
		fmt.Sprintf("Branch name (blank for current): %s", m.BranchName),
	}
	btmMsg := "\nPress [ctrl+c] to go back, press [enter] to confirm.\n"
	return utils.ShowMenu(m, s, choices, btmMsg)
}

func DeleteBranch(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				m.Cursor++
			case 1:
				if m.BranchName != "" {
					output := utils.RunCommand("git", "branch", "-D", m.BranchName)
					m.StatusMessage = output
					m.State = "status"
					m.Cursor = 0
					m.BranchName = ""
				} else {
					m.StatusMessage = "Branch name cannot be empty"
					m.State = "status"
				}
			}
		case "ctrl+c":
			m.State = "branches"
			m.BranchName = ""
			m.Cursor = 0
		case "backspace":
			if len(m.BranchName) > 0 {
				m.BranchName = m.BranchName[:len(m.BranchName)-1]
			}
		default:
			switch m.Cursor {
			case 0:
				m.BranchName += keyMsg.String()
			}
		}
	}
	return m, nil
}

func ShowDeleteBranch(m utils.Model) string {
	s := "Delete branch"
	choices := []string{
		fmt.Sprintf("Branch name: %s", m.BranchName),
		"[Delete branch]",
	}
	btmMsg := "\nPress [ctrl+c] to go back, press [enter] to confirm.\n"
	return utils.ShowMenu(m, s, choices, btmMsg)
}

func RenameBranch(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				m.Cursor++
			case 1:
				m.Cursor++
			case 2:
				if m.BranchName != "" {
					if m.OldBranchName == "" {
						m.OldBranchName = strings.TrimSpace(utils.RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD"))
					}
					output := utils.RunCommand("git", "branch", "-m", m.OldBranchName, m.BranchName)
					m.StatusMessage = output + "\n\n renamed " + m.OldBranchName + " to " + m.BranchName
					m.State = "status"
					m.Cursor = 0
					m.BranchName = ""
					m.OldBranchName = ""
				} else {
					m.StatusMessage = "Branch name cannot be empty"
					m.State = "status"
					m.Cursor = 0
				}
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 2 {
				m.Cursor++
			}
		case "ctrl+c":
			m.State = "branches"
			m.BranchName = ""
			m.Cursor = 0
		case "backspace":
			if len(m.BranchName) > 0 {
				m.BranchName = m.BranchName[:len(m.BranchName)-1]
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

func ShowRenameBranch(m utils.Model) string {
	s := "Rename branch"
	choices := []string{
		fmt.Sprintf("Branch name (blank for current): %s", m.OldBranchName),
		fmt.Sprintf("New branch name: %s", m.BranchName),
		"[Rename branch]",
	}
	btmMsg := "\nPress [ctrl+c] to go back, press [enter] to confirm.\n"
	return utils.ShowMenu(m, s, choices, btmMsg)
}

func MergeBranch(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				m.Cursor++
			case 1:
				m.Cursor++
			case 2:
				if m.BranchName != "" {
					if m.OldBranchName == "" {
						m.OldBranchName = strings.TrimSpace(utils.RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD"))
					}
					output := utils.RunCommand("git", "checkout", m.BranchName)
					output += utils.RunCommand("git", "merge", m.OldBranchName, m.BranchName)
					m.StatusMessage = output + "\n\n merged " + m.OldBranchName + " into " + m.BranchName
					m.State = "status"
					m.Cursor = 0
					m.BranchName = ""
					m.OldBranchName = ""
				} else {
					m.StatusMessage = "Branch name cannot be empty"
					m.State = "status"
					m.Cursor = 0
				}
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 2 {
				m.Cursor++
			}
		case "ctrl+c":
			m.State = "branches"
			m.BranchName = ""
			m.Cursor = 0
		case "backspace":
			if len(m.BranchName) > 0 {
				m.BranchName = m.BranchName[:len(m.BranchName)-1]
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

func ShowMergeBranch(m utils.Model) string {
	s := "Merge branch"
	choices := []string{
		fmt.Sprintf("Source branch (blank for current): %s", m.OldBranchName),
		fmt.Sprintf("Target branch: %s", m.BranchName),
		"[Merge branches]",
	}
	btmMsg := "\nPress [ctrl+c] to go back, press [enter] to confirm.\n"
	return utils.ShowMenu(m, s, choices, btmMsg)
}

func RebaseBranch(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				m.Cursor++
			case 1:
				m.Cursor++
			case 2:
				if m.BranchName != "" {
					if m.OldBranchName == "" {
						m.OldBranchName = strings.TrimSpace(utils.RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD"))
					}
					output := utils.RunCommand("git", "rebase", m.OldBranchName, m.BranchName)
					m.StatusMessage = output + "\n\n rebased " + m.OldBranchName + " onto " + m.BranchName
					m.State = "status"
					m.Cursor = 0
					m.BranchName = ""
					m.OldBranchName = ""
				} else {
					m.StatusMessage = "Branch name cannot be empty"
					m.State = "status"
					m.Cursor = 0
				}
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 2 {
				m.Cursor++
			}
		case "ctrl+c":
			m.State = "branches"
			m.BranchName = ""
			m.Cursor = 0
		case "backspace":
			if len(m.BranchName) > 0 {
				m.BranchName = m.BranchName[:len(m.BranchName)-1]
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

func ShowRebaseBranch(m utils.Model) string {
	s := "Rebase branch"
	choices := []string{
		fmt.Sprintf("Source branch (blank for current): %s", m.OldBranchName),
		fmt.Sprintf("Target branch: %s", m.BranchName),
		"[Rebase branches]",
	}
	btmMsg := "\nPress [ctrl+c] to go back, press [enter] to confirm.\n"
	return utils.ShowMenu(m, s, choices, btmMsg)
}
