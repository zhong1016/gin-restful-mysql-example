package user

import (
	service "todolist/internal/service/user"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	user := r.Group("/users")
	{
		// Get all or ?id=1
		user.GET("", func(c *gin.Context) {

			id := c.Request.URL.Query().Get("id")
			if id == "" {
				c.JSON(200, service.GetAll())
				return
			}
			c.JSON(200, service.GetById(id))
		})

		// Get /1
		user.GET("/:id", func(c *gin.Context) {
			c.JSON(200, service.GetById(c.Param("id")))
		})

		// Post
		user.POST("", func(c *gin.Context) {

			// Json to map
			// var m map[string]interface{}
			// err := c.Bind(&m)

			// Json to struct
			var u service.User

			err := c.Bind(&u)
			if err != nil {
				return
			}

			c.JSON(200, service.Insert(u))
		})

		// Delete /1
		user.DELETE("/:id", func(c *gin.Context) {
			c.JSON(200, service.Delete(c.Param("id")))
		})

		// Update
		user.PUT("", func(c *gin.Context) {
			var u service.User

			err := c.Bind(&u)
			if err != nil {
				return
			}

			c.JSON(200, service.Update(u))
		})
	}
}
