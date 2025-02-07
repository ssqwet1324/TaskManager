package models

import (
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`          // id задачи (UUID)
	Title       string    `json:"name"`        // заголовок
	Description string    `json:"description"` // краткое описание
	Deadline    string    `json:"deadline"`    // дедлайн (теперь string)
	Completed   bool      `json:"completed"`   // статус задачи
}
