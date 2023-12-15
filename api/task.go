package api

import (
	"fmt"
	"net/http"
	"strconv"

	restful "go-exercise/internal"
	"go-exercise/service"

	"github.com/gin-gonic/gin"
)

// TaskHandler handles requests related tasks
type TaskHandler struct {
	service service.TaskService
}

// CreateTaskRequest is the request body for creating a task
type CreateTaskRequest struct {
	Name string `json:"name"`
}

// UpdateTaskRequest is the request body for updating a task
type UpdateTaskRequest struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func newTaskRoute(g *gin.RouterGroup) {
	taskHandler := NewtaskHandler()
	g.GET("/tasks", taskHandler.GetTaskList)
	g.POST("/task", taskHandler.CreateTask)
	g.PUT("/task/:id", taskHandler.UpdateTask)
	g.DELETE("/task/:id", taskHandler.DeleteTask)
}

// NewtaskHandler creates a new task handler
func NewtaskHandler() *TaskHandler {
	return &TaskHandler{
		service: service.NewTaskService(),
	}
}

// GetTaskList returns a list of tasks
func (t *TaskHandler) GetTaskList(c *gin.Context) {
	tasks, err := t.service.GetTasks()
	if err != nil {
		restful.ResponseFail(c, http.StatusInternalServerError, restful.NewErrorMessage("Failed to get task list", err))
		return
	}

	resultList := []gin.H{}
	for _, task := range tasks {
		resultList = append(resultList, gin.H{
			"name":   task.Name,
			"status": boolToInt(task.Status),
			"id":     task.ID,
		})
	}
	restful.ResponseSuccess(c, nil, resultList)
}

// CreateTask creates a new task
func (t *TaskHandler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.Bind(&req); err != nil {
		restful.ResponseFail(c, http.StatusBadRequest, restful.NewErrorMessage("Failed to parse json body", err))
		return
	}
	fmt.Printf("CreateTaskRequest = %+v\n", req)

	task, err := t.service.CreateTask(req.Name)
	if err != nil {
		restful.ResponseFail(c, http.StatusInternalServerError, restful.NewErrorMessage("Failed to create a task", err))
		return
	}

	restful.ResponseSuccess(c, gin.H{
		"name":   task.Name,
		"status": boolToInt(task.Status),
		"id":     task.ID,
	}, nil)
}

// UpdateTask updates a task
func (t *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		restful.ResponseFail(c, http.StatusBadRequest, restful.NewErrorMessage("Failed to parse id", err))
		return
	}

	var req UpdateTaskRequest
	if err := c.Bind(&req); err != nil {
		restful.ResponseFail(c, http.StatusBadRequest, restful.NewErrorMessage("Failed to parse json body", err))
		return
	}
	fmt.Printf("UpdateTaskRequest = %+v\n", req)

	task, err := t.service.UpdateTask(
		service.Task{
			ID:     uint(id), // Use the id from param because it's unchangeable
			Name:   req.Name,
			Status: intToBool(req.Status),
		},
	)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == service.ErrTaskNotFound {
			statusCode = http.StatusNotFound
		}
		restful.ResponseFail(c, statusCode, restful.NewErrorMessage("Failed to update the task", err))
		return
	}

	restful.ResponseSuccess(c, gin.H{
		"name":   task.Name,
		"status": boolToInt(task.Status),
		"id":     task.ID,
	}, nil)
}

// DeleteTask deletes a task
func (t *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		restful.ResponseFail(c, http.StatusBadRequest, restful.NewErrorMessage("Failed to parse id", err))
		return
	}
	fmt.Printf("id = %v\n", id)

	err = t.service.DeleteTask(uint(id))
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == service.ErrTaskNotFound {
			statusCode = http.StatusNotFound
		}
		restful.ResponseFail(c, statusCode, restful.NewErrorMessage("Failed to delete the task", err))
		return
	}

	restful.ResponseSuccess(c, nil, nil)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func intToBool(i int) bool {
	return i == 1
}
