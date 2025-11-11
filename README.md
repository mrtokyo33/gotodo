# GoTodo

![GoTodo logo](images/GOTODO.png)

![Go Version](https://img.shields.io/badge/go-1.18%2B-blue.svg)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macos%20%7C%20windows-lightgrey.svg)

## Overview
GoTodo is a lightweight and efficient command-line application for task management. It allows users to quickly add, list, complete, and remove tasks directly from the terminal. Perfect for anyone looking for a straightforward way to organize their to-do items.

This project is built with Go and the [Cobra](https://cobra.dev/) library.

## Installation

You can install `gotodo` directly using `go install`.

```bash
go install github.com/mrtokyo33/gotodo@latest
```

Ensure that your Go bin directory is included in your system's PATH. You can add it to your `.bashrc` or `.zshrc` file:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Quick Usage

Here is a quick demonstration of the main workflow:

```bash
# Add new tasks
$ gotodo add "Buy groceries for the week"
Task successfully added: "Buy groceries for the week"

$ gotodo add "Finish the project documentation"
Task successfully added: "Finish the project documentation"

$ gotodo add "Call dentist to schedule appointment"
Task successfully added: "Call dentist to schedule appointment"

$ gotodo add "Review pull requests"
Task successfully added: "Review pull requests"

# List all tasks
$ gotodo list

Current To-Do List:
[ ] 1: Buy groceries for the week
[ ] 2: Finish the project documentation
[ ] 3: Call dentist to schedule appointment
[ ] 4: Review pull requests

# Mark tasks as done
$ gotodo do 1
Task 1 marked as done

$ gotodo do 2
Task 2 marked as done

# Remove a task
$ gotodo rm 3
Task 3 removed successfully

# List tasks again to see the changes
$ gotodo list

Current To-Do List:
[✓] 1: Buy groceries for the week
[✓] 2: Finish the project documentation
[ ] 4: Review pull requests
```

## Commands

Here is a complete list of all available commands and their descriptions.

| Command | Description |
|---------|-------------|
| `help` | Help about any command. |
| `add` | Quickly create a new task. |
| `do` | Mark a task as done. |
| `undo` | Mark a task as pending. |
| `list` | View all tasks. |
| `rm` | Delete a task permanently. |

## Global Flags

| Flag | Description |
|------|-------------|
| `-h, --help` | Help for gotodo. |
| `-t, --toggle` | Help message for toggle. |

## Examples

### Add a Task

You can add any task as a string argument.

```bash
$ gotodo add "Call the mechanic to schedule service"
```

### Complete a Task

Use the `do` command followed by the task's ID number.

```bash
# Assuming "Buy groceries" is task 5
$ gotodo do 5
```

### Undo a Task

If you complete a task by mistake, you can mark it as pending again.

```bash
$ gotodo undo 5
```

### Remove a Task

To permanently delete a task, use `rm`.

```bash
$ gotodo rm 2
```

### Get Help

You can get help for any specific command.

```bash
$ gotodo add --help
```