package restful

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponseSuccess_WithResult(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		ResponseSuccess(c, gin.H{"key": "value"}, nil)
	})

	// Test
	// Create HTTP GET request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response ResponseOK
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	assert.Equal(t, gin.H{"key": "value"}, response.Result)
}

func TestResponseSuccess_WithResultList(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		ResponseSuccess(c, nil, []gin.H{{"item1": "value1"}, {"item2": "value2"}})
	})

	// Test
	// Create HTTP GET request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response ResponseOKWithList
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	assert.Len(t, response.Result, 2)
	assert.Equal(t, "value1", response.Result[0]["item1"])
	assert.Equal(t, "value2", response.Result[1]["item2"])
}

func TestResponseSuccess_WithoutResult(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		ResponseSuccess(c, nil, nil)
	})

	// Test
	// Create HTTP GET request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestResponseFail(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		ResponseFail(c, http.StatusBadRequest, "Bad Request")
	})

	// Test
	// Create HTTP GET request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ResponseError
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, "Bad Request", response.Message)
}

func TestNewErrorMessage(t *testing.T) {
	err := errors.New("test error")
	message := NewErrorMessage("An error occurred", err)
	assert.Equal(t, "An error occurred: test error", message)
}
