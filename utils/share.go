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
	// Commit
	IsTypingMsg   bool
	CommitMessage string
	CommitDesc    string
	// Add
	FileName string
	// Repo create
	RepoName    string
	RepoDesc    string
	IsPublic    bool
	Source      string
	CreateClone bool
	// Checkouts and branches
	BranchName    string
	CreateBranch  bool
	OldBranchName string
	// PR
	Target      string
	Title       string
	BodyMessage string
	ID          string
	Comment     string
}

// Exported function to create a new model
func InitialModel() Model {
	return Model{
		Choices:     []string{"Add", "Commit", "Push", "Clone", "Show", "Branch", "Pull request", "Init", "Create repo"},
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
