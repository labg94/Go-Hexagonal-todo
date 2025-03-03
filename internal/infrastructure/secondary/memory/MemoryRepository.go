package memory

import (
	"awesomeProject/internal/core/application"
	"awesomeProject/internal/core/domain"
	"strconv"
)

type memoryRepository struct {
	todos []domain.Todo
}

func (receiver *memoryRepository) GetAll() []domain.Todo {
	todosCopy := make([]domain.Todo, len(receiver.todos))
	copy(todosCopy, receiver.todos)
	return todosCopy

}

func (receiver *memoryRepository) FindById(id string) *domain.Todo {
	for i, oldTodo := range receiver.todos {
		if oldTodo.Id == id {
			return &receiver.todos[i]
		}
	}
	return nil
}

func (receiver *memoryRepository) Save(todo *domain.Todo) *domain.Todo {

	todo.Id = strconv.Itoa(len(receiver.todos) + 1)

	receiver.todos = append(receiver.todos, *todo)

	return todo
}

func (receiver *memoryRepository) Update(todo *domain.Todo) *domain.Todo {
	existingTodo := receiver.FindById(todo.Id)
	if existingTodo == nil {
		return nil
	}

	*existingTodo = *todo

	return existingTodo
}

func (receiver *memoryRepository) Delete(id string) {

	for i := range receiver.todos {
		if receiver.todos[i].Id == id {
			receiver.todos = append(receiver.todos[:i], receiver.todos[i+1:]...)
			break
		}
	}

}

func NewMemoryRepository() application.TodoRepository {
	return &memoryRepository{
		todos: []domain.Todo{},
	}
}
