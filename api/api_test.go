package api

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup
	router := gin.New()
	RegisterRoutes(router)

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
