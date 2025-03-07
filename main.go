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
	statusMessage string
	state         string
	isTypingMsg   bool // Commit
	commitMessage string
	commitDesc    string
	repoName      string // Repo create
	repoDesc      string
	isPublic      bool
	source        string
	createClone   bool
}

func initialModel() model {
	return model{
		choices:     []string{"Add", "Commit", "Push", "Init", "Create repo"},
		selected:    make(map[int]struct{}),
		state:       "menu", // default state
		isPublic:    true,
		createClone: true,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func showStatus(m model, msg string) {
	m.statusMessage = msg
	m.state = "status"
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
				// Add
				m.state = "add"
				m.cursor = 0
			case 1:
				// Commit
				m.isTypingMsg = true
				m.state = "commitMessage"
				m.cursor = 0
			case 2:
				// Push
				m.cursor = 0
				showStatus(m, "Pushing to remote...")
				output := runCommand("git", "push")
				showStatus(m, output)
			case 3:
				output := runCommand("git", "init")
				showStatus(m, output)
			case 4:
				m.state = "createRepo"
				m.cursor = 0
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

	s += "\nPress [q] to quit.\n"
	return s
}

// /////// RUN GIT COMMAND //////////
func runCommand(name string, args ...string) string {
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
				runCommand("git", "add", ".")
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
			runCommand("git", "add", m.commitMessage)
			m.state = "menu"
			m.commitMessage = ""
		case "ctrl+c":
			m.state = "menu"
			m.commitMessage = ""
		case "backspace":
			if len(m.commitMessage) > 0 {
				m.commitMessage = m.commitMessage[:len(m.commitMessage)-1]
			}
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

	s += "\nPress [ctrl+c] to cancel.\n"
	return s
}

///////////////////////////////////
/////////// COMMIT ////////////////
///////////////////////////////////

// Get keypresses and update the commit message
func typeCommitMessage(m model, msg tea.Msg) (model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+s", "enter":
			// Commit when message is entered
			if m.commitMessage == "" {
				m.statusMessage = "Commit message cannot be empty!"
				m.state = "status"
				return m, nil
			}
			output := runCommand("git", "commit", "-m", m.commitMessage)
			if m.commitDesc != "" {
				output = runCommand("git", "commit", "-m", m.commitMessage, "-m", m.commitDesc)
			}
			m.statusMessage = output
			m.state = "status"
			m.commitMessage = ""
			m.commitDesc = ""
		case "ctrl+d":
			m.state = "commitDesc"
		case "backspace":
			// Handle backspace for commit message
			if len(m.commitMessage) > 0 {
				m.commitMessage = m.commitMessage[:len(m.commitMessage)-1]
			}
			// Handle backspace for commit description
			if len(m.commitDesc) > 0 {
				m.commitDesc = m.commitDesc[:len(m.commitDesc)-1]
			}
		case "ctrl+c":
			// Handle exit to menu and clear both fields
			m.state = "menu"
			m.commitMessage = ""
			m.commitDesc = ""
		default:
			m.commitMessage += keyMsg.String()
		}
	}
	return m, nil
}

// Get keypresses and update the commit description
func typeCommitDesc(m model, msg tea.Msg) (model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+s", "enter":
			// Commit when description is entered
			output := runCommand("git", "commit", "-m", m.commitMessage, "-m", m.commitDesc)
			m.statusMessage = output
			m.state = "status"
			m.commitMessage = ""
			m.commitDesc = ""
		case "backspace":
			// Handle backspace for commit description
			if len(m.commitDesc) > 0 {
				m.commitDesc = m.commitDesc[:len(m.commitDesc)-1]
			}
		case "ctrl+c":
			// Handle exit to menu and clear both fields
			m.state = "menu"
			m.commitMessage = ""
			m.commitDesc = ""
		default:
			// Append input to commit description
			m.commitDesc += keyMsg.String()
		}
	}
	return m, nil
}

///////////////////////////////////
/////////// CREATE REPO ///////////
///////////////////////////////////

/* Some gh commands ill use:
1. gh repo create <repo-name> --description "<repo-description>" --public --source .
2. gh repo create <repo-name> --description "<repo-description>" --private --source .
This one creates a local folder and makes a gh repo
3. gh repo create my-project --public --clone
*/

/*
gh repo create <repo-name> --description "<repo-description>" --? --source .
>	Create repo from ./
	>	Repo name: [Default: {wd}]
		Repo description:
		Source: [Default: .]
		[*] Public
		[ ] Readme
		[ ] .gitignore

gh repo create <repo-name> --description "<repo-description>" --public --clone
	Create empty remote and clone it
	>	Repo name:
		Repo description
		[*] Public
		[ ] Readme
		[ ] .gitignore
	Create empty remote repo
	>	Repo name:
		Repo description
		[*] Public
		[ ] Readme
		[ ] .gitignore
*/

// Create repo menu 1
func repoCreate(m model, msg tea.Msg) (model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.cursor {
			case 0:
				m.state = "fromLocal"
				m.cursor = 0
			case 1:
				m.state = "createEmpty"
				m.cursor = 0
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < 2 {
				m.cursor++
			}
		case "ctrl+c", "q":
			m.state = "menu"
		}
	}
	return m, nil
}

func showCreateRepoMenu(m model) string {
	s := "What would you want to do?\n\n"
	createChoices := []string{"Create repo from ./", "Create empty remote"}
	for i, choice := range createChoices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to cancel, press [enter] to confirm.\n"
	return s
}

