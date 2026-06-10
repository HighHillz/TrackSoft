# TrackSoft — CLI Project Tracker

A high-performance, terminal-first project and deadline tracker built in **Go** and **SQLite**, designed to streamline developer workflows. Track deadlines, manage project health, and stay organized entirely from your shell.

![Latest Version](https://img.shields.io/badge/Latest%20Version-1.0.0-blue)
![Status](https://img.shields.io/badge/Status-Inactive-yellow)


## 🌟 Features

*   **Interactive TUI Dashboard**: A gorgeous Bubble Tea & Lip Gloss dashboard showing a live list of your projects, their category, deadlines, and a custom color-coded health bar.
*   **Natural Language Due Dates**: Powered by `go-naturaldate`, you can define deadlines using natural language like `"tomorrow"`, `"next friday"`, `"in 2 weeks"`, or `"monday afternoon"`.
*   **ASCII Tabular Summary**: A quick-scan table printed to the terminal showing all project statistics (ID, Name, Type, Due Date, Health, Status).
*   **Active Overdue Alerts**: Automatic desktop notifications via `beeep` if any projects are detected as overdue during listing.
*   **JSON Backups**: Quick export of all database projects to a local JSON file for backup and portability.
*   **Local SQLite Storage**: Fast, persistent storage utilizing a local SQLite database file in the user's config directory.

## 🚀 How to Run

### Installation

If you have Go installed on your system, you can build the project from source:

```bash
go build -o tracker main.go
```

Alternatively, you can use the precompiled `track` binary in the repository root:

```bash
# Set execution permissions
chmod +x track

# Run track directly
./track
```

## 📁 Data & Configuration Paths

*   **SQLite Database**: `~/.config/tracker/projects.db` (automatically created on initialization)
*   **Backup Path**: `~/.config/tracker/backup.json` (generated via the `backup` command)

## 🛠 Command Reference

### 1. `tracker dash`
Launches the interactive Terminal User Interface (TUI) dashboard.
*   **Keys**:
    *   `q` or `Ctrl+C`: Quit the dashboard.
    *   `d`: Interactively delete/remove the currently selected project.
    *   `↑` / `↓` / `k` / `j`: Navigate through the project list.

```bash
./track dash
```

### 2. `tracker add`
Adds a new project to the tracker database.
*   **Flags**:
    *   `-n, --name string`: Name of the project (can also be passed as an argument).
    *   `-t, --type string`: Project type (options: `UI/UX`, `Web`, `CLI`, `GUI`, `Script`, `Core`. Defaults to `CLI`).
    *   `-d, --due string`: Natural language due date (defaults to `"next friday"` if unspecified).

If no type is specified, the application will display an interactive selection prompt using `promptui`.

```bash
# Example with inline argument and default deadline
./track add "My Go Project"

# Example specifying type and natural language due date
./track add -n "Design Mockup" -t "UI/UX" -d "next Monday"
```

### 3. `tracker list`
Displays a clean, styled ASCII table of all projects, showing their current health percentages and status.
*   **Statuses**:
    *   `NEW`: Project created within the last 24 hours.
    *   `Pending`: Active project with time remaining.
    *   `OVERDUE`: Project deadline has passed.
    *   `Done`: Completed project.
*   **Note**: If any projects are overdue, a system notification will alert you.

```bash
./track list
```

### 4. `tracker edit`
Modifies an existing project's deadline.
*   **Flags**:
    *   `-d, --due string` (Required): New natural language due date.

```bash
# Update project with ID 1 to be due next Friday
./track edit 1 --due "next Friday"
```

### 5. `tracker rm`
Removes a project from the tracker database by ID.

```bash
./track rm 1
```

### 6. `tracker backup`
Exports all database records to a JSON file format.

```bash
./track backup
```

---

## ⚡ Tech Stack

*   **CLI Framework**: [Cobra](https://github.com/spf13/cobra)
*   **TUI Components**: [Bubble Tea](https://github.com/charmbracelet/bubbletea) & [Lip Gloss](https://github.com/charmbracelet/lipgloss)
*   **Natural Language Parser**: [go-naturaldate](https://github.com/tj/go-naturaldate)
*   **Interactive Prompts**: [Promptui](https://github.com/manifoldco/promptui)
*   **Notification Engine**: [Beeep](https://github.com/gen2brain/beeep)
*   **SQLite Driver**: [modernc.org/sqlite](https://modernc.org/sqlite) (pure Go SQLite driver)