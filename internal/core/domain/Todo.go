package domain

import "time"

type Status int8

const (
	Created Status = iota
	InProgress
	Done
)

type Todo struct {
	Id          string
	Title       string
	Description string
	Status      Status
	LastUpdate  time.Time
}

func TodoFrom(title string, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Status:      Created,
		LastUpdate:  time.Now(),
	}
}

func (todo *Todo) UpdateStatus() {

	if todo.Status == Created {
		todo.Status = InProgress
	} else {
		todo.Status = Done
	}
	todo.LastUpdate = time.Now()
}
