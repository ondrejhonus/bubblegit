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
				m.State = "setUpstream"
			case 2:
				m.State = "deleteBranch"
			case 3:
				m.State = "renameBranch"
			case 4:
				m.State = "mergeBranch"
			case 5:
				m.State = "rebaseBranch"
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
		}
	}
	return m, nil
}

// Print the branches menu on the screen
func ShowBranchesMenu(m utils.Model) string {
	s := "Branches\n\n"
	branchChoices := []string{
		"Checkout branch",
		"Set upstream",
		"Delete branch",
		"Rename branch",
		"Merge branch",
		"Rebase branch",
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
func CheckoutBranch(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
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
					output := utils.RunCommand("git", "checkout", "-b", m.BranchName)
					utils.ShowStatus(m, output)

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
func ShowCheckoutBranch(m utils.Model) string {
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

func SetUpstream(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			if m.BranchName == "" {
				currentBranch := utils.RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
				output := utils.RunCommand("git", "branch", "--set-upstream-to", "origin/"+currentBranch)
				utils.ShowStatus(m, output)
				m.State = "menu"
				m.Cursor = 0
				m.BranchName = ""
				// testing a branch
			} else {
				output := utils.RunCommand("git", "branch", "--set-upstream-to", m.BranchName)
				utils.ShowStatus(m, output)
				m.State = "menu"
				m.Cursor = 0
				m.BranchName = ""
			}
		case "ctrl+c":
			m.State = "menu"
			m.BranchName = ""
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
	s := "Set upstream\n\n"
	branchChoices := []string{
		fmt.Sprintf("Branch name (blank for current): %s", m.BranchName),
	}
	for i, choice := range branchChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to cancel, press [enter] to confirm.\n"
	return s
}
