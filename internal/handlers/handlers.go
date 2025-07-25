package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"taskmanager/internal/entity"
	"taskmanager/internal/repository"
)

// TaskHandler — структура для обработки запросов, связанных с задачами.
type TaskHandler struct {
	storage *repository.Repository
}

// New — конструктор, создающий новый обработчик задач.
func New(storage *repository.Repository) *TaskHandler {
	return &TaskHandler{storage: storage}
}

// GetTasks — обработчик для получения всех задач.
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// Загружаем задачи из хранилища
	tasks, err := h.storage.GetTasks()
	if err != nil {
		http.Error(w, "Не удалось загрузить задачи", http.StatusInternalServerError)

		return
	}

	// Отправляем задачи как JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		return
	}

}

// AddTask — обработчик для добавления новой задачи.
func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)

	// Декодируем JSON из тела запроса в структуру Task.
	var req entity.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный формат JSON", http.StatusBadRequest)
		return
	}

	if err := h.storage.AddTask(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(req)
	if err != nil {
		return
	}
}

// EditTask обработчик для редактирования задачи
func (h *TaskHandler) EditTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)

		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)

	var updateTask entity.Task
	if err := json.NewDecoder(r.Body).Decode(&updateTask); err != nil {
		http.Error(w, "Некорректный формат JSON", http.StatusBadRequest)

		return
	}

	if err := h.storage.EditTask(updateTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(updateTask)
	if err != nil {
		return
	}
}

// DeleteTask обработчик для удаления задачи
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)

		return
	}

	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		http.Error(w, "ID не найден", http.StatusBadRequest)

		return
	}

	if err := h.storage.DeleteTask(taskID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Задача успешно удалена"))
	if err != nil {
		return
	}
}
