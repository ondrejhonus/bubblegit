package pkg

import (
	"bubblegit/utils"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

///////////////////////////////////
/////////// CREATE REPO ///////////
///////////////////////////////////

// Create repo menu 1
func RepoCreate(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				m.State = "fromLocal"
				m.Cursor = 0
			case 1:
				m.State = "createEmpty"
				m.Cursor = 0
			}
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < 2 {
				m.Cursor++
			}
		case "ctrl+c", "q":
			m.State = "menu"
		}
	}
	return m, nil
}

func ShowCreateRepoMenu(m utils.Model) string {
	s := "What would you want to do?\n\n"
	createChoices := []string{"Create repo from ./", "Create empty remote"}
	for i, choice := range createChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] or [q] to go back, [enter] to confirm.\n"
	return s
}

// Get keypresses and update the file name to add
func FromLocal(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				// Repo name
				if m.RepoName == "" {
					m.RepoName = "bubblegit-repo"
				}
				m.Cursor++
			case 1:
				// Repo description
				m.Cursor++
			case 2:
				// Source
				if m.Source == "" {
					m.Source = "."
				}
				m.Cursor++
			case 3:
				// Public
				m.IsPublic = !m.IsPublic
				m.Cursor++
			case 4:
				// Create repo
				var visibility string
				if m.IsPublic {
					visibility = "--public"
				} else {
					visibility = "--private"
				}
				utils.RunCommand("gh", "repo", "create", m.RepoName, "--description", m.RepoDesc, visibility, "--source", m.Source)
				m.State = "menu"
				m.RepoName = ""
				m.RepoDesc = ""
				m.Source = ""
				m.IsPublic = false
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 4 {
				m.Cursor++
			}
		case "ctrl+c":
			m.State = "menu"
			m.RepoName = ""
			m.RepoDesc = ""
			m.Source = ""
			m.IsPublic = false

		case "backspace":
			switch m.Cursor {
			case 0:
				if len(m.RepoName) > 0 {
					m.RepoName = m.RepoName[:len(m.RepoName)-1]
				}
			case 1:
				if len(m.RepoDesc) > 0 {
					m.RepoDesc = m.RepoDesc[:len(m.RepoDesc)-1]
				}
			case 2:
				if len(m.Source) > 0 {
					m.Source = m.Source[:len(m.Source)-1]
				}
			}
		default:
			switch m.Cursor {
			case 0:
				m.RepoName += keyMsg.String()
			case 1:
				m.RepoDesc += keyMsg.String()
			case 2:
				m.Source += keyMsg.String()
			}
		}
	}
	return m, nil
}

func CreateEmpty(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.Cursor {
			case 0:
				// Repo name
				if m.RepoName == "" {
					m.RepoName = "bubblegit-repo"
				}
				m.Cursor++
			case 1:
				// Repo description
				m.Cursor++
			case 2:
				// Public?
				m.IsPublic = !m.IsPublic
				m.Cursor++
			case 3:
				// Clone?
				m.CreateClone = !m.CreateClone
				m.Cursor++
			case 4:
				// Create repo
				var visibility string
				if m.IsPublic {
					visibility = "--public"
				} else {
					visibility = "--private"
				}
				var clone string
				if m.CreateClone {
					clone = "--clone"
				} else {
					clone = ""
				}

				// gh repo create <repo-name> --description "<repo-description>" --public --clone
				utils.RunCommand("gh", "repo", "create", m.RepoName, "--description", m.RepoDesc, visibility, clone)
				m.State = "menu"
				m.RepoName = ""
				m.RepoDesc = ""
				m.Source = ""
				m.IsPublic = false
				m.CreateClone = false
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "tab":
			if m.Cursor < 4 {
				m.Cursor++
			}
		case "ctrl+c":
			m.State = "menu"
			m.RepoName = ""
			m.RepoDesc = ""
			m.Source = ""
			m.IsPublic = false

		case "backspace":
			switch m.Cursor {
			case 0:
				if len(m.RepoName) > 0 {
					m.RepoName = m.RepoName[:len(m.RepoName)-1]
				}
			case 1:
				if len(m.RepoDesc) > 0 {
					m.RepoDesc = m.RepoDesc[:len(m.RepoDesc)-1]
				}
			case 2:
				if len(m.Source) > 0 {
					m.Source = m.Source[:len(m.Source)-1]
				}
			}
		default:
			switch m.Cursor {
			case 0:
				m.RepoName += keyMsg.String()
			case 1:
				m.RepoDesc += keyMsg.String()
			case 2:
				m.Source += keyMsg.String()
			}
		}
	}
	return m, nil
}

func ShowCreateFromLocal(m utils.Model) string {
	s := "Enter the following details:\n\n"
	createChoices := []string{
		fmt.Sprintf("Name: %s", m.RepoName),
		fmt.Sprintf("Description: %s", m.RepoDesc),
		fmt.Sprintf("Source (default = ./): %s", m.Source),
		fmt.Sprintf("Public: %t", m.IsPublic),
		"[Create repo]",
	}

	for i, choice := range createChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to cancel, press [enter] to toggle true/false.\n"
	return s
}

func ShowCreateEmpty(m utils.Model) string {
	s := "Enter the following details:\n\n"
	createChoices := []string{
		fmt.Sprintf("Name: %s", m.RepoName),
		fmt.Sprintf("Description: %s", m.RepoDesc),
		fmt.Sprintf("Public: %t", m.IsPublic),
		fmt.Sprintf("Clone: %t", m.CreateClone),
		"[Create repo]",
	}

	for i, choice := range createChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to cancel, press [enter] to toggle true/false.\n"
	return s
}
