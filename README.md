# Git Helper

Git Helper is a command-line interface tool designed to simplify common Git operations. It provides an intuitive set of commands for managing Git repositories efficiently.

## Features

- **Branch Management**: List, create, switch, and delete branches.
- **Commit Operations**: Commit changes with options for all or specific files.
- **Push and Pull**: Easily synchronize your local repository with remote.
- **Reset Options**: Reset your working directory or undo commits.
- **Logging**: Tracks all Git operations performd by the tool.
- **Custom Configuration**: Set default values for remote and branch names.

## Installation

To install Git Helper, clone the repository and build the project:

```bash
git clone https://github.com/yourusername/git-helper.git
cd git-helper
go build
```

## Usage

### Basic Commands

- **List Branches**:
  ```bash
  git-helper branch list
  ```

- **Create a New Branch**:
  ```bash
  git-helper branch create <branch_name>
  ```

- **Switch Branches**:
  ```bash
  git-helper branch switch <branch_name>
  ```

- **Delete a Branch**:
  ```bash
  git-helper branch delete <branch_name>
  ```

- **Commit Changes**:
  ```bash
  git-helper commit --all --mmessage "Commit message"
  git-helper commit -a -m "Commit message"

  git-helper commit --file <file_name> --mmessage "Commit message"
  git-helper commit -f <file_name> -m "Commit message"

  git-helper commit --file <file_name> --file <file_name> --message "Commit message"
  git-helper commit -f <file_name> -f <file_name> -m "Commit message"
  ```

- **Logging**:
  ```bash
  git-helper log

  git-helper log --oneline
  git-helper log -o
  ```

- **Reset**:
  ```bash
  git-helper reset --discard
  git-helper reset -d

  git-helper reset --undo
  git-helper reset -u
  ```

- **Status**:
  ```bash
  git-helper status

  git-helper status --tracked
  git-helper status -t

  git-helper status --modified
  git-helper status -m

  git-helper status -untracked
  git-helper status -u
  ```

- **Automated Pull and Push**:
  ```bash
  git-helper sync
  ```

- **Set default remote**:
  ```bash
  git-helper config remote origin
  git-helper config default_branch main
  ```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any features or improvements.

