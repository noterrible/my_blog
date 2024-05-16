package main

import (
	"my_blog/global"
	"my_blog/initialization"
	"my_blog/routers"
)

func main() {
	initialization.InitConfig("conf.yaml")
	global.Log = initialization.InitLogger()
	initialization.InitMysql()
	initialization.InitBucket()
	r := routers.InitRouter()
	global.Log.Infof("run: %s", global.Config.System.Addr())
	err := r.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Error(err)
	}
}
