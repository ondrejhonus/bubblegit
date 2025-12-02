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
		Choices:     []string{"1 | Add", "2 | Commit", "3 | Push", "4 | Clone", "5 | Show", "6 | Branch", "7 | Pull request", "8 | Init", "9 | Create repo"},
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
