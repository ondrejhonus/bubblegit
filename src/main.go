package main

import (
	"bubblegit/pkg"
	"bubblegit/utils"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// type model struct {
// 	choices       []string
// 	cursor        int
// 	selected      map[int]struct{}
// 	statusMessage string
// 	state         string
// 	isTypingMsg   bool // Commit
// 	commitMessage string
// 	commitDesc    string
// 	repoName      string // Repo create
// 	repoDesc      string
// 	isPublic      bool
// 	source        string
// 	createClone   bool
// }

// func initialModel() model {
// 	return model{
// 		choices:     []string{"Add", "Commit", "Push", "Init", "Create repo"},
// 		selected:    make(map[int]struct{}),
// 		state:       "menu", // default state
// 		isPublic:    true,
// 		createClone: true,
// 	}
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func utils.ShowStatus(m model, msg string) model {
// 	m.StatusMessage = msg
// 	m.State = "status"
// 	return m
// }

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

// // /////// RUN GIT COMMAND //////////
// func utils.RunCommand(name string, args ...string) string {
// 	cmd := exec.Command(name, args...)

// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return fmt.Sprintf("Error: %s\n%s", err, output)
// 	}
// 	return string(output)
// }

// ///////////////////////////////////
// /////////// ADD ///////////////////
// ///////////////////////////////////

// // Handle keypresses for the add menu
// func add(m model, msg tea.Msg) (model, tea.Cmd) {
// 	if keyMsg, ok := msg.(tea.KeyMsg); ok {
// 		switch keyMsg.String() {
// 		case "enter":
// 			switch m.Cursor {
// 			case 0:
// 				utils.RunCommand("git", "add", ".")
// 				m.State = "menu"
// 			case 1:
// 				m.State = "addFile"
// 			}
// 		case "up", "k":
// 			if m.Cursor > 0 {
// 				m.Cursor--
// 			}
// 		case "down", "j":
// 			if m.Cursor < 1 {
// 				m.Cursor++
// 			}
// 		case "ctrl+c", "q":
// 			m.State = "menu"
// 		}
// 	}
// 	return m, nil
// }

// // Get keypresses and update the file name to add
// func addFile(m model, msg tea.Msg) (model, tea.Cmd) {
// 	if keyMsg, ok := msg.(tea.KeyMsg); ok {
// 		switch keyMsg.String() {
// 		case "enter":
// 			utils.RunCommand("git", "add", m.CommitMessage)
// 			m.State = "menu"
// 			m.CommitMessage = ""
// 		case "ctrl+c":
// 			m.State = "menu"
// 			m.CommitMessage = ""
// 		case "backspace":
// 			if len(m.CommitMessage) > 0 {
// 				m.CommitMessage = m.CommitMessage[:len(m.CommitMessage)-1]
// 			}
// 		default:
// 			m.CommitMessage += keyMsg.String()
// 		}
// 	}
// 	return m, nil
// }

// // Print the add menu on the screen
// func showAddMenu(m model) string {
// 	s := "What would you like to add?\n\n"
// 	addChoices := []string{"All files", "Specific file"}

// 	for i, choice := range addChoices {
// 		cursor := " "
// 		if m.Cursor == i {
// 			cursor = ">"
// 		}
// 		s += fmt.Sprintf("%s %s\n", cursor, choice)
// 	}

// 	s += "\nPress [ctrl+c] to cancel.\n"
// 	return s
// }

///////////////////////////////////
/////////// COMMIT ////////////////
///////////////////////////////////

// Get keypresses and update the commit message
func typeCommitMessage(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+s", "enter":
			// Commit when message is entered
			if m.CommitMessage == "" {
				m.StatusMessage = "Commit message cannot be empty!"
				m.State = "status"
				return m, nil
			}
			output := utils.RunCommand("git", "commit", "-m", m.CommitMessage)
			if m.CommitDesc != "" {
				output = utils.RunCommand("git", "commit", "-m", m.CommitMessage, "-m", m.CommitDesc)
			}
			m.StatusMessage = output
			m.State = "status"
			m.CommitMessage = ""
			m.CommitDesc = ""
		case "ctrl+d":
			m.State = "commitDesc"
		case "backspace":
			// Handle backspace for commit message
			if len(m.CommitMessage) > 0 {
				m.CommitMessage = m.CommitMessage[:len(m.CommitMessage)-1]
			}
			// Handle backspace for commit description
			if len(m.CommitDesc) > 0 {
				m.CommitDesc = m.CommitDesc[:len(m.CommitDesc)-1]
			}
		case "ctrl+c":
			// Handle exit to menu and clear both fields
			m.State = "menu"
			m.CommitMessage = ""
			m.CommitDesc = ""
		default:
			m.CommitMessage += keyMsg.String()
		}
	}
	return m, nil
}

// Get keypresses and update the commit description
func typeCommitDesc(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+s", "enter":
			// Commit when description is entered
			output := utils.RunCommand("git", "commit", "-m", m.CommitMessage, "-m", m.CommitDesc)
			m.StatusMessage = output
			m.State = "status"
			m.CommitMessage = ""
			m.CommitDesc = ""
		case "backspace":
			// Handle backspace for commit description
			if len(m.CommitDesc) > 0 {
				m.CommitDesc = m.CommitDesc[:len(m.CommitDesc)-1]
			}
		case "ctrl+c":
			// Handle exit to menu and clear both fields
			m.State = "menu"
			m.CommitMessage = ""
			m.CommitDesc = ""
		default:
			// Append input to commit description
			m.CommitDesc += keyMsg.String()
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
func repoCreate(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
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

func showCreateRepoMenu(m utils.Model) string {
	s := "What would you want to do?\n\n"
	createChoices := []string{"Create repo from ./", "Create empty remote"}
	for i, choice := range createChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress [ctrl+c] to cancel, press [enter] to confirm.\n"
	return s
}

// Get keypresses and update the file name to add
func fromLocal(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
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

func createEmpty(m utils.Model, msg tea.Msg) (utils.Model, tea.Cmd) {
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

func showCreateFromLocal(m utils.Model) string {
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

func showCreateEmpty(m utils.Model) string {
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
		m.Model, cmd = typeCommitMessage(m.Model, msg)
	case "commitDesc":
		m.Model, cmd = typeCommitDesc(m.Model, msg)
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
		m.Model, cmd = repoCreate(m.Model, msg)
	case "fromLocal":
		m.Model, cmd = fromLocal(m.Model, msg)
	case "createEmpty":
		m.Model, cmd = createEmpty(m.Model, msg)
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
		return showCreateRepoMenu(m.Model)
	case "fromLocal":
		return showCreateFromLocal(m.Model)
	case "createEmpty":
		return showCreateEmpty(m.Model)
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
