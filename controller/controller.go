package controller

import (
	"go_blog/dao"
	"go_blog/plugin"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := dao.Mgr.CheckUser(username)
	passwordHash := plugin.CheckHashCode(password, user.Salt)
	if user.Username == "" {
		c.HTML(200, "login.html", "账号或者密码错误")
	} else {
		if user.Password != passwordHash {
			c.HTML(200, "login.html", "账号或密码错误")
		} else {
			c.Redirect(301, "/")
		}
	}
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
