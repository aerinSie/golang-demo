package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"golang-demo/src/config"
	"golang-demo/src/usermgmt" //golang-demo是 專案名稱
)

var (
	myenv string
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email seekastyle@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
func main() {
	// 初始化引擎
	config.Config(myenv)
	engine := gin.Default()
	// http://localhost:9205/swagger/index.html
	engine.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
	// 注册一个路由和处理函数
	engine.Any("/", WebRoot)
	engine.POST("/login", usermgmt.PostLogin)
	user := engine.Group("user")
	user.GET("/info", usermgmt.GetUserInfoByAccount)
	// user.GET("/:userID", userhandler.GetUserInfo)
	user.POST("/register", usermgmt.PostRegister)
	user.PUT("/change-password", usermgmt.PutUserChangePassword)

	// 绑定端口，然后启动应用
	// 本地 $go run main.go local
	// 遠端 $go run main.go prd
	engine.Run(":9205")
}

// WebRoot is ...
func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