// Get keypresses and update the file name to add
func fromLocal(m model, msg tea.Msg) (model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.cursor {
			case 0:
				// Repo name
				if m.repoName == "" {
					m.repoName = "bubblegit-repo"
				}
				m.cursor++
			case 1:
				// Repo description
				m.cursor++
			case 2:
				// Source
				if m.source == "" {
					m.source = "."
				}
				m.cursor++
			case 3:
				// Public
				m.isPublic = !m.isPublic
				m.cursor++
			case 4:
				// Create repo
				var visibility string
				if m.isPublic {
					visibility = "--public"
				} else {
					visibility = "--private"
				}
				runCommand("gh", "repo", "create", m.repoName, "--description", m.repoDesc, visibility, "--source", m.source)
				m.state = "menu"
				m.repoName = ""
				m.repoDesc = ""
				m.source = ""
				m.isPublic = false
			}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "tab":
			if m.cursor < 4 {
				m.cursor++
			}
		case "ctrl+c":
			m.state = "menu"
			m.repoName = ""
			m.repoDesc = ""
			m.source = ""
			m.isPublic = false

		case "backspace":
			switch m.cursor {
			case 0:
				if len(m.repoName) > 0 {
					m.repoName = m.repoName[:len(m.repoName)-1]
				}
			case 1:
				if len(m.repoDesc) > 0 {
					m.repoDesc = m.repoDesc[:len(m.repoDesc)-1]
				}
			case 2:
				if len(m.source) > 0 {
					m.source = m.source[:len(m.source)-1]
				}
			}
		default:
			switch m.cursor {
			case 0:
				m.repoName += keyMsg.String()
			case 1:
				m.repoDesc += keyMsg.String()
			case 2:
				m.source += keyMsg.String()
			}
		}
	}
	return m, nil
}

func createEmpty(m model, msg tea.Msg) (model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			switch m.cursor {
			case 0:
				// Repo name
				if m.repoName == "" {
					m.repoName = "bubblegit-repo"
				}
				m.cursor++
			case 1:
				// Repo description
				m.cursor++
			case 2:
				// Public?
				m.isPublic = !m.isPublic
				m.cursor++
			case 3:
				// Clone?
				m.createClone = !m.createClone
				m.cursor++
			case 4:
				// Create repo
				var visibility string
				if m.isPublic {
					visibility = "--public"
				} else {
					visibility = "--private"
				}
				var clone string
				if m.createClone {
					clone = "--clone"
				} else {
					clone = ""
				}

				// gh repo create <repo-name> --description "<repo-description>" --public --clone
				runCommand("gh", "repo", "create", m.repoName, "--description", m.repoDesc, visibility, clone)
				m.state = "menu"
				m.repoName = ""
				m.repoDesc = ""
				m.source = ""
				m.isPublic = false
				m.createClone = false
			}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "tab":
			if m.cursor < 4 {
				m.cursor++
			}
		case "ctrl+c":
			m.state = "menu"
			m.repoName = ""
			m.repoDesc = ""
			m.source = ""
			m.isPublic = false

		case "backspace":
			switch m.cursor {
			case 0:
				if len(m.repoName) > 0 {
					m.repoName = m.repoName[:len(m.repoName)-1]
				}
			case 1:
				if len(m.repoDesc) > 0 {
					m.repoDesc = m.repoDesc[:len(m.repoDesc)-1]
				}
			case 2:
				if len(m.source) > 0 {
					m.source = m.source[:len(m.source)-1]
				}
			}
		default:
			switch m.cursor {
			case 0:
				m.repoName += keyMsg.String()
			case 1:
				m.repoDesc += keyMsg.String()
			case 2:
				m.source += keyMsg.String()
			}
		}
	}
	return m, nil
}

func showCreateFromLocal(m model) string {
	s := "Enter the following details:\n\n"
	createChoices := []string{
		fmt.Sprintf("Name: %s", m.repoName),
		fmt.Sprintf("Description: %s", m.repoDesc),
		fmt.Sprintf("Source (default = ./): %s", m.source),
		fmt.Sprintf("Public: %t", m.isPublic),
		"[Create repo]",
	}

	for i, choice := range createChoices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to cancel, press [enter] to toggle true/false.\n"
	return s
}

func showCreateEmpty(m model) string {
	s := "Enter the following details:\n\n"
	createChoices := []string{
		fmt.Sprintf("Name: %s", m.repoName),
		fmt.Sprintf("Description: %s", m.repoDesc),
		fmt.Sprintf("Public: %t", m.isPublic),
		fmt.Sprintf("Clone: %t", m.createClone),
		"[Create repo]",
	}

	for i, choice := range createChoices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to cancel, press [enter] to toggle true/false.\n"
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
	case "commitDesc":
		m, cmd = typeCommitDesc(m, msg)
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
	case "createRepo":
		m, cmd = repoCreate(m, msg)
	case "fromLocal":
		m, cmd = fromLocal(m, msg)
	case "createEmpty":
		m, cmd = createEmpty(m, msg)
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
		return fmt.Sprintf("Enter commit message: %s\n\nPress [enter] to commit, [ctrl+d] to add description or [ctrl+c] to cancel.\n", m.commitMessage)
	case "commitDesc":
		return fmt.Sprintf("Enter commit description: %s\n\nPress [enter] to commit or [ctrl+c] to cancel.\n", m.commitDesc)
	case "add":
		return showAddMenu(m)
	case "addFile":
		return fmt.Sprintf("Enter file name to add: %s\n\nPress [enter] to add or [ctrl+c] to cancel.\n", m.commitMessage)
	case "status":
		return fmt.Sprintf("%s\n\nPress [enter] to return to menu.", m.statusMessage)
	case "createRepo":
		return showCreateRepoMenu(m)
	case "fromLocal":
		return showCreateFromLocal(m)
	case "createEmpty":
		return showCreateEmpty(m)
	}

	return ""
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
