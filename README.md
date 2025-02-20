# Bubbletea Git TUI App

This is a simple TUI (Text User Interface) application written in Go using the [Bubbletea](https://github.com/charmbracelet/bubbletea) framework. The app provides some useful Git actions to help you manage your repositories from the terminal.

## Features

- **View Status**: Check the current status of your Git repository.
- **Stage Changes**: Stage files for commit.
- **Commit Changes**: Commit staged changes with a message.
- **View Log**: View the commit history of the repository.
- **Switch Branches**: List and switch between branches.

## Installation

To install the app, you need to have Go installed on your machine. Then, you can clone the repository and build the app:

```sh
git clone https://github.com/yourusername/bubbletea-git-tui.git
cd bubbletea-git-tui
go build -o git-tui
```

## Usage

Run the app from the terminal:

```sh
./git-tui
```

Use the keyboard to navigate and perform Git actions.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with any improvements or bug fixes.

## License

This project is licensed under the MIT License.