package pkg

import (
	"bubblegit/utils"
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

/*
gh pr create --base [target] 	   --head [source]
gh pr create --base my-base-branch --head my-changed-branch
*/

/*
1. create
2. list
3. status
4. checkout
5. view
6. approve
8. close
7. merge
8. reopen
9. delete
*/

func PullRequestSubmenu(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				// Create PR
				m.State = "createPR"
				m.Cursor = 0
			case 1:
				// List
				output := utils.RunCommand("gh", "pr", "ls")
				m.StatusMessage = output
				m.State = "status"
				m.Cursor = 0
			case 2:
				// Status
				output := utils.RunCommand("gh", "pr", "status")
				m.StatusMessage = output
				m.State = "status"
				m.Cursor = 0
			case 3:
				// Checkout
				m.State = "checkoutPR"
				m.Cursor = 0
			case 4:
				// View PR
				m.State = "viewPR"
				m.Cursor = 0
			case 5:
				// Approve
				m.State = "approvePR"
				m.Cursor = 0
			case 6:
				m.Cursor++
			case 7:
				m.Cursor++
			case 8:
				m.Cursor++
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 9 {
				m.Cursor++
			}
		case "ctrl+c", "q":
			m.State = "menu"
			m.BranchName = ""
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

func ShowPullRequestSubmenu(m utils.Model) string {
	prChoices := []string{
		"1. Create a pull request",
		"2. List pull requests",
		"3. Check pull request status",
		"4. Checkout a pull request",
		"5. View a pull request",
		"6. Approve a pull request",
		"7. Close a pull request",
		"8. Merge a pull request",
		"9. Reopen a pull request",
		"10. Delete a pull request",
	}
	btmMsg := "Press [q] or [ctrl+c] to go back to the main menu"
	return utils.ShowMenu(m, "Pull request", prChoices, btmMsg)
}

func CreatePR(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				m.Cursor++
			case 1:
				m.Cursor++
			case 2:
				m.Cursor++
			case 3:
				m.Cursor++
			case 4:
				if m.Source == "" {
					m.Source = strings.TrimSpace(utils.RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD"))
				}
				if m.Target == "" {
					m.Target = "main"
				}
				output := ""
				if m.BodyMessage == "" && m.Title == "" {
					output = utils.RunCommand("gh", "pr", "create", "-B", m.Target, "-H", m.Source)
				} else if m.BodyMessage == "" && m.Title != "" {
					output = utils.RunCommand("gh", "pr", "create", "-B", m.Target, "-H", m.Source, "--title", m.Title)
				} else if m.Title == "" && m.BodyMessage != "" {
					output = utils.RunCommand("gh", "pr", "create", "-B", m.Target, "-H", m.Source, "--body", m.BodyMessage)
				} else {
					output = utils.RunCommand("gh", "pr", "create", "-B", m.Target, "-H", m.Source, "--title", m.Title, "--body", m.BodyMessage)
				}
				m.StatusMessage = output
				m.State = "status"
				m.Source = ""
				m.Target = ""
				m.Target = ""
				m.BodyMessage = ""
				m.Cursor = 0

			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 9 {
				m.Cursor++
			}
		case "ctrl+c":
			m.State = "menu"
			m.BranchName = ""
		case "backspace":
			switch m.Cursor {
			case 0:
				if len(m.Source) > 0 {
					m.Source = m.Source[:len(m.Source)-1]
				}
			case 1:
				if len(m.Target) > 0 {
					m.Target = m.Target[:len(m.Target)-1]
				}
			case 2:
				if len(m.Title) > 0 {
					m.Title = m.Title[:len(m.Title)-1]
				}
			case 3:
				if len(m.BodyMessage) > 0 {
					m.BodyMessage = m.BodyMessage[:len(m.BodyMessage)-1]
				}
			}
		case "ctrl+s":
			m.Source = strings.TrimSpace(utils.RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD"))
			m.Target = "main"
			output := utils.RunCommand("gh", "pr", "create", "-B", m.Target, "-H", m.Source)
			m.StatusMessage = output
			m.State = "status"
			m.Source = ""
			m.Target = ""
			m.BodyMessage = ""
			m.Cursor = 0

		default:
			switch m.Cursor {
			case 0:
				m.Source += keyMsg.String()
			case 1:
				m.Target += keyMsg.String()
			case 2:
				m.Title += keyMsg.String()
			case 3:
				m.BodyMessage += keyMsg.String()
			}
		}
	}
	return m, nil
}

func ShowCreatePR(m utils.Model) string {
	createChoices := []string{
		fmt.Sprintf("Source branch (blank for current): %s", m.Source),
		fmt.Sprintf("Target branch (blank for main): %s", m.Target),
		fmt.Sprintf("Title (blank for default): %s", m.Title),
		fmt.Sprintf("Body message (can be empty): %s", m.BodyMessage),
		fmt.Sprintf("[PR %s > %s]", m.Source, m.Target),
	}
	topMsg := "Create a pull request"
	btmMsg := "Press [ctrl+c] to go back to the main menu, [ctrl+s] to quick PR"
	return utils.ShowMenu(m, topMsg, createChoices, btmMsg)
}

func CheckoutPR(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			output := ""
			if m.ID == "" {
				m.StatusMessage = "ID can't be blank"
				m.State = "status"
				m.ID = ""
				break
			} else if _, err := strconv.Atoi(m.ID); err != nil {
				m.StatusMessage = "ID has to be a valid number"
				m.State = "status"
				m.ID = ""
				break
			} else {
				output = utils.RunCommand("gh", "pr", "checkout", m.ID)
			}
			m.StatusMessage = output
			m.State = "status"
			m.ID = ""
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 9 {
				m.Cursor++
			}
		case "ctrl+c", "q":
			m.State = "menu"
			m.ID = ""
		case "backspace":
			if len(m.ID) > 0 {
				m.ID = m.ID[:len(m.ID)-1]
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			switch m.Cursor {
			case 0:
				m.ID += keyMsg.String()
			}
		}
	}
	return m, nil
}

func ShowCheckoutPR(m utils.Model) string {
	createChoices := []string{
		fmt.Sprintf("PR ID: %s", m.ID),
	}
	topMsg := "Checkout a PR"
	btmMsg := "Press [ctrl+c] to go back to the main menu, [enter] to checkout"
	return utils.ShowMenu(m, topMsg, createChoices, btmMsg)
}

func ViewPR(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			output := ""
			if m.ID == "" {
				m.StatusMessage = "ID can't be blank"
				m.State = "status"
				m.ID = ""
				break
			} else if _, err := strconv.Atoi(m.ID); err != nil {
				m.StatusMessage = "ID has to be a valid number"
				m.State = "status"
				m.ID = ""
				break
			} else {
				output = utils.RunCommand("gh", "pr", "view", m.ID)
			}
			m.StatusMessage = output
			m.State = "status"
			m.ID = ""
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 9 {
				m.Cursor++
			}
		case "ctrl+c", "q":
			m.State = "menu"
			m.ID = ""
		case "backspace":
			if len(m.ID) > 0 {
				m.ID = m.ID[:len(m.ID)-1]
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			switch m.Cursor {
			case 0:
				m.ID += keyMsg.String()
			}
		}
	}
	return m, nil
}

