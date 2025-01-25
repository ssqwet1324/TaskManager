package models

// Структура задачи и её методы

type Task struct {
	ID          string `json:"id"`          // id задачи
	Title       string `json:"name"`        // заголовок
	Description string `json:"description"` // краткое описание
	Deadline    string `json:"deadline"`    // дедлайн
	Completed   bool   `json:"completed"`   // статус задачи
}
