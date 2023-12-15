package api

import (
	"bytes"
	"encoding/json"
	restful "go-exercise/internal"
	"go-exercise/service"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTaskRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup
	router := gin.New()
	newTaskRoute(router.Group(""))

	// Test
	var (
		routes    = router.Routes()
		hasGet    = false
		hasPost   = false
		hasPut    = false
		hasDelete = false
	)

	for _, route := range routes {
		switch route.Method {
		case "GET":
			if route.Path == "/tasks" {
				hasGet = true
			}
		case "POST":
			if route.Path == "/task" {
				hasPost = true
			}
		case "PUT":
			if route.Path == "/task/:id" {
				hasPut = true
			}
		case "DELETE":
			if route.Path == "/task/:id" {
				hasDelete = true
			}
		}
	}

	// Assertions
	assert.True(t, hasGet, "Expected GET /tasks route")
	assert.True(t, hasPost, "Expected POST /task route")
	assert.True(t, hasPut, "Expected PUT /task/:id route")
	assert.True(t, hasDelete, "Expected DELETE /task/:id route")
}

func TestTaskHandler_CreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup
	router := gin.New()
	taskHandler := NewtaskHandler()
	taskHandler.service = &mockTaskService{}

	router.POST("/task", taskHandler.CreateTask)

	// Test
	// Create HTTP POST request with JSON body
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/task", io.NopCloser(toJsonBody(t, CreateTaskRequest{Name: "NewTask"})))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response restful.ResponseOK
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "NewTask", response.Result["name"])
	assert.Equal(t, float64(0), response.Result["status"])
	assert.Greater(t, response.Result["id"], float64(0))
}

func TestTaskHandler_GetTaskList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup
	router := gin.New()
	taskHandler := NewtaskHandler()
	taskHandler.service = &mockTaskService{}

	router.GET("/tasks", taskHandler.GetTaskList)

	// Test
	// Create HTTP GET request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response restful.ResponseOKWithList
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Len(t, response.Result, 2)
	assert.Equal(t, "Task1", response.Result[0]["name"])
	assert.Equal(t, float64(1), response.Result[0]["status"])
}

func TestTaskHandler_UpdateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup
	router := gin.New()
	taskHandler := NewtaskHandler()
	taskHandler.service = &mockTaskService{}

	router.PUT("/task/:id", taskHandler.UpdateTask)

	// Test
	// Create HTTP PUT request with JSON body
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/task/1", io.NopCloser(toJsonBody(t, UpdateTaskRequest{ID: 1, Name: "UpdatedTask", Status: 1})))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response restful.ResponseOK
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "UpdatedTask", response.Result["name"])
	assert.Equal(t, float64(1), response.Result["status"])
	assert.Equal(t, float64(1), response.Result["id"])
}

func TestTaskHandler_DeleteTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup
	router := gin.New()
	taskHandler := NewtaskHandler()
	taskHandler.service = &mockTaskService{}

	router.DELETE("/task/:id", taskHandler.DeleteTask)

	// Test
	// Create HTTP DELETE request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/task/1", nil)

	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBoolToInt(t *testing.T) {
	// Test when input is true
	result := boolToInt(true)
	assert.Equal(t, 1, result, "Expected result to be 1 for true input")

	// Test when input is false
	result = boolToInt(false)
	assert.Equal(t, 0, result, "Expected result to be 0 for false input")
}

func TestIntToBool(t *testing.T) {
	// Test when input is 1
	result := intToBool(1)
	assert.True(t, result, "Expected result to be true for input 1")

	// Test when input is 0
	result = intToBool(0)
	assert.False(t, result, "Expected result to be false for input 0")

	// Test when input is any other value
	result = intToBool(123)
	assert.False(t, result, "Expected result to be false for any other input value")
}

// toJsonBody convert struct to JSON
func toJsonBody(t *testing.T, v interface{}) *bytes.Buffer {
	body, err := json.Marshal(v)
	require.NoError(t, err)
	return bytes.NewBuffer(body)
}

// Mock implementation of TaskService for testing
type mockTaskService struct {
	tasks map[uint]*service.Task
}

func (m *mockTaskService) GetTasks() (map[uint]*service.Task, error) {
	// Mock implementation
	// Return a map of tasks for testing
	return map[uint]*service.Task{
		1: {ID: 1, Name: "Task1", Status: true},
		2: {ID: 2, Name: "Task2", Status: false},
	}, nil
}

func (m *mockTaskService) CreateTask(name string) (*service.Task, error) {
	// Mock implementation
	// Create a new task and return it for testing
	return &service.Task{ID: 3, Name: name, Status: false}, nil
}

func (m *mockTaskService) UpdateTask(task service.Task) (*service.Task, error) {
	// Mock implementation
	// Update the task and return it for testing
	return &service.Task{ID: task.ID, Name: task.Name, Status: task.Status}, nil
}

func (m *mockTaskService) DeleteTask(id uint) error {
	// Mock implementation
	// Delete the task for testing
	delete(m.tasks, id)
	return nil
}
