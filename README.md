# Bubbletea Git TUI App

This is a simple TUI (Text User Interface) application written in Go using the [Bubbletea](https://github.com/charmbracelet/bubbletea) framework. The app provides some useful Git actions to help you manage your repositories from the terminal.

## Features

- **Add Files**: Add files to the staging area.
- **Commit Changes**: Commit staged changes with a message.
- **Push Changes**: Push committed changes to the remote repository.
- **Initialize Repository**: Initialize a new Git repository.
- **Create a repository**: Create a repo from local dir or empty remote
- **Checkout Branches**: Switch between different branches in your repository.
- **Merge Branches**: Merge changes from one branch to another.
- **Rebase Branches**: Rebase your current branch onto another branch.
- **Stash Changes**: Stash your uncommitted changes.
- **Pull Requests**: Create and manage pull requests.

## TO-DO
- **View Logs**: View commit logs and history.
- **Add list branches to menu > Branch**
- **Apply Stash**: Apply stashed changes back to your working directory.
- **Resolve Conflicts**: Resolve merge conflicts.

## Installation

To install the app, you need to have Go installed on your machine. Then, you can clone the repository and build the app:

```sh
git clone https://github.com/ondrejhonus/bubblegit.git
cd bubblegit
./build.sh
```

## Usage

Then you can use the ```bubblegit``` command run the app from anywhere in the terminal:

Use the keyboard arrows or 'hjkl' to navigate and perform Git actions.
