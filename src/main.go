package main

import (
	"bubblegit/pkg"
	"bubblegit/utils"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

///////////////////////////////////
/////////// UPDATE ////////////////
///////////////////////////////////

// Create a local model based on Model from utils
type localModel struct {
	utils.Model
}

func (m localModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.State {
	case "menu":
		m.Model, cmd = pkg.MenuFunctions(m.Model, msg) // FIX: Use `=` instead of `:=`
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
	case "branches":
		m.Model, cmd = pkg.BranchControl(m.Model, msg)
	case "checkoutBranch":
		m.Model, cmd = pkg.CheckoutBranch(m.Model, msg)
	case "setUpstream":
		m.Model, cmd = pkg.SetUpstream(m.Model, msg)
	case "deleteBranch":
		m.Model, cmd = pkg.DeleteBranch(m.Model, msg)
	case "renameBranch":
		m.Model, cmd = pkg.RenameBranch(m.Model, msg)
	case "mergeBranch":
		m.Model, cmd = pkg.MergeBranch(m.Model, msg)
	case "rebaseBranch":
		m.Model, cmd = pkg.RebaseBranch(m.Model, msg)
	case "clone":
		m.Model, cmd = pkg.CloneRepo(m.Model, msg)
	}

	return m, cmd
}

func (m localModel) View() string {
	switch m.State {
	case "menu":
		return pkg.ShowMenu(m.Model)
	case "commitMessage":
		return fmt.Sprintf("Enter commit message: %s\n\nPress [enter] to commit, [ctrl+d] to add description or [ctrl+c] to cancel.\n", m.CommitMessage)
	case "commitDesc":
		return fmt.Sprintf("Enter commit description: %s\n\nPress [enter] to commit or [ctrl+c] to cancel.\n", m.CommitDesc)
	case "add":
		return pkg.ShowAddMenu(m.Model)
	case "addFile":
		return fmt.Sprintf("Enter file name to add: %s\n\nPress [enter] to add or [ctrl+c] to cancel.\n", m.CommitMessage)
	case "status":
		return fmt.Sprintf("%s\n\nPress [enter]/[q] to close.", m.StatusMessage)
	case "createRepo":
		return pkg.ShowCreateRepoMenu(m.Model)
	case "fromLocal":
		return pkg.ShowCreateFromLocal(m.Model)
	case "createEmpty":
		return pkg.ShowCreateEmpty(m.Model)
	case "branches":
		return pkg.ShowBranchesMenu(m.Model)
	case "checkoutBranch":
		return pkg.ShowCheckoutBranch(m.Model)
	case "setUpstream":
		return pkg.ShowSetUpstream(m.Model)
	case "deleteBranch":
		return pkg.ShowDeleteBranch(m.Model)
	case "renameBranch":
		return pkg.ShowRenameBranch(m.Model)
	case "mergeBranch":
		return pkg.ShowMergeBranch(m.Model)
	case "rebaseBranch":
		return pkg.ShowRebaseBranch(m.Model)
	case "clone":
		return pkg.ShowCloneRepo(m.Model)
	}

	return ""
}

func main() {
	p := tea.NewProgram(localModel{Model: utils.InitialModel()}, tea.WithoutBracketedPaste())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
