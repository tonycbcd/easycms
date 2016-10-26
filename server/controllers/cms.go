// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// The Api Controllers.

package controllers

import (
    "github.com/tonycbcd/easycms/server/components/egin"
)

func CmsDefault(c *egin.Context) {
    //module  := modules.NewArticle()
    // get the article list data.
    //_, list := module.GetList(models.M{"state": "pass"}, 1, 8, "created desc", true)

    // get the hot article list data.
    //_, hotList := module.GetList(models.M{"state": "pass"}, 1, 3, "commented desc, liked desc", false)

    params  := egin.H{
        "title": "登录管理后台",
        "customCss": "cms/css/index.css",
        "breadcrumbs": []egin.H {
            {"title": "首页", "link": ""},
        },
        //"list": list,
        //"hotList": hotList,
    }
    c.HTML(200, "cms_default.html", params)
}

