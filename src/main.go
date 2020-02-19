package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang-demo/src/handler"//golang-demo是 go.mod檔的module 
)

func main() {
	// 初始化引擎

	engine := gin.Default()
	// 注册一个路由和处理函数
	engine.Any("/", WebRoot)

	engine.POST("/login", handler.PostLogin)

	user := engine.Group("user")
	user.GET("/info", handler.GetUserInfoByAccount)
	// user.GET("/:userID", userhandler.GetUserInfo)
	user.POST("/register", handler.PostRegister)
	user.PUT("/change-password", handler.PutUserChangePassword)

	// 绑定端口，然后启动应用
	engine.Run(":9205")
}

//WebRoot is ...
/*
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
 */
func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
