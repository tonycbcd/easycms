// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// The Api Controllers.

package controllers

import (
    "github.com/tonycbcd/easycms/server/components/egin"

    // mid "github.com/tonycbcd/easycms/server/middlewares"
)

func OmLogin(c *egin.Context) {
    isWrong := false
    email   := ""

    params  := egin.H{
        "title": "登录管理后台",
        "showLoginForm": true,
        "isWrong": isWrong,
        "email": email,
    }
    c.HTML(200, "om_login.html", params)
}

