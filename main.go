package main

import (
	"fmt"
	"log"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices       []string
	cursor        int
	selected      map[int]struct{}
	isTypingMsg   bool
	commitMessage string
	commitDesc    string
	state         string
	statusMessage string
}

func initialModel() model {
	return model{
		choices:  []string{"Add", "Commit", "Push", "Init"},
		selected: make(map[int]struct{}),
		state:    "menu", // default state
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

///////////////////////////////////
/////////// MENU ////////////////
///////////////////////////////////

// Get keypresses and update the cursor
func menuFunctions(m model, msg tea.Msg) (model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			switch m.cursor {
			case 0:
				m.state = "add"
			case 1:
				m.isTypingMsg = true
				m.state = "commitMessage"
			case 2:
				runGitCommand("git", "push")
				m.statusMessage = "Pushed to remote."
				m.state = "status"
			case 3:
				output := runGitCommand("git", "init")
				m.statusMessage = output
				m.state = "status"
			}
		}
	}
	return m, nil
}

// Print the menu on the screen
func showMenu(m model) string {
	s := "What would you like to do?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}

///////////////////////////////////
/////////// COMMIT ////////////////
///////////////////////////////////

// Get keypresses and update the commit message
func typeCommitMessage(m model, msg tea.Msg) (model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			if m.state == "commitMessage" {
				if m.commitMessage == "" {
					m.statusMessage = "Commit message cannot be empty!"
					m.state = "status"
					return m, nil
				}
				m.state = "commitDesc"
			} else if m.state == "commitDesc" {
				output := runGitCommand("git", "commit", "-m", m.commitMessage, "-m", m.commitDesc)
				m.statusMessage = output
				m.state = "status"
				m.commitMessage = ""
				m.commitDesc = ""
			}
		case "backspace":
			if m.state == "commitMessage" && len(m.commitMessage) > 0 {
				m.commitMessage = m.commitMessage[:len(m.commitMessage)-1]
			} else if m.state == "commitDesc" && len(m.commitDesc) > 0 {
				m.commitDesc = m.commitDesc[:len(m.commitDesc)-1]
			}
		case "ctrl+c", "q":
			m.state = "menu"
			m.commitMessage = ""
			m.commitDesc = ""
		default:
			if m.state == "commitMessage" {
				m.commitMessage += keyMsg.String()
			} else if m.state == "commitDesc" {
				m.commitDesc += keyMsg.String()
			}
		}
	}
	return m, nil
}

// /////// RUN GIT COMMAND //////////
func runGitCommand(name string, args ...string) string {
	cmd := exec.Command(name, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error: %s\n%s", err, output)
	}
	return string(output)
}

///////////////////////////////////
/////////// ADD ///////////////////
///////////////////////////////////

// Handle keypresses for the add menu
func add(m model, msg tea.Msg) (model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.cursor {
			case 0:
				runGitCommand("git", "add", ".")
				m.state = "menu"
			case 1:
				m.state = "addFile"
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < 1 {
				m.cursor++
			}
		case "ctrl+c", "q":
			m.state = "menu"
		}
	}
	return m, nil
}

// Get keypresses and update the file name to add
func addFile(m model, msg tea.Msg) (model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			runGitCommand("git", "add", m.commitMessage)
			m.state = "menu"
			m.commitMessage = ""
		case "ctrl+c", "q":
			m.state = "menu"
			m.commitMessage = ""
		default:
			m.commitMessage += keyMsg.String()
		}
	}
	return m, nil
}

// Print the add menu on the screen
func showAddMenu(m model) string {
	s := "What would you like to add?\n\n"
	addChoices := []string{"All files", "Specific file"}

	for i, choice := range addChoices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to cancel.\n"
	return s
}

///////////////////////////////////
/////////// UPDATE ////////////////
///////////////////////////////////

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.state {
	case "menu":
		m, cmd = menuFunctions(m, msg)
	case "commitMessage":
		m, cmd = typeCommitMessage(m, msg)
	case "add":
		m, cmd = add(m, msg)
	case "addFile":
		m, cmd = addFile(m, msg)
	case "status":
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			if keyMsg.String() == "enter" || keyMsg.String() == "q" {
				m.state = "menu"
			}
		}
	}

	return m, cmd
}

///////////////////////////////////
/////////// VIEW ////////////////
///////////////////////////////////

func (m model) View() string {
	switch m.state {
	case "menu":
		return showMenu(m)
	case "commitMessage":
		return fmt.Sprintf("Enter commit message: %s\n\nPress enter to add description or q to cancel.\n", m.commitMessage)
	case "commitDesc":
		return fmt.Sprintf("Enter commit description: %s\n\nPress enter to commit or q to cancel.\n", m.commitDesc)
	case "add":
		return showAddMenu(m)
	case "addFile":
		return fmt.Sprintf("Enter file name to add: %s\n\nPress enter to add or q to cancel.\n", m.commitMessage)
	case "status":
		return fmt.Sprintf("%s\n\nPress enter to return to menu.", m.statusMessage)
	}
	return ""
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
