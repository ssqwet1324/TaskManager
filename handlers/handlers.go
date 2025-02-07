package handlers

import (
	"encoding/json"
	"net/http"
	"taskmanager/models"
	"taskmanager/repository" // хранилище
)

// TaskHandler — структура для обработки запросов, связанных с задачами.
type TaskHandler struct {
	storage *repository.Repository
}

// NewHandler — конструктор, создающий новый обработчик задач.
func NewHandler(storage *repository.Repository) *TaskHandler {
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
	json.NewEncoder(w).Encode(tasks)

}

// AddTask — обработчик для добавления новой задачи.
func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Декодируем JSON из тела запроса в структуру Task.
	var req models.Task
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
	json.NewEncoder(w).Encode(req)
}

// обработчик для редактирования задачи
func (h *TaskHandler) EditTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var updateTask models.Task
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
	json.NewEncoder(w).Encode(updateTask)
}

// обработчикк для удаления задачи
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
	w.Write([]byte("Задача успешно удалена"))
}
