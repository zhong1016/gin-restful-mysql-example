package api

import (
	v1 "todolist/api/v1"
	"todolist/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	api := r.Group("api")
	v1.V1Route(api)

	return r
}
