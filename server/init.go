package server

import (
    "runtime"
    "github.com/gin-gonic/gin"
    "github.com/tonycbcd/easycms/server/conf"
    "github.com/tonycbcd/easycms/server/components/egin"
)

var (
    app *egin.Engine
)

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    gin.SetMode(conf.Config.Global.Mode)
    app = egin.NewEngine(gin.Default())
}
