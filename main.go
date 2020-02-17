package main

import (
	"net/http"

	"./handler/userhandler"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	engine.Any("/", WebRoot)

	engine.POST("/login", userhandler.PostLogin)

	user := engine.Group("user")
	user.GET("/user-info", userhandler.GetUserInfoByAccount)
	// user.GET("/:userID", userhandler.GetUserInfo)
	user.POST("/register", userhandler.PostRegister)
	user.PUT("/change-password", userhandler.PutUserChangePassword)

	// 绑定端口，然后启动应用
	engine.Run(":9205")
}

/**
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
 */
func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
