# Bubbletea Git TUI App

This is a simple TUI (Text User Interface) application written in Go using the [Bubbletea](https://github.com/charmbracelet/bubbletea) framework. The app provides some useful Git actions to help you manage your repositories from the terminal.

## Features

- **Add Files**
- **Commit Changes**
- **Push Changes**
- **Initialize Repository**
- **Create a repository**
- **Checkout Branches**
- **Merge Branches**
- **Rebase Branches**
- **Stash Changes**
- **Pull Requests**
- **Pull Requests**
- **View Logs**: View commit logs and history.

## TO-DO
- **Add list branches to menu > Branch**
- **Apply Stash**
- **Resolve Merge Conflicts**
- **Edit Diffs**

## Installation
#### Dependency list:
- golang
- git
- [gh-cli](https://cli.github.com/)

### Install dependencies

- ##### OpenSUSE
  ```bash
  sudo zypper addrepo https://cli.github.com/packages/rpm/gh-cli.repo
  sudo zypper ref
  sudo zypper install gh
  ```
- ##### Arch
  ```bash
  sudo pacman -Syu git go github-cli
  ```
- ##### Fedora (New DNF5)
  ```bash
  sudo dnf install dnf5-plugins
  sudo dnf config-manager addrepo --from-repofile=https://cli.github.com/packages/rpm/gh-cli.repo
  sudo dnf install gh --repo gh-cli
  ```
- ##### Fedora (Old DNF4)
  ```bash
  sudo dnf install 'dnf-command(config-manager)'
  sudo dnf config-manager --add-repo https://cli.github.com/packages/rpm/gh-cli.repo
  sudo dnf install gh --repo gh-cli
  ```
- ##### Debian
  ```bash
  (type -p wget >/dev/null || (sudo apt update && sudo apt install wget -y)) \
  	&& sudo mkdir -p -m 755 /etc/apt/keyrings \
  	&& out=$(mktemp) && wget -nv -O$out https://cli.github.com/packages/githubcli-archive-keyring.gpg \
  	&& cat $out | sudo tee /etc/apt/keyrings/githubcli-archive-keyring.gpg > /dev/null \
  	&& sudo chmod go+r /etc/apt/keyrings/githubcli-archive-keyring.gpg \
  	&& sudo mkdir -p -m 755 /etc/apt/sources.list.d \
  	&& echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null \
  	&& sudo apt update \
  	&& sudo apt install gh -y
  ```

### Install the tool globally
To install the app, you need to have Go installed on your machine. Then, you can clone the repository and build the app:

```sh
git clone https://github.com/ondrejhonus/bubblegit.git
cd bubblegit
./build.sh
```

## Usage

Then you can use the ```bubblegit``` command run the app from anywhere in the terminal:

Use the keyboard arrows or 'hjkl' to navigate and perform Git actions.
