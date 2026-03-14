# Project Tracker CLI

A personal, terminal-based project tracker designed for the modern developer workflow. Built with Go, Cobra, SQLite, Lipgloss, and Bubble Tea.

![Version](https://img.shields.io/badge/Latest%20Version-1.0.0-blue)
![License](https://img.shields.io/badge/license-MIT-black)

## Features

- **Local & Fast**: Pure Go SQLite database stored automatically in your OS user config directory.
- **Smart Inputs**: Interactive project type selection and natural language deadline parsing (e.g., "next friday", "tomorrow").
- **Sensory Dashboard**: Color-coded project health and status tracking. 
- **TUI Mode**: A full-screen interactive dashboard (Bubble Tea) to browse and manage projects.
- **Desktop Alerts**: Automatic OS-level popup notifications for overdue projects.
- **Backup Support**: Export your database to a JSON file at any time.

## Prerequisites

- Go 1.24+ (if compiling from source).

## Installation

1. Clone or download the repository.
2. Build the binary:
   ```bash
   go build -o track main.go
   ```
3. Move the binary to a location in your system's PATH:
   ```bash
   sudo mv track /usr/local/bin/
   ```
   *(Optional)* Add an alias in your `~/.bashrc` or `~/.zshrc`:
   ```bash
   alias track='/usr/local/bin/track'
   ```

## Usage

### 1. Add a Project
Add a project interactively (you will be prompted to select a type):
```bash
./track add --name "My New App"
```

Add a project inline with natural language dates:
```bash
./track add --name "Weekend MVP" --due "this sunday" --type "Web"
./track add --name "Fix tests" --due "tomorrow" --type "CLI"
./track add --name "Client Proposal" --due "2026-04-15" --type "UI/UX"
```

### 2. View Projects
Show the standard table view. Health bars show remaining time until the deadline.
```bash
./track list
```

### 3. Interactive Dashboard (TUI)
Launch the full-screen terminal UI to browse your projects interactively.
```bash
./track dash
```
- Navigate with **Arrow Keys** or **j/k**.
- Press **d** to delete the currently selected project.
- Press **q** or **Ctrl+C** to quit.

### 4. Edit a Project's Due Date
You can modify the deadline of an existing project by using its ID.
```bash
./track edit <ID> --due "new date"
# Example: ./track edit 3 --due "next month"
```

### 5. Remove a Project
Remove a project by its ID (as seen in the `list` or `dash` view).
```bash
./track rm <ID>
# Example: ./track rm 3
```

### 6. Backup Database
Export your current SQLite database to a formatted JSON file located in your user config directory.
```bash
./track backup
```

## Database Location
The SQLite database is automatically created and stored in your OS config directory:
- Linux: `~/.config/tracker/projects.db`
- macOS: `~/Library/Application Support/tracker/projects.db`
- Windows: `%AppData%\tracker\projects.db`

## License
MIT License. See `LICENSE` for more information.
