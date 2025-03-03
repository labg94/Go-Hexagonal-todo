package rest

import (
	"awesomeProject/internal/core/application"
	"awesomeProject/internal/core/domain"
	"encoding/json"
	"net/http"
	"strings"
)

type TodoResponseDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var statusMapper = map[domain.Status]string{
	domain.Created:    "Created",
	domain.Done:       "Done",
	domain.InProgress: "In Progress",
}

func mapTodoToResponseDTO(todo *domain.Todo) TodoResponseDTO {
	return TodoResponseDTO{
		ID:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      statusMapper[todo.Status],
	}
}

type CreateTodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ResponseError struct {
	Message string `json:"message"`
}

// AppHandler with a TodoService instance
type AppHandler struct {
	service *application.TodoService
}

func NewAppHandler(service *application.TodoService) *AppHandler {
	return &AppHandler{service: service}
}

// Helper method to write JSON response
func writeJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
		return
	}
}

func (h *AppHandler) GetAllTodos(w http.ResponseWriter) {
	todos := h.service.GetAll()

	todosResponse := make([]TodoResponseDTO, len(todos))
	for i, todo := range todos {
		todosResponse[i] = mapTodoToResponseDTO(&todo)
	}

	writeJSON(w, http.StatusOK, todosResponse)
}

func (h *AppHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	todo, err := h.service.FindById(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, ResponseError{Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, mapTodoToResponseDTO(todo))
}

func (h *AppHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var dto CreateTodoDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		writeJSON(w, http.StatusBadRequest, ResponseError{Message: "Invalid JSON"})
		return
	}
	todo := h.service.NewFrom(dto.Title, dto.Description)
	writeJSON(w, http.StatusCreated, mapTodoToResponseDTO(todo))
}

func (h *AppHandler) UpdateTodoStatus(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	todo, err := h.service.UpdateStatus(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, ResponseError{Message: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, mapTodoToResponseDTO(todo))
}

func (h *AppHandler) DeleteTodoByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	h.service.DeleteById(id)
	writeJSON(w, http.StatusNoContent, nil)
}
