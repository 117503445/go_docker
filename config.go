package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	var1 string
	var2 string
}

func (config *config) Load() {
	pflag.String("var1", "code_var1", "var1")
	pflag.String("var2", "code_var2", "var2")
	pflag.Parse()

	// 绑定命令行
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Println(err)
	}
	err = viper.ReadInConfig()
	if err != nil {
		//panic(err)
	}
	config.var1 = viper.GetString("var1")
	config.var2 = viper.GetString("var2")
}

var AppConfig *config = &config{}
