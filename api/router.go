package api

import (
	"graduation_project/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	group := router.Group("")
	user_group := group.Group("/user")
	{
		user_group.GET("/host", handler.UserHostShow)
	}
	admin_group := group.Group("/admin")
	{
		admin_group.GET("/host", handler.AdminHostShow)
	}
	return router
}
