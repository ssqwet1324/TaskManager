package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"taskmanager/models"
	"time"
)

type Repository struct {
	db *sql.DB
}

// Создание нового репозитория (подключение к БД)
func NewRepository() *Repository {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Error connection to BD:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("БД недоступна:", err)
	}

	log.Println("✅ Connection to Postgres Comleted!")
	return &Repository{db}
}

// Получение всех задач из БД
func (r *Repository) GetTasks() ([]models.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, deadline FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		var deadlineRaw string

		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &deadlineRaw); err != nil {
			return nil, err
		}

		parsedTime, err := time.Parse(time.RFC3339, deadlineRaw)
		if err != nil {
			fmt.Println("Ошибка парсинга даты:", err)
			task.Deadline = ""
		} else {
			task.Deadline = parsedTime.Format("02-01-2006")
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Добавление задачи в БД
func (r *Repository) AddTask(task models.Task) error {
	if task.ID == uuid.Nil {
		task.ID = uuid.New()
	}
	query := "INSERT INTO tasks (id, title, description, deadline, completed) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, task.ID, task.Title, task.Description, task.Deadline, task.Completed)
	//fmt.Printf("addtask", task.Title, task.Description, task.Deadline, task.ID)
	return err
}

// Редактирование задачи в БД
func (r *Repository) EditTask(task models.Task) error {
	_, err := r.db.Exec("UPDATE tasks SET title=$1, description=$2, deadline=$3 WHERE id=$4", task.Title, task.Description, task.Deadline, task.ID)
	//fmt.Printf("edittask", task.Title, task.Description, task.Deadline, task.ID)
	return err
}

// Удаление задачи из БД
func (r *Repository) DeleteTask(id string) error {
	//log.Println("UUID:", id)
	// Преобразуем строку id в тип uuid.UUID
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("incorrect UUID: %v", err)
	}

	_, err = r.db.Exec("DELETE FROM tasks WHERE id=$1", parsedID)
	return err
}
