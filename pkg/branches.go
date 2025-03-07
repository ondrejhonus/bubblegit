package pkg

import (
	"bubblegit/utils"
	"fmt"

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
			case 1:
				m.State = "deleteBranch"
			case 2:
				m.State = "renameBranch"
			case 3:
				m.State = "mergeBranch"
			case 4:
				m.State = "rebaseBranch"
			case 5:
				m.State = "setUpstream"
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

// Print the branches menu on the screen
func ShowBranchesMenu(m utils.Model) string {
	s := "Branches\n\n"
	branchChoices := []string{
		"Checkout branch",
		"Delete branch",
		"Rename branch",
		"Merge branch",
		"Rebase branch",
		"Set upstream",
	}

	for i, choice := range branchChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to go back.\n"
	return s
}

// Get keypresses and update the file name to add
func TypeCheckout(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				if m.BranchName == "" {
					utils.ShowStatus(m, "Branch name cannot be empty")
				} else {
					m.Cursor++
				}
			case 1:
				m.CreateBranch = !m.CreateBranch
				m.Cursor++
			case 2:
				if m.CreateBranch {
					utils.RunCommand("git", "checkout", "-b", m.BranchName)
				} else {
					utils.RunCommand("git", "checkout", m.BranchName)
				}
				m.State = "menu"
				m.Cursor = 0
				m.BranchName = ""
				m.CreateBranch = false
			}
		case "ctrl+c":
			m.State = "menu"
			m.CommitMessage = ""
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
func ShowCheckoutMenu(m utils.Model) string {
	s := "Configure branch checkout\n\n"
	branchChoices := []string{
		fmt.Sprintf("Branch name: %s", m.BranchName),
		fmt.Sprintf("Create branch: %t", m.CreateBranch),
		"[Checkout branch]",
	}

	for i, choice := range branchChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to cancel, press [enter] to confirm or change bool value.\n"
	return s
}
