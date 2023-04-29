package main

import (
	"go.uber.org/zap"
)

func init() {
	Viper.InitViper("dev")
	objs := []*Object{
		{
			Name:         "mysql",
			New:          Mysql.InitMySQL,
			NewEveryTime: false,
		},
		{
			Name:         "zap",
			New:          Zap.InitZap,
			NewEveryTime: false,
		},
	}
	if err := Provide(objs...); err != nil {
		panic(err)
	}
}

func main() {
	var z *zap.Logger
	if err := Populate("zap", &z); err != nil {
		panic(err)
	}

	z.Info("zj is yyds")
}
