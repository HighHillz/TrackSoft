package db

import (
	"database/sql"
	"os"
	"path/filepath"
	"time"
	"tracker/models"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	appDir := filepath.Join(configDir, "tracker")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return err
	}

	dbPath := filepath.Join(appDir, "projects.db")
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		type TEXT,
		deadline DATETIME,
		created_at DATETIME,
		is_completed BOOLEAN DEFAULT 0
	);`

	_, err = DB.Exec(createTableSQL)
	return err
}

func AddProject(p *models.Project) error {
	_, err := DB.Exec("INSERT INTO projects (name, type, deadline, created_at, is_completed) VALUES (?, ?, ?, ?, ?)",
		p.Name, p.Type, p.Deadline, p.CreatedAt, p.IsCompleted)
	return err
}

func GetAllProjects() ([]models.Project, error) {
	rows, err := DB.Query("SELECT id, name, type, deadline, created_at, is_completed FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Deadline, &p.CreatedAt, &p.IsCompleted)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

func UpdateProjectDeadline(id int64, newDeadline time.Time) error {
	_, err := DB.Exec("UPDATE projects SET deadline = ? WHERE id = ?", newDeadline, id)
	return err
}

func RemoveProject(id int64) error {
	_, err := DB.Exec("DELETE FROM projects WHERE id = ?", id)
	return err
}
