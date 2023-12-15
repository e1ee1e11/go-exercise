package service

import (
	"errors"
	"math"
)

var ErrTaskNotFound = errors.New("task is not found")

// TaskService represents a service for managing tasks.
type TaskService interface {
	GetTasks() (map[uint]*Task, error)
	// GetTask() (Task, error)  // not implemented
	CreateTask(name string) (*Task, error)
	UpdateTask(task Task) (*Task, error)
	DeleteTask(id uint) error
}

// Task represents a task entity.
type Task struct {
	ID     uint
	Name   string
	Status bool
}

// TaskMapService implements TaskService using a map.
type TaskMapService struct {
	tasks     map[uint]*Task
	taskCount uint
}

// NewTaskService creates a new task service.
func NewTaskService() TaskService {
	return &TaskMapService{tasks: make(map[uint]*Task)}
}

// CreateTask creates a new task with the given name.
func (t *TaskMapService) CreateTask(name string) (*Task, error) {
	// Check if taskCount reaches uint max value
	if t.taskCount == math.MaxUint {
		return nil, errors.New("task list is full")
	}

	t.taskCount++

	task := &Task{ID: t.taskCount, Name: name, Status: false}
	t.tasks[task.ID] = task

	return task, nil
}

// GetTasks returns all tasks.
func (t *TaskMapService) GetTasks() (map[uint]*Task, error) {
	return t.tasks, nil
}

// UpdateTask updates the given task.
func (t *TaskMapService) UpdateTask(task Task) (*Task, error) {
	// Check if task exists
	if _, ok := t.tasks[task.ID]; !ok {
		return nil, ErrTaskNotFound
	}

	// Update task
	t.tasks[task.ID] = &task

	return &task, nil
}

// DeleteTask deletes the task with the given ID.
func (t *TaskMapService) DeleteTask(id uint) error {
	// Check if task exists
	if _, ok := t.tasks[id]; !ok {
		return ErrTaskNotFound
	}

	// Delete task
	delete(t.tasks, id)

	return nil
}

// TODO: implement TaskDBService to handle tasks stored in a database.
// TaskDBService implements TaskService using a database.
// type TaskDBService struct {}
