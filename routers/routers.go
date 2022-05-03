package routers

import (
	"go_blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/assets", "./assets")

	engine.GET("/login", controller.GoLogin)
	engine.POST("/login", controller.Login)
	engine.GET("/", controller.Index)
	engine.Run()
}
