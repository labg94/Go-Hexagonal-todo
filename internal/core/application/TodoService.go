package application

import (
	"awesomeProject/internal/core/domain"
	"errors"
)

type TodoRepository interface {
	GetAll() []domain.Todo
	FindById(id string) *domain.Todo
	Save(todo *domain.Todo) *domain.Todo
	Update(todo *domain.Todo) *domain.Todo
	Delete(id string)
}

type TodoService struct {
	repository TodoRepository
}

func (service *TodoService) GetAll() []domain.Todo {
	return service.repository.GetAll()
}

func (service *TodoService) FindById(id string) (*domain.Todo, error) {
	byId := service.repository.FindById(id)
	if byId == nil {
		return nil, errors.New("there is not a todo with id " + id)
	}
	return byId, nil
}

func (service *TodoService) NewFrom(title string, description string) *domain.Todo {
	todo := domain.TodoFrom(title, description)
	return service.repository.Save(todo)
}

func (service *TodoService) UpdateStatus(id string) (*domain.Todo, error) {
	findByID, err := service.FindById(id)
	if err != nil {
		return nil, err
	}
	findByID.UpdateStatus()
	return service.repository.Update(findByID), nil
}

func (service *TodoService) DeleteById(id string) {
	service.repository.Delete(id)
}

func NewTodoService(todoRepository TodoRepository) *TodoService {
	return &TodoService{repository: todoRepository}
}