func ShowViewPR(m utils.Model) string {
	createChoices := []string{
		fmt.Sprintf("PR ID: %s", m.ID),
	}
	topMsg := "View PR"
	btmMsg := "Press [ctrl+c] to go back to the main menu, [enter] to view"
	return utils.ShowMenu(m, topMsg, createChoices, btmMsg)
}

func ApprovePR(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			output := ""
			if m.ID == "" {
				m.StatusMessage = "ID can't be blank"
				m.State = "status"
				m.ID = ""
				break
			} else if _, err := strconv.Atoi(m.ID); err != nil {
				m.StatusMessage = "ID has to be a valid number"
				m.State = "status"
				m.ID = ""
				break
			} else {
				output = utils.RunCommand("gh", "pr", "review", m.ID, "--approve")
			}
			m.StatusMessage = output
			m.State = "status"
			m.ID = ""
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 9 {
				m.Cursor++
			}
		case "ctrl+c", "q":
			m.State = "menu"
			m.ID = ""
		case "backspace":
			if len(m.ID) > 0 {
				m.ID = m.ID[:len(m.ID)-1]
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			switch m.Cursor {
			case 0:
				m.ID += keyMsg.String()
			}
		}
	}
	return m, nil
}

func ShowApprovePR(m utils.Model) string {
	createChoices := []string{
		fmt.Sprintf("PR ID: %s", m.ID),
	}
	topMsg := "Approve a PR"
	btmMsg := "Press [ctrl+c] to go back to the main menu, [enter] to approve"
	return utils.ShowMenu(m, topMsg, createChoices, btmMsg)
}
