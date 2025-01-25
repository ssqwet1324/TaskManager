package repository

import (
	"fmt"
	"taskmanager/models"
)

type Repository struct {
	taskRepo map[string]models.Task
}

// хранилище
func NewRepository() *Repository {
	return &Repository{
		taskRepo: make(map[string]models.Task),
	}
}

// создаем загрузчик задачек
func (r *Repository) GetTasks() ([]models.Task, error) {
	var resp []models.Task
	fmt.Println("GetTask до", resp)
	for _, task := range r.taskRepo {
		resp = append(resp, task)
	}
	fmt.Println("GetTask после", resp)
	return resp, nil
}

func (r *Repository) AddTasks(task models.Task) error {
	r.taskRepo[task.ID] = task
	fmt.Println("AddTask", r.taskRepo)
	return nil
}

func (r *Repository) EditTask(task models.Task) error {
	fmt.Println("Edit до", r.taskRepo)
	r.taskRepo[task.ID] = task
	fmt.Println("EditTask", r.taskRepo)
	return nil
}

func (r *Repository) DeleteTask(id string) error {
	fmt.Println("Delete до", r.taskRepo)
	delete(r.taskRepo, id)
	fmt.Println("Delete после", r.taskRepo)
	return nil
}
