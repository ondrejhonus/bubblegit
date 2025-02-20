# Bubbletea Git TUI App

This is a simple TUI (Text User Interface) application written in Go using the [Bubbletea](https://github.com/charmbracelet/bubbletea) framework. The app provides some useful Git actions to help you manage your repositories from the terminal.

## Features

- **Add Files**: Add files to the staging area.
- **Commit Changes**: Commit staged changes with a message.
- **Push Changes**: Push committed changes to the remote repository.
- **Initialize Repository**: Initialize a new Git repository.

## Installation

To install the app, you need to have Go installed on your machine. Then, you can clone the repository and build the app:

```sh
git clone https://github.com/ondrejhonus/bubblegit.git
cd bubblegit
go build -o bubblegit
```

## Usage

Run the app from the terminal:

```sh
./bubblegit
```

Use the keyboard to navigate and perform Git actions.