package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	AppConfig.Load()
	fmt.Println(AppConfig.var1)
	fmt.Println(AppConfig.var2)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		_, _ = c.Writer.WriteString("hello, world" + AppConfig.var1 + AppConfig.var2)
	})
	_ = r.Run("0.0.0.0:80")
}
