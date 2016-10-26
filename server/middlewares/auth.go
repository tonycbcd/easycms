// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// To initialize the Basic Admin Env.

package middlewares

import (
    "errors"
    "github.com/tonycbcd/easycms/server/components/egin"
    "github.com/gin-gonic/contrib/sessions"
)

var (
    CurAdminSession sessions.Session
)

func AdminAuthMiddleware(c *egin.Context) {
    session := sessions.Default(c.Context)
    userId  := session.Get("adminUserId")

    if userId != nil && userId != "" {
        CurAdminSession  = session
        c.Next()
    } else {
        if c.Request.Method == "POST" {
            c.Request.Method    = "GET"
        }
        c.Redirect(302, "/om/login")
    }
}

func GetCurAdminUserId() (userId string, err error) {
    if CurAdminSession == nil {
        err = errors.New("no cur admin session")
        return
    }
    id  := CurAdminSession.Get("adminUserId")
    if id == nil {
        err = errors.New("no admin user id in current session")
        return
    }

    idStr   := id.(string)
    if idStr == "" {
        err = errors.New("the user id is wrong")
        return
    }

    userId  = idStr

    return
}

