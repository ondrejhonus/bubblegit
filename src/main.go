package main

import (
	"bubblegit/pkg"
	"bubblegit/utils"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

///////////////////////////////////
/////////// MENU ////////////////
///////////////////////////////////

// Get keypresses and update the cursor
func menuFunctions(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "enter":
			switch m.Cursor {
			case 0:
				// Add
				m.State = "add"
				m.Cursor = 0
			case 1:
				// Commit
				m.IsTypingMsg = true
				m.State = "commitMessage"
				m.Cursor = 0
			case 2:
				// Push

				output := utils.RunCommand("git", "push")
				m = utils.ShowStatus(m, output)
				utils.ShowStatus(m, output)
			case 3:
				output := utils.RunCommand("git", "init")
				utils.ShowStatus(m, output)
			case 4:
				m.State = "createRepo"
				m.Cursor = 0
			}
		}
	}
	return m, nil
}

// Print the menu on the screen
func showMenu(m utils.Model) string {
	s := "What would you like to do?\n\n"

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [q] to quit.\n"
	return s
}

///////////////////////////////////
/////////// UPDATE ////////////////
///////////////////////////////////

type localModel struct {
	utils.Model
}

// Implement Bubble Tea's Update function on localModel
func (m localModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.State {
	case "menu":
		m.Model, cmd = menuFunctions(m.Model, msg) // FIX: Use `=` instead of `:=`
	case "commitMessage":
		m.Model, cmd = pkg.TypeCommitMessage(m.Model, msg)
	case "commitDesc":
		m.Model, cmd = pkg.TypeCommitDesc(m.Model, msg)
	case "add":
		m.Model, cmd = pkg.Add(m.Model, msg)
	case "addFile":
		m.Model, cmd = pkg.AddFile(m.Model, msg)
	case "status":
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			if keyMsg.String() == "enter" || keyMsg.String() == "q" {
				m.State = "menu"
			}
		}
	case "createRepo":
		m.Model, cmd = pkg.RepoCreate(m.Model, msg)
	case "fromLocal":
		m.Model, cmd = pkg.FromLocal(m.Model, msg)
	case "createEmpty":
		m.Model, cmd = pkg.CreateEmpty(m.Model, msg)
	}

	return m, cmd
}

// Implement Bubble Tea's View function on localModel
func (m localModel) View() string {
	switch m.State {
	case "menu":
		return showMenu(m.Model)
	case "commitMessage":
		return fmt.Sprintf("Enter commit message: %s\n\nPress [enter] to commit, [ctrl+d] to add description or [ctrl+c] to cancel.\n", m.CommitMessage)
	case "commitDesc":
		return fmt.Sprintf("Enter commit description: %s\n\nPress [enter] to commit or [ctrl+c] to cancel.\n", m.CommitDesc)
	case "add":
		return pkg.ShowAddMenu(m.Model)
	case "addFile":
		return fmt.Sprintf("Enter file name to add: %s\n\nPress [enter] to add or [ctrl+c] to cancel.\n", m.CommitMessage)
	case "status":
		return fmt.Sprintf("%s\n\nPress [enter] to return to menu.", m.StatusMessage)
	case "createRepo":
		return pkg.ShowCreateRepoMenu(m.Model)
	case "fromLocal":
		return pkg.ShowCreateFromLocal(m.Model)
	case "createEmpty":
		return pkg.ShowCreateEmpty(m.Model)
	}

	return ""
}

func main() {
	// FIX: Wrap utils.InitialModel() in localModel
	p := tea.NewProgram(localModel{Model: utils.InitialModel()})

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
