package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskMapService_CreateTask(t *testing.T) {
	taskService := NewTaskService()
	task, err := taskService.CreateTask("Test Task")

	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, "Test Task", task.Name)
	assert.Equal(t, uint(1), task.ID)
}

func TestTaskMapService_GetTasks(t *testing.T) {
	taskService := NewTaskService()
	task, _ := taskService.CreateTask("Test Task")
	tasks, err := taskService.GetTasks()

	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, task, tasks[task.ID])
}

func TestTaskMapService_UpdateTask(t *testing.T) {
	taskService := NewTaskService()
	task, _ := taskService.CreateTask("Test Task")
	task.Name = "Updated Task"
	updatedTask, err := taskService.UpdateTask(*task)

	assert.NoError(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, "Updated Task", updatedTask.Name)
}

func TestTaskMapService_DeleteTask(t *testing.T) {
	taskService := NewTaskService()
	task, _ := taskService.CreateTask("Test Task")
	err := taskService.DeleteTask(task.ID)

	assert.NoError(t, err)
	assert.Len(t, taskService.(*TaskMapService).tasks, 0)
}

func TestTaskMapService_DeleteTask_NotFound(t *testing.T) {
	taskService := NewTaskService()
	err := taskService.DeleteTask(1)

	assert.EqualError(t, err, ErrTaskNotFound.Error())
}
