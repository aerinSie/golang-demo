package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"golang-demo/src/docs"
)

// $ go run main.go -env="prd" arg1
func Config(myenv string) {

	if os.Args[1] != "" {
		myenv = os.Args[1]
	} else {
		myenv = "prd"
	}

	// &env是用戶命令行的-env後的參數值
	// flag.StringVar(&myenv, "myenv", "local", "環境名，默認為local")
	// flag.Parse()               //解析所有註冊的flag
	viper.SetConfigName(myenv)           // 设置配置文件名 (不带后缀)
	viper.AddConfigPath("../src/config") // 第一个搜索路径
	err := viper.ReadInConfig()          // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

	host := viper.GetString("swagger.baseURL")
	docs.SwaggerInfo.Title = "Golang-Demo Swagger API"
	docs.SwaggerInfo.Description = "This is a sample server Golang-Demo server."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
