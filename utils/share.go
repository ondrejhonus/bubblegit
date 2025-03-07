package utils

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Exported Model (capitalized)
type Model struct {
	Choices       []string
	Cursor        int
	Selected      map[int]struct{}
	StatusMessage string
	State         string
	IsTypingMsg   bool // Commit
	CommitMessage string
	CommitDesc    string
	RepoName      string // Repo create
	RepoDesc      string
	IsPublic      bool
	Source        string
	CreateClone   bool
	BranchName    string // Checkout
	CreateBranch  bool
}

// Exported function to create a new model
func InitialModel() Model {
	return Model{
		Choices:     []string{"Add", "Commit", "Push", "Branch", "Init", "Create repo"},
		Selected:    make(map[int]struct{}),
		State:       "menu", // default state
		IsPublic:    true,
		CreateClone: true,
	}
}

// Exported Init function
func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return ""
}
