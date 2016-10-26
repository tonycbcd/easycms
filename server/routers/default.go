// Copyright 2016, Tonyxu All rights reserved.
// Author TonyXu <tonycbcd@gmail.com>
// Build on dev-0.0.1
// MIT Licensed

// The default routers.

package routers

import (
    "github.com/tonycbcd/easycms/server/components/egin"
    ctl "github.com/tonycbcd/easycms/server/controllers"
    // mid "git.18.tl/tony/liangli/server/middlewares"
)

func SetRouters(app *egin.Engine) {
    // 前端CMS页面路由.
    setCmsRouter(app.Group("/"))

    // 后台OM管理路由.
    setOmRouter(app.Group("/om"))

    // Api接口相关路由.
    setApiRouter(app.Group("/api"))
}

func setCmsRouter(g *egin.RouterGroup) {
    g.GET("/", ctl.CmsDefault)
    //g.GET("/articles", ctl.CmsArticles)
    //g.GET("/articles/:categoryId", ctl.CmsArticles)
    //g.GET("/search/articles", ctl.CmsSearch)
    //g.GET("/article/:articleId", ctl.CmsDetail)
}

func setOmRouter(g *egin.RouterGroup) {
    g.Any("/login", ctl.OmLogin)
    //g.GET("/", mid.AdminAuthMiddleware, ctl.OmDefault)
    //g.Any("/logout", mid.AdminAuthMiddleware, ctl.OmLogout)
    //g.Any("/setting", mid.AdminAuthMiddleware, ctl.OmSetting)

    //g.Any("/user", mid.AdminAuthMiddleware, ctl.OmUser)
    //g.Any("/user/add", mid.AdminAuthMiddleware, ctl.OmAddUser)
    //g.GET("/cms/category", mid.AdminAuthMiddleware, ctl.OmCmsCategory)
    //g.Any("/cms/category/add", mid.AdminAuthMiddleware, ctl.OmCmsAddCategory)
    //g.Any("/cms/category/edit", mid.AdminAuthMiddleware, ctl.OmCmsEditCategory)
    //g.GET("/cms/category/delete", mid.AdminAuthMiddleware, ctl.OmCmsDeleteCategory)
    //g.GET("/cms/article", mid.AdminAuthMiddleware, ctl.OmCmsArticle)
    //g.GET("/cms/article/add", mid.AdminAuthMiddleware, ctl.OmCmsAddArticle)
    //g.POST("/cms/article/add", mid.AdminAuthMiddleware, ctl.OmCmsAddArticle)
    //g.GET("/cms/article/edit/:articleId", mid.AdminAuthMiddleware, ctl.OmCmsAddArticle)
    //g.POST("/cms/article/edit/:articleId", mid.AdminAuthMiddleware, ctl.OmCmsAddArticle)
    //g.GET("/cms/article/delete/:articleId", mid.AdminAuthMiddleware, ctl.OmCmsDeleteArticle)
    //g.GET("/picture", mid.AdminAuthMiddleware, ctl.OmPicture)
    //g.GET("/tag", mid.AdminAuthMiddleware, ctl.OmTag)
    //g.GET("/keyword", mid.AdminAuthMiddleware, ctl.OmKeyword)
}

func setApiRouter(g *egin.RouterGroup) {
    //g.GET("/:controller/:action", ctl.ApiGet)
    //g.POST("/:controller/:action", ctl.ApiSet)
    //g.DELETE("/:controller/:action", ctl.ApiDelete)
    //g.OPTIONS("/:controller/:action", ctl.ApiOption)
}


