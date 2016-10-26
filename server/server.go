// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// To initialize the config data.

package server

import (
    "github.com/tonycbcd/easycms/server/conf"
    "github.com/tonycbcd/easycms/server/models"
    "github.com/tonycbcd/easycms/server/routers"
    mid "github.com/tonycbcd/easycms/server/middlewares"
)

func Run() {
    models.Init()
    app.Use(mid.SetEnv(app))

    routers.SetRouters(app)

    app.Run(":" + conf.Config.Global.Port)
}
