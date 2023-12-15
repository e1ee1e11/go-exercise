package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	g := r.Group("")

	newTaskRoute(g)

	// TODO: add more routes here
}
