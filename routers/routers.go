package routers

import (
	"go_blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/assets", "./assets")

	engine.GET("/index", controller.ListUser)
	engine.POST("/index", controller.AddUser)
	engine.Run()
}
