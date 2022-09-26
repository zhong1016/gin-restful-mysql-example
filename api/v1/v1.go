package v1

import (
	"todolist/api/v1/user"

	"github.com/gin-gonic/gin"
)

func V1Route(r *gin.RouterGroup) {

	v1 := r.Group("v1")

	user.UserRoute(v1)
}
